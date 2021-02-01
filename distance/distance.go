package distance

import (
	"math"
	"fmt"
)

func GetDistance(lat1, lng1, lat2, lng2 float64) float64 {
	radlat1 := lat1 * 3.1415926 / 180.0
	radlat2 := lat2 * 3.1415926 / 180.0
	radlng1 := lng1 * 3.1415926 / 180.0
	radlng2 := lng2 * 3.1415926 / 180.0
	ff := (radlat1 + radlat2) / 2.0
	gg := (radlat1 - radlat2) / 2.0
	ll := (radlng1 - radlng2) / 2.0
	ss := math.Pow((math.Sin(gg)), 2)*math.Pow((math.Cos(ll)), 2) + math.Pow((math.Cos(ff)), 2)*math.Pow((math.Sin(ll)), 2)
	cc := math.Pow((math.Cos(gg)), 2)*math.Pow((math.Cos(ll)), 2) + math.Pow((math.Sin(ff)), 2)*math.Pow((math.Sin(ll)), 2)
	ww := math.Atan(math.Sqrt(ss / cc))
	banjin := 6378.135 //地球半径km
	dist := 2 * ww * banjin
	rr := math.Sqrt(ss*cc) / ww
	h1 := (3*rr - 1) / (2 * cc)
	h2 := (3*rr + 1) / (2 * ss)
	xx := 1 / 298.257223543 //修正率
	var dm float64
	if ww != 0 {
		dm = dist * (1 + xx*h1*math.Pow((math.Sin(ff)), 2)*math.Pow((math.Cos(gg)), 2) - xx*h2*math.Pow((math.Cos(ff)), 2)*math.Pow((math.Sin(gg)), 2))
	} else {
		dm = 0
	}
	return dm
}
func Distance(dm float64) string {
	if dm < 1 {
		return fmt.Sprintf("%.0fm", dm*1000)
	}
	return fmt.Sprintf("%.1fkm", dm)
}
func GetDisEx(lat1, lng1, lat2, lng2 float64) string {
	dis := GetDistance(lat1, lng1, lat2, lng2)
	return Distance(dis)
}
