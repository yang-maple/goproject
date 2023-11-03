package model

import (
	"strconv"
	"time"
)

// GetAge 获取node 时间戳
func GetAge(CTT int64) string {
	age := time.Now().Unix() - CTT
	switch {
	case age < 60:
		//秒
		return strconv.FormatInt(age, 10) + "s"
	case age < 3600:
		//分
		return strconv.FormatInt(age/60, 10) + "min"
	case age < 86400:
		//时
		return strconv.FormatInt(age/3600, 10) + "h"
	case age < 31536000:
		//天
		return strconv.FormatInt(age/86400, 10) + "d"
	default:
		//年
		return strconv.FormatInt(age/31536000, 10) + "y"
	}

}
