package admin

import (
	"context"
	"encoding/json"
	v1 "zeal_be/api/admin/v1"
	"zeal_be/internal/consts"
	"zeal_be/internal/dao"
	"zeal_be/internal/service"
	"zeal_be/utility"

	"github.com/gogf/gf/v2/frame/g"
)

type sAdmin struct{}

func init() {
	service.RegisterAdmin(New())
}

func New() service.IAdmin {
	return &sAdmin{}
}

func (s *sAdmin) CheckAdmin(ctx context.Context) int {
	r := g.RequestFromCtx(ctx)
	user_email := r.GetCtxVar("Email").String()
	md := dao.Userprivilege.Ctx(ctx)
	var privilege = struct {
		Privilege string `json:"privilege"`
	}{}
	md.Where("email_address", user_email).Fields("privilege").Scan(&privilege)
	var P []string
	json.Unmarshal([]byte(privilege.Privilege), &P)
	if len(P) == 0 {
		return consts.UserNoPrivilege
	}
	for _, p := range P {
		if p == "admin" {
			return 0
		}
	}
	return consts.UserNoPrivilege
}

func (s *sAdmin) GetUserList(ctx context.Context, in v1.ULReqInfo) (stat int, userList []v1.UserInfo, count int) {
	stat = s.CheckAdmin(ctx)
	if stat == consts.UserNoPrivilege {
		return
	}
	md := dao.Userdata.Ctx(ctx)
	if in.Order == "desc" {
		md = md.OrderDesc(in.OrderItem)
	} else if in.Order == "asc" {
		md = md.OrderAsc(in.OrderItem)
	}

	c := make(chan int, 1)
	defer close(c)
	if in.NeedCount == 1 {
		go func() {
			num, _ := md.Count()
			c <- num
		}()
	}

	md = md.Page(in.PageIndex, in.PageSize).Fields("user_id")
	g.Model("userdata as a,? as b", md).Fields("a.*").Where("b.user_id = a.user_id").Scan(&userList)

	for i := range userList {
		userList[i].Password, _ = utility.Crypt.Decrypt(userList[i].Password)
		if len(userList[i].Avatar_url) != 0 {
			userList[i].Avatar_url = g.Cfg().MustGet(ctx, "base_url").String() + userList[i].Avatar_url
		}
	}
	if in.NeedCount == 1 {
		count = <-c
	}
	return
}

func (s *sAdmin) GetUserPrivs(ctx context.Context, t_email string) (privs []string) {
	md := dao.Userprivilege.Ctx(ctx)
	var privilege = struct {
		Privilege string `json:"privilege"`
	}{}
	md.Where("email_address", t_email).Fields("privilege").Scan(&privilege)
	json.Unmarshal([]byte(privilege.Privilege), &privs)
	return

}

func (s *sAdmin) UpdateUserPrivs(ctx context.Context, t_email string, privs []string) {
	privilegeJSON, _ := json.Marshal(privs)
	dao.Userprivilege.Ctx(ctx).Where("email_address", t_email).Data("privilege", privilegeJSON).Update()
}
