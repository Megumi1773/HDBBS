package test

import (
	"Backend/config"
	"Backend/utils"
	"testing"
)

// 初始化测试配置
func init() {
	// 模拟配置初始化
	config.AppConfig = &config.Config{}
	config.AppConfig.Jwt.Key = "KURIYAMAMIRAI"
	config.AppConfig.Jwt.Exp = 24
}

// TestParseTokenWithInvalidInput 测试无效输入的token解析
func TestParseTokenWithInvalidInput(t *testing.T) {
	testCases := []struct {
		name        string
		tokenString string
		expectError bool
	}{
		{
			name:        "空字符串token",
			tokenString: "",
			expectError: true,
		},
		{
			name:        "随机字符串token",
			tokenString: "random_invalid_token",
			expectError: true,
		},
		{
			name:        "格式错误的JWT",
			tokenString: "invalid.jwt.format.extra",
			expectError: true,
		},
		{
			name:        "只有两个部分的JWT",
			tokenString: "header.payload",
			expectError: true,
		},
		{
			name:        "包含特殊字符的token",
			tokenString: "!@#$%^&*()",
			expectError: true,
		},
		{
			name:        "超长token",
			tokenString: string(make([]byte, 10000)),
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 测试不应该导致程序panic
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("ParseToken caused panic with input '%s': %v", tc.tokenString, r)
				}
			}()

			_, err := utils.ParseToken(tc.tokenString)
			
			if tc.expectError && err == nil {
				t.Errorf("Expected error for input '%s', but got none", tc.tokenString)
			}
			
			if !tc.expectError && err != nil {
				t.Errorf("Unexpected error for input '%s': %v", tc.tokenString, err)
			}
		})
	}
}

// TestCreateAndParseValidToken 测试创建和解析有效token
func TestCreateAndParseValidToken(t *testing.T) {
	username := "testuser"
	
	// 创建token
	tokenString, err := utils.CreateToken(username)
	if err != nil {
		t.Fatalf("Failed to create token: %v", err)
	}
	
	// 移除Bearer前缀进行解析
	if len(tokenString) > 6 && tokenString[:6] == "Bearer" {
		tokenString = tokenString[6:]
	}
	
	// 解析token
	claims, err := utils.ParseToken(tokenString)
	if err != nil {
		t.Fatalf("Failed to parse valid token: %v", err)
	}
	
	// 验证claims
	if claims == nil {
		t.Fatal("Claims should not be nil")
	}
}

// BenchmarkParseTokenInvalid 性能测试：解析无效token
func BenchmarkParseTokenInvalid(b *testing.B) {
	invalidToken := "invalid.token.string"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = utils.ParseToken(invalidToken)
	}
}
