package utils

import (
	"log"
	"runtime/debug"
)

// SafeExecute 安全执行函数，捕获panic并记录错误
func SafeExecute(fn func() error, operation string) error {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Panic recovered in %s: %v\nStack trace:\n%s", 
				operation, r, string(debug.Stack()))
		}
	}()
	
	return fn()
}

// LogError 统一的错误日志记录
func LogError(operation string, err error) {
	if err != nil {
		log.Printf("Error in %s: %v", operation, err)
	}
}

// IsTokenError 判断是否为token相关错误
func IsTokenError(err error) bool {
	if err == nil {
		return false
	}
	
	errorMsg := err.Error()
	tokenErrors := []string{
		"token",
		"expired",
		"invalid",
		"malformed",
		"signature",
		"claims",
	}
	
	for _, keyword := range tokenErrors {
		if contains(errorMsg, keyword) {
			return true
		}
	}
	
	return false
}

// contains 检查字符串是否包含子字符串（不区分大小写）
func contains(s, substr string) bool {
	return len(s) >= len(substr) && 
		   (s == substr || 
		    (len(s) > len(substr) && 
		     (s[:len(substr)] == substr || 
		      s[len(s)-len(substr):] == substr || 
		      indexOfSubstring(s, substr) >= 0)))
}

// indexOfSubstring 查找子字符串位置
func indexOfSubstring(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
