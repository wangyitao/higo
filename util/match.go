package util

import "regexp"

// 匹配是否为ip
func IsIp(ip string) (b bool) {
	if m, _ := regexp.MatchString("^((25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.){3}(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)$", ip); !m {
		return false
	}
	return true
}
