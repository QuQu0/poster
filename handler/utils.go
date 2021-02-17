package handler

import (
	"net/http"
	"os"
)

// IsURLAvailable ... 判断URL是否可用
func IsURLAvailable(url string) bool {
	header, err := http.Get(url)
	if err != nil || header.Header["X-Errno"] != nil {
		return false
	}
	return true
}

// IsFileExist ... 判断本地文件是否存在
func IsFileExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}
