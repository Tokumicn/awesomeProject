package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

func main() {

	token := simpleDemo()
	time.Sleep(time.Second * 2)
	checkToken(token)
	withCustomClaimsToken()
}

var secretKey = []byte("mySecretKey") // 用于签名和验证的密钥

func checkToken(tokenString string) {
	// tokenString := "your_jwt_token_string_here" // 从客户端获取的 JWT 字符串

	// 解析并验证 JWT
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 校验 token 使用的签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}

		// 返回密钥用于验证签名
		return secretKey, nil
	})
	if err != nil {
		log.Fatal("Error parsing token:", err)
	}

	// 验证 token 是否有效
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("get claims error")
		return
	}

	if token.Valid {
		fmt.Println("Token is valid")
		fmt.Println("Claims:", claims)

		// 过期时间校验  Parse过程中就检验过过期时间等是否有效
		//if exp, ok := claims["exp"].(float64); ok {
		//	if int64(exp) < time.Now().Unix() {
		//		fmt.Println("Token has expired")
		//	} else {
		//		fmt.Println("Token is valid")
		//	}
		//}
	}

}

type CustomClaims struct {
	Username             string `json:"username"`
	jwt.RegisteredClaims        // 嵌套了 jwt.RegisteredClaims，以支持标准声明
}

func withCustomClaimsToken() {
	// 创建一个新的 token
	claims := &CustomClaims{
		Username: "john_doe",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "myApp",
			Subject:   "somebody",
			ID:        "1",
			Audience:  []string{"somebody_else"},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Fatal("Failed to sign token:", err)
	}

	fmt.Println("Generated Custom Token:", tokenString)
}

func simpleDemo() string {
	// 创建一个新的 token
	token := jwt.New(jwt.SigningMethodHS256)

	// 设置声明（Claims）
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "1234567890"
	claims["name"] = "John Doe"
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Second * 1).Unix() // 过期时间

	// 签名并产生完整的 token 字符串
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Fatalf("Error signing the token: %v", err)
	}

	fmt.Println("Generated Token:", tokenString)
	return tokenString
}
