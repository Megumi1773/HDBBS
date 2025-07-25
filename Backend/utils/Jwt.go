package utils

import (
	"Backend/config"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

func CreateToken(username string) (string, error) {
	key := config.AppConfig.Jwt.Key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * time.Duration(config.AppConfig.Jwt.Exp)).Unix(),
	})
	s, err := token.SignedString([]byte(key))

	return "Bearer " + s, err
}

func ParseToken(tokenString string) (jwt.Claims, error) {
	// 输入验证
	if tokenString == "" {
		return nil, fmt.Errorf("token string is empty")
	}

	// 安全的token解析，使用defer recover防止panic
	defer func() {
		if r := recover(); r != nil {
			log.Printf("token parse panic recovered: %v", r)
		}
	}()

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.AppConfig.Jwt.Key), nil
	})

	if err != nil {
		// 记录错误但不使用panic，避免程序崩溃
		log.Printf("token parse error: %v", err)
		return nil, err
	}

	// 验证token有效性和claims
	if token == nil {
		return nil, fmt.Errorf("token is nil")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// 额外验证claims的基本结构
		if claims["username"] == nil || claims["exp"] == nil {
			return nil, fmt.Errorf("invalid token claims structure")
		}
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
