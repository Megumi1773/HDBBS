package middle

import (
	"Backend/utils"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthTokenMiddle 认证中间件
func AuthTokenMiddle(c *gin.Context) {
	// 获取Authorization头
	auth := c.GetHeader("Authorization")
	if auth == "" {
		utils.Respond(c, http.StatusUnauthorized, 401, "未提供令牌", nil)
		c.Abort()
		return
	}

	// 正确解析Bearer token格式
	var token string
	if after, found := strings.CutPrefix(auth, "Bearer "); found {
		token = after
	} else if after, found := strings.CutPrefix(auth, "Bearer"); found {
		token = after
	} else {
		// 如果不是Bearer格式，直接使用原始值（兼容性处理）
		token = auth
	}

	// 验证token不为空
	token = strings.TrimSpace(token)
	if token == "" {
		utils.Respond(c, http.StatusUnauthorized, 401, "令牌格式无效", nil)
		c.Abort()
		return
	}

	// 使用defer recover防止token解析过程中的panic
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Token middleware panic recovered: %v", r)
			utils.Respond(c, http.StatusInternalServerError, 500, "服务器内部错误", nil)
			c.Abort()
		}
	}()

	// 解析token
	info, err := utils.ParseToken(token)
	if err != nil {
		// 详细的错误分类处理
		var msg string
		switch {
		case errors.Is(err, jwt.ErrSignatureInvalid):
			msg = "无效签名"
		case strings.Contains(err.Error(), "expired"):
			msg = "令牌已过期"
		case strings.Contains(err.Error(), "malformed"):
			msg = "令牌格式错误"
		case strings.Contains(err.Error(), "empty"):
			msg = "令牌为空"
		case strings.Contains(err.Error(), "unexpected signing method"):
			msg = "不支持的签名方法"
		case strings.Contains(err.Error(), "invalid token claims"):
			msg = "令牌内容无效"
		default:
			msg = "无效令牌"
		}

		log.Printf("Token validation failed: %v", err)
		utils.Respond(c, http.StatusUnauthorized, 401, msg, nil)
		c.Abort()
		return
	}

	// 将用户信息存储到上下文中，供后续处理使用
	if claims, ok := info.(jwt.MapClaims); ok {
		if username, exists := claims["username"]; exists {
			c.Set("username", username)
		}
		c.Set("claims", claims)
	}
	c.Next()
}
