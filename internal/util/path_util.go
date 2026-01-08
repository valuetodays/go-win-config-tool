package util

import (
	"os"
	"regexp"
)

func ExpandWinEnv(s string) string {
	re := regexp.MustCompile(`%([^%]+)%`)
	return re.ReplaceAllStringFunc(s, func(match string) string {
		// 去掉 % 符号
		key := match[1 : len(match)-1]
		return os.Getenv(key)
	})
}
