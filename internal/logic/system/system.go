package system

import (
	"context"
	"encoding/json"
	"time"
	"zeal_be/internal/consts"
	"zeal_be/internal/service"
	"zeal_be/utility"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mojocn/base64Captcha"
)

type sSystem struct{}

func init() {
	service.RegisterSystem(New())
}

func New() service.ISystem {
	return &sSystem{}
}

var (
	driver = &base64Captcha.DriverString{
		Height:          80,
		Width:           240,
		NoiseCount:      50,
		ShowLineOptions: 20,
		Length:          4,
		Source:          "abcdefghjkmnpqrstuvwxyz23456789",
		Fonts:           []string{"chromohv.ttf"},
	}
	store = base64Captcha.DefaultMemStore
)

// 获取图形验证码
func (s *sSystem) GetVerifyImgString(ctx context.Context) (idKeyC string, base64stringC string, err error) {
	driver := driver.ConvertFonts()
	c := base64Captcha.NewCaptcha(driver, store)
	idKeyC, base64stringC, err = c.Generate()
	return
}

// 图形验证码验证
func (s *sSystem) VerifyCaptcha(ctx context.Context, idKey, captcha string) bool {
	return store.Verify(idKey, captcha, true)
}

// 登录功能
func (s *sSystem) Login(ctx context.Context, email string, password string) (stat int, tokenString string) {
	md := g.Model("userdata").Safe()
	exist, _ := md.Where("email_address", email).Exist()
	if !exist {
		return consts.UserNotExist, ""
	}

	ps, _ := md.Fields("password").Where("email_address", email).Value()
	d_pass, _ := utility.Crypt.Decrypt(ps.String())
	if ps.String() != password && password != d_pass {
		return consts.UserFalsePassword, ""
	} else if ps.String() == password {
		e_pass, _ := utility.Crypt.Encrypt(password)
		md.Where("email_address", email).Update(g.Map{"password": e_pass})
	}

	s.UpdateIPAddress(ctx, email)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":    email,
		"clientIP": g.RequestFromCtx(ctx).GetClientIp(),
		"exp":      time.Now().Add(time.Hour * 6).Unix(),
	})

	jwt_key, _ := g.Cfg().Get(ctx, "jwt-secret")
	tokenString, _ = token.SignedString([]byte(jwt_key.String()))
	return 0, tokenString
}

func (s *sSystem) UpdateIPAddress(ctx context.Context, email string) {
	md := g.Model("userdata").Safe()
	ip_address_json := struct {
		Ip_address string `json:"ip_address"`
	}{}
	md.Where("email_address", email).Fields("ip_address").Scan(&ip_address_json)
	ip_address := []string{}
	json.Unmarshal([]byte(ip_address_json.Ip_address), &ip_address)
	curr_ip := g.RequestFromCtx(ctx).GetClientIp()
	for i, v := range ip_address {
		if v == curr_ip {
			last_index := len(ip_address) - 1
			temp_ip := ip_address[last_index]
			ip_address[last_index] = curr_ip
			ip_address[i] = temp_ip
			save_ip, _ := json.Marshal(ip_address)
			md.Where("email_address", email).Data("ip_address", nil).Update()
			md.Where("email_address", email).Data("ip_address", save_ip).Update()
			return
		}
	}
	ip_address = append(ip_address, curr_ip)
	if len(ip_address) > 3 {
		ip_address = ip_address[len(ip_address)-3:]
	}
	save_ip, _ := json.Marshal(ip_address)
	md.Where("email_address", email).Data("ip_address", save_ip).Update()
}

// 注册
func (s *sSystem) Register(ctx context.Context, username string, email string, password string) int {
	md := g.Model("userdata").Safe()
	exist, _ := md.Where("email_address", email).Exist()
	if exist {
		return consts.UserAlreadyExist
	}
	password, _ = utility.Crypt.Encrypt(password)
	data := g.Map{
		"user_name":     username,
		"email_address": email,
		"password":      password,
	}
	md.Insert(data)
	userID, _ := md.Where("email_address", email).Fields("user_id").Value()
	g.Model("userprivilege").Safe().Insert(g.Map{
		"user_id":       userID.Int(),
		"email_address": email,
		"privilege":     []string{"stack"},
	})
	return 0
}

func (s *sSystem) UpdatePwd(ctx context.Context, email string, password string) int {
	md := g.Model("userdata").Safe()
	exist, _ := md.Where("email_address", email).Exist()
	if !exist {
		return consts.UserNotExist
	}
	md.Where("email_address", email).Data("password", password).Update()
	return 0
}
