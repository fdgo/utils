package cron

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/robfig/cron"
	"log"
	"interview/baiy/model/mysql_mdl"
	"interview/baiy/support/lib/mysqlex"
	"interview/baiy/support/utils/constex"
	"interview/baiy/support/utils/timex"
)

// 6个参数:  秒(0-59)， 分(0-59)， 时(0-23)， 一个月中某天(1-31)， 月(1-12)， 星期几(0-6)
func Recycle() error {
	log.Println("Start Recycleing ...")
	c := cron.New()
	c.AddFunc("15 * * * * *", func() { //每分钟的15秒  (周期分钟)  比如 2018-07-03 14:25:15     2018-07-03 14:26:15    2018-07-03 14:27:15

	})
	c.AddFunc("10 * * * * *", func() { //每天夜里0点
		tx := mysqlex.Dbmysql().Begin()
		err := SetVipExperiedTo0(tx) //过期vip，设置查看相册和连麦次数为0
		if err != nil {
			return
		}
		err = SetVipOnedayCount(tx) //设置vip用户每天相册和连麦次数为0
		if err != nil {
			return
		}
		err = SetIdentifyOnedayCount(tx) //设置真人认证每天看相册和连麦次数
		if err != nil {
			return
		}
		err = SetNoVipOnedayCountForMainpage(tx) //设置非vip用户每天查看主页次数
		if err != nil {
			return
		}
		err = tx.Commit().Error
		if err != nil {
			tx.Rollback()
			return
		}
	})
	c.Start()
	select {}
}
func SetVipExperiedTo0(tx *gorm.DB) error { //设置vip过期用户每天可用次数为0
	var auth_experied []mysql_mdl.TbAuth
	err := tx.Where("consume_type=? and end_time <?", constex.AUTH_VIP_NUM, timex.GetCurrentTimeStamp()).Find(&auth_experied).Error
	if err != nil {
		tx.Rollback()
		return errors.New("服务正忙...")
	}
	for _, oneexperie := range auth_experied {
		err = tx.Model(&mysql_mdl.TbVipcount{}).Where("vipcount_id=?", oneexperie.UseradvID).Update().
			Update("oneday_counts", 0).
			Update("begin_time", timex.GetCurrentTimeStamp()).
			Update("end_time", timex.GetCurrentTimeStamp()).Error
		if err != nil {
			tx.Rollback()
			return errors.New("服务正忙...")
		}
	}
	return nil
}
func SetVipOnedayCount(tx *gorm.DB) error { //设置vip用户每天次数为全新总次数
	var auth_oneday []mysql_mdl.TbAuth
	err := tx.Where("consume_type=? and end_time >?", constex.AUTH_VIP_NUM, timex.GetCurrentTimeStamp()).Find(&auth_oneday).Error
	if err != nil {
		tx.Rollback()
		return errors.New("服务正忙...")
	}
	for _, oneexperie := range auth_oneday { //排除过期
		err = tx.Model(&mysql_mdl.TbVipcount{}).Where("vipcount_id=?", oneexperie.UseradvID).Update().
			Update("oneday_counts", constex.VIP_PHOABLUM_COUNT).
			Update("begin_time", timex.GetCurrentTimeStamp()).
			Update("end_time", timex.TimeStrToTimeStamp(timex.GetNextDayByCurrentTime(timex.GetCurrentTime()))).Error
		if err != nil {
			tx.Rollback()
			return errors.New("服务正忙...")
		}
	}
	return nil
}
func SetNoVipOnedayCountForMainpage(tx *gorm.DB) error { //非vip用户每天查看主页次数
	var vip_experied []mysql_mdl.TbAuth
	err := tx.Where("consume_type=? and end_time >?", constex.AUTH_VIP_NUM, timex.GetCurrentTimeStamp()).Find(&vip_experied).Error
	if err != nil {
		tx.Rollback()
		return errors.New("服务正忙...")
	}
	var identifs []mysql_mdl.TbAuth
	err = tx.Where("consume_type=? and is_effect=?", constex.AUTH_IDENTIFY_NUM, 1).Find(&identifs).Error
	if err != nil {
		tx.Rollback()
		return errors.New("服务正忙...")
	}
	var alluser []mysql_mdl.TbUserbase
	err = tx.Find(&alluser).Error
	if err != nil {
		tx.Rollback()
		return errors.New("服务正忙...")
	}
	for _, vip := range vip_experied {
		alluser = DeleteOther(alluser, vip.UseradvID)
	}
	for _, identify := range identifs {
		alluser = DeleteOther(alluser, identify.UseradvID)
	}
	for _, nor := range alluser {
		err = tx.Model(&mysql_mdl.TbFreecount{}).Where("freecount_id=?", nor.UserbaseID).
			Update("oneday_counts", constex.NORMAL_MAINPAGE_COUNT).
			Update("begin_time", timex.GetCurrentTimeStamp()).
			Update("end_time", timex.TimeStrToTimeStamp(timex.GetNextDayByCurrentTime(timex.GetCurrentTime()))).Error
		if err != nil {
			tx.Rollback()
			return errors.New("服务正忙...")
		}
	}
	return nil
}
func DeleteOther(allusers []mysql_mdl.TbUserbase, except string) []mysql_mdl.TbUserbase {
	for i := 0; i < len(allusers); i++ {
		if allusers[i].UserbaseID == except {
			allusers = append(allusers[:i], allusers[i+1:]...)
			i--
		}
	}
	return allusers
}
func SetIdentifyOnedayCount(tx *gorm.DB) error { //设置真人认证每天看相册和连麦次数
	var identifs []mysql_mdl.TbAuth
	err := tx.Where("consume_type=? and is_effect=?", constex.AUTH_IDENTIFY_NUM, 1).Find(&identifs).Error
	if err != nil {
		tx.Rollback()
		return errors.New("服务正忙...")
	}
	for _, oneexperie := range identifs {
		err = tx.Model(&mysql_mdl.TbIdetifycount{}).Where("idetifycount_id=?", oneexperie.UseradvID).Update().
			Update("oneday_counts", constex.IDENTIFY_PHOABLUM_COUNT).
			Update("begin_time", timex.GetCurrentTimeStamp()).
			Update("end_time", timex.TimeStrToTimeStamp(timex.GetNextDayByCurrentTime(timex.GetCurrentTime()))).Error
		if err != nil {
			tx.Rollback()
			return errors.New("服务正忙...")
		}
	}
	return nil
}
