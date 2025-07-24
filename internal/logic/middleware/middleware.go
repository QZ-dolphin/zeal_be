package middleware

import (
	"strings"
	"zeal_be/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/golang-jwt/jwt/v5"
)

type sMiddleware struct{}

func init() {
	service.RegisterMiddleware(New())
} // 用于服务初始化注册

func New() service.IMiddleware {
	return &sMiddleware{}
}

func (s *sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func (s *sMiddleware) Auth(r *ghttp.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		r.Response.WriteStatusExit(401)
	}
	ctx := gctx.New()
	jwt_key, _ := g.Cfg().Get(ctx, "jwt-secret")

	// 提取Token
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwt_key.String()), nil
	})

	// 验证Token
	if err != nil || !token.Valid {
		r.Response.WriteStatusExit(401)
	}

	// 存储用户邮箱信息到上下文
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		r.SetCtxVar("Email", claims["email"])
	} else {
		r.Response.WriteStatusExit(401)
	}

	r.Middleware.Next()
}
