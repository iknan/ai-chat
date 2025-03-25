package middleware

import (
	"ai_chat/internal/config"
	redisInfra "ai_chat/internal/infra/redis"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/redis/go-redis/v9"
	"net/http"
)

type AuthTokenMiddleware struct {
	Config *config.Config
	redis  *redis.Client
}

func NewAuthTokenMiddleware(c *config.Config, client *redis.Client) *AuthTokenMiddleware {
	return &AuthTokenMiddleware{
		Config: c,
		redis:  client,
	}
}

func (m *AuthTokenMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.URL.Query().Get("token")

		if token == "" {
			// url没有，从header中获取
			token = r.Header.Get("Authorization")
		}

		exist := m.redis.Exists(r.Context(), fmt.Sprintf(redisInfra.KeyUserToken, token)).Val()
		if exist == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("token is not valid"))
			return
		}

		next(w, r)
	}
}

// 解析 Token 并获取 uid
func (m *AuthTokenMiddleware) ParseToken(tokenString string) (jwt.MapClaims, error) {
	// 会自动判断token是否过期
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 确保算法是 HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// 使用相同的秘钥进行验证
		return []byte(m.Config.JwtAuth.AccessSecret), nil
	})

	if err != nil {
		return nil, err
	}

	// 确保 Token 没有被篡改
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// 从 Payload 中获取 uid
		return claims, nil
	}
	return nil, fmt.Errorf("token is not valid")
}

func ParseToken(tokenString string, accessSecret string) (uid int64, isTourist bool, platform int, err error) {
	// 会自动判断token是否过期
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 确保算法是 HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// 使用相同的秘钥进行验证
		return []byte(accessSecret), nil
	})

	if err != nil {
		return
	}

	// 确保 Token 没有被篡改
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// 从 Payload 中获取 uid
		uid = int64(claims["uid"].(float64))
		platform = int(claims["platform"].(float64))
		isTourist, _ = claims["isTourist"].(bool)
		return
	}
	err = fmt.Errorf("token is not valid")
	return
}
