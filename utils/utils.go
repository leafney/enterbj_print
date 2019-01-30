package utils

import (
	"crypto/md5"
	"fmt"
	"time"
)

//计算字符串的MD5值
func GetMd5Hash(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func CreateUnix() int64 {
	// timeLayout := "2006-01-02 15:04:05"
	//[Go 时间戳 · Go示例学 · 看云](https://www.kancloud.cn/itfanr/go-by-example/81697)
	return time.Now().Unix()
}
