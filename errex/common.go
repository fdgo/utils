package errex

import (
	"net/http"
)

var (
	ERROR_SYSTEMBASE = 10000
	ERROR_SYSTEM     = ERROR_SYSTEMBASE + 1 //系统在忙

	ERROR_NORMALBASE     = 20000
	ERROR_TOKEN_NEED     = ERROR_NORMALBASE + 1 //缺少token
	ERROR_TOKEN_EXPERIED = ERROR_NORMALBASE + 2 //toke过期
	ERROR_TOKEN_ERROR    = ERROR_NORMALBASE + 3 //toke错误
	ERROR_PARAM_ERROR    = ERROR_NORMALBASE + 4 //参数错误
	ERROR_USERACC_ERROR  = ERROR_NORMALBASE + 5 //用户账号错误
	ERROR_USERPWD_ERROR  = ERROR_NORMALBASE + 6 //用户密码错误

	ERROR_USERBASE                  = 30000
	ERROR_USER_NOTEXIST             = ERROR_USERBASE + 1   //用户不存在
	ERROR_USERDETAIL_MISS           = ERROR_USERBASE + 2   //用户资料不完善
	ERROR_USER_PHOTOALBUM_QUERY     = ERROR_USERBASE + 3   //用户相册查询失败
	ERROR_USER_BLACK_QUERY          = ERROR_USERBASE + 4   //用户黑名单查询失败
	ERROR_USER_FAV_QUERY            = ERROR_USERBASE + 5   //用户喜欢人员查询失败
	ERROR_USER_CARRY_QUERY          = ERROR_USERBASE + 6   //用户职业查询失败
	ERROR_USER_APPRAISE_QUERY       = ERROR_USERBASE + 7   //用户评论查询失败
	ERROR_USER_AUTH_QUERY           = ERROR_USERBASE + 8   //用户认真方式查询失败
	ERROR_USER_ADDRESS_QUERY        = ERROR_USERBASE + 9   //用户常住地址查询失败
	ERROR_USER_PGRAM_QUERY        = ERROR_USERBASE + 10  //用户交友节目查询失败
	ERROR_USER_OBJEXP_QUERY         = ERROR_USERBASE + 11  //用户期望对象查询失败
	ERROR_USER_REMARK_QUERY         = ERROR_USERBASE + 12  //用户备注人员查询失败
	ERROR_USER_HEIGHT_QUERY         = ERROR_USERBASE + 100 //用户体重查询失败
	ERROR_USER_WEIGHT_QUERY         = ERROR_USERBASE + 101 //用户身高查询失败
	ERROR_USER_BSTPRO_QUERY         = ERROR_USERBASE + 102 //用户广播节目查询失败
	ERROR_USER_MIKETAB_QUERY        = ERROR_USERBASE + 103 //用户miketab查询失败
	ERROR_USER_PARKFEMALETABS_QUERY = ERROR_USERBASE + 104 //用户ParkFemaleTabs查询失败
	ERROR_USER_DATETIMETABS_QUERY   = ERROR_USERBASE + 105 //用户DatetimeTabs查询失败
	ERROR_USER_NEWREGISTER_QUERY    = ERROR_USERBASE + 106 //用户最新注册失败
	ERROR_USER_GODDNESS_QUERY       = ERROR_USERBASE + 107 //用户女神失败
	ERROR_USER_VIP_QUERY            = ERROR_USERBASE + 108 //用户vip失败
	ERROR_USER_IDENTYCARD_QUERY     = ERROR_USERBASE + 109 //用户真人失败
	ERROR_USER_PHOTO_QUERY          = ERROR_USERBASE + 110 //用户照片查询失败
	ERROR_USER_DYNAMIC_QUERY        = ERROR_USERBASE + 111 //用户动态查询失败
	ERROR_USER_DYNAMICPHO_QUERY        = ERROR_USERBASE + 112 //用户动态照片查询失败
	ERROR_USER_PROGRAM_QUERY        = ERROR_USERBASE + 113 //用户节目查询失败
	ERROR_USER_PROGRAMPHO_QUERY        = ERROR_USERBASE + 114 //用户节目照片查询失败

	ERROR_USER_BASE_SAVE      = ERROR_USERBASE + 13 //用户基础信息保存失败
	ERROR_USER_CARRER_SAVE    = ERROR_USERBASE + 14 //用户职业信息保存失败
	ERROR_USER_PROGRAM_SAVE   = ERROR_USERBASE + 15 //用户交友节目保存失败
	ERROR_USER_OBJEXP_SAVE    = ERROR_USERBASE + 16 //用户期望对象保存失败
	ERROR_USER_ADDRESS_SAVE   = ERROR_USERBASE + 17 //用户地址信息保存失败
	ERROR_USER_LOGINTIME_SAVE = ERROR_USERBASE + 18 //用户登陆时间保存失败
	ERROR_USER_PHOALBUM_SAVE  = ERROR_USERBASE + 19 //用户保存相册隐私失败
	ERROR_USERDETAIL_EXIST    = ERROR_USERBASE + 20 //用户资料已经完善,无需重复填写

)

const BAD_REQUEST string = "BAD_REQUEST"
const UNAUTHORIZED string = "UNAUTHORIZED"
const FORBIDDEN string = "FORBIDDEN"
const NO_PERMISSION string = "NO_PERMISSION"
const INTERNAL_SERVER_ERROR string = "INTERNAL_SERVER_ERROR"
const INVALID_REQUEST_TYPE string = "INVALID_REQUEST_TYPE"
const SYSTEM_ERROR string = "SYSTEM_ERROR"
const INVALID_PARAMETER string = "INVALID_PARAMETER"
const NOT_FOUND string = "NOT_FOUND"
const interview/access_TOO_OFTEN string = "interview/access_TOO_OFTEN"
const CLIENT_TIME_INVALID string = "CLIENT_TIME_INVALID"

var (
	OK                     = NewError(http.StatusOK, "OK", "Success")
	BadRequest             = NewError(http.StatusOK, BAD_REQUEST, "Invalid request")
	Unauthorized           = NewError(http.StatusOK, UNAUTHORIZED, "unauthorized")
	Forbidden              = NewError(http.StatusOK, FORBIDDEN, "forbidden")
	NoPermission           = NewError(http.StatusOK, NO_PERMISSION, "no permission")
	InternalServerError    = NewError(http.StatusOK, INTERNAL_SERVER_ERROR, "server error")
	NotFound               = NewError(http.StatusOK, NOT_FOUND, "not found")
	ParameterError         = NewError(http.StatusOK, INVALID_PARAMETER, "Invalid parameter")
	InvalidReqError        = NewError(http.StatusOK, INVALID_REQUEST_TYPE, "Illegal request")
	SysError               = NewError(http.StatusOK, SYSTEM_ERROR, "System error")
	Sysinterview/access2OftenError   = NewError(http.StatusOK, interview/access_TOO_OFTEN, "interview/access too often")
	ClientTimeInvalidError = NewError(http.StatusOK, CLIENT_TIME_INVALID, "Please check your system time.")
)
