package system

import (
	"context"
	"crypto/tls"
	"fmt"
	"strings"
	"zeal_be/utility"

	"github.com/gogf/gf/v2/frame/g"
	"gopkg.in/gomail.v2"
)

// 发送验证码
func (s *sSystem) EmailCodeSend(ctx context.Context, to string) (send bool, ttl int64) {
	code := utility.GenerateVerificationCode(6)
	key := "verification_code:" + to
	value, _ := g.Redis().Get(ctx, key)
	if !value.IsNil() {
		utility.Clog("code", value.String())
		ttl, _ = g.Redis().TTL(ctx, key)
		return true, ttl
	}

	smtpHost := g.Cfg().MustGet(ctx, "smtp.host").String()
	smtpPort := g.Cfg().MustGet(ctx, "smtp.port").Int()
	senderEmail := g.Cfg().MustGet(ctx, "smtp.email").String()
	senderPassword := g.Cfg().MustGet(ctx, "smtp.password").String()

	htmlBody := fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	<head>
		<meta charset="UTF-8">
		<title>您的 TechStack 验证码</title>
		<style>
			body { font-family: Arial, sans-serif; background-color: #1b2838; color: #ffffff; text-align: center; padding: 20px; }
			.container { max-width: 500px; margin: 0 auto; background-color: #2a475e; padding: 20px; border-radius: 10px; }
			.header { font-size: 24px; font-weight: bold; margin-bottom: 10px; }
			.code { font-size: 32px; font-weight: bold; color: #66c0f4; margin: 20px 0; }
			.footer { font-size: 14px; color: #bbbbbb; margin-top: 20px; }
			.signature { font-size: 14px; color: #bbbbbb; margin-top: 30px; border-top: 1px solid #444; padding-top: 10px; }
			.signature a { color: #66c0f4; text-decoration: none; }
		</style>
	</head>
	<body>
		<div class="container">
			<div class="header">您的 TechStack 验证码</div>
			<p>请使用以下验证码完成验证：</p>
			<div class="code">%s</div>
			<p>请在 <strong>5 分钟</strong> 内使用，否则验证码将失效。</p>
			<div class="footer">如果您未请求此验证码，请忽略此邮件。</div>
			<div class="signature">
				<hr>
				<p>
					由 <strong>TechStack</strong> 提供支持<br>
					官网: <a href="http://113.45.183.72/">Tech-Stack</a><br>
					邮箱: <a href="mailto:dearqueenszeal@163.com">dearqueenszeal@163.com</a>
				</p>
			</div>
		</div>
	</body>
	</html>
	`, code)

	m := gomail.NewMessage()

	m.SetHeader("From", senderEmail)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "您的 TechStack 验证码")
	m.SetBody("text/html", htmlBody)

	// 配置 SMTP 服务器（使用 SSL）
	d := gomail.NewDialer(smtpHost, smtpPort, senderEmail, senderPassword)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true} // 解决可能的证书问题
	utility.Clog("code", code)
	err := d.DialAndSend(m)
	if err != nil {
		utility.Clog(err)
		return false, 0
	}

	g.Redis().SetEX(ctx, key, code, 60*5)
	return true, 60 * 5
}

// 验证
func (s *sSystem) EmailCodeVerify(ctx context.Context, to string, code string) bool {
	key := "verification_code:" + to
	value, _ := g.Redis().Get(ctx, key)
	if value.IsNil() {
		return false
	}
	return strings.EqualFold(value.String(), code)
}

func (s *sSystem) EmailCodeSend2(ctx context.Context, to string) (send bool, ttl int64) {
	code := utility.GenerateVerificationCode(6)
	key := "verification_code:" + to
	value, _ := g.Redis().Get(ctx, key)
	if !value.IsNil() {
		utility.Clog("code", value.String())
		ttl, _ = g.Redis().TTL(ctx, key)
		return true, ttl
	}
	utility.Clog("code", code)
	g.Redis().SetEX(ctx, key, code, 60*5)
	return true, 60 * 5
}
