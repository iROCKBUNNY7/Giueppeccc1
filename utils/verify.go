package utils

import (
	"strings"
)

func IsMd5Str(str string) bool {
	return regexpURLParse.MatchString(str)
}

//IsType 判断类型是否允许上传
func IsType(typeStr string) bool {
	for _, v := range ImageTypes {
		if strings.Contains(typeStr, strings.ToLower(v)) {
			return true
		}
	}

	return false
}

func IsAllow(ip string) bool {
	if len(AdminIPs) == 1 && AdminIPs[0] == "*" {
		return true
	}

	for _, v := range AdminIPs {
		if v == ip {
			return true
		}
	}

	return false
}
