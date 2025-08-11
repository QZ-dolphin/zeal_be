package user

import (
	"context"
	"time"
	v1 "zeal_be/api/user/v1"
	"zeal_be/internal/service"
	"zeal_be/utility"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() service.IUser {
	return &sUser{}
}

func (s *sUser) UpateAvatar(ctx context.Context, file *ghttp.UploadFile) (filepath string) {
	md := g.Model("userdata").Safe()
	r := g.RequestFromCtx(ctx)
	user_email := r.GetCtxVar("Email").String()
	avatar_path := "resource/public/upload/" + user_email + "/" + "avatar"
	file.Filename = "avatar.jpg"
	utility.ClearDir(avatar_path)
	file.Save(avatar_path)
	filepath = avatar_path + "/" + file.Filename
	md.Where("email_address", user_email).Update(g.Map{"avatar_url": filepath})
	filepath = g.Cfg().MustGet(ctx, "base_url").String() + filepath
	return
}

func (s *sUser) GetAvatar(ctx context.Context) (avatar string) {
	r := g.RequestFromCtx(ctx)
	user_email := r.GetCtxVar("Email").String()
	avatar_url, _ := g.Model("userdata").Safe().Where("email_address", user_email).Fields("avatar_url").Value()
	if len(avatar_url.String()) == 0 {
		return
	}
	return g.Cfg().MustGet(ctx, "base_url").String() + avatar_url.String()
}

func (s *sUser) SetAnniversary(ctx context.Context, day_name string, day_value string) {
	r := g.RequestFromCtx(ctx)
	user_email := r.GetCtxVar("Email").String()
	g.Model("userdata").Safe().Where("email_address", user_email).Update(g.Map{"day_name": day_name, "day_value": day_value})
}

func (s *sUser) GetAnniversary(ctx context.Context) (day_name string, day_value string) {
	r := g.RequestFromCtx(ctx)
	user_email := r.GetCtxVar("Email").String()
	var Data = struct {
		Day_name  string
		Day_value string
	}{}
	g.Model("userdata").Safe().Where("email_address", user_email).Fields("day_name", "day_value").Scan(&Data)
	return Data.Day_name, Data.Day_value
}

func (s *sUser) GetSelfInfo(ctx context.Context) (infoData *v1.SelfInfo) {
	r := g.RequestFromCtx(ctx)
	user_email := r.GetCtxVar("Email").String()
	g.Model("userdata").Safe().Where("email_address", user_email).Scan(&infoData)
	if len(infoData.Avatar_url) != 0 {
		infoData.Avatar_url = g.Cfg().MustGet(ctx, "base_url").String() + infoData.Avatar_url
	}
	return
}

type StackData struct {
	UserEmail string `bson:"user_email,omitempty"`
	UpdatedAt string `bson:"updated_at,omitempty"` // Last update timestamp
	StackText string `bson:"stack_text,omitempty"`
}

func (s *sUser) UpdateStackText(ctx context.Context, stacktext string) (stat bool, update_at string) {
	r := g.RequestFromCtx(ctx)
	user_email := r.GetCtxVar("Email").String()
	mongoClient, _ := utility.NewMongoClient()
	collection := mongoClient.Database("zeal").Collection("user")
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{Key: "user_email", Value: user_email}}
	const layout = "2006年1月2日 15:04"
	update_at = time.Now().Format(layout)
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "user_email", Value: user_email},
			{Key: "updated_at", Value: update_at},
			{Key: "stack_text", Value: stacktext},
		}},
		{Key: "$setOnInsert", Value: bson.D{
			{Key: "created_at", Value: update_at},
		}},
	}
	_, err := collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		utility.Clog("err", err)
		stat = false
		return
	}
	stat = true
	return
}

func (s *sUser) GetStackText(ctx context.Context) (stacktext string, update_at string) {
	r := g.RequestFromCtx(ctx)
	user_email := r.GetCtxVar("Email").String()
	mongoClient, _ := utility.NewMongoClient()
	collection := mongoClient.Database("zeal").Collection("user")
	filter := bson.D{{Key: "user_email", Value: user_email}}
	userdate := &StackData{}
	err := collection.FindOne(ctx, filter).Decode(&userdate)
	if err != nil {
		utility.Clog("err", err)
		return "# no_userdata", ""
	}
	// utility.Clog(userdate.UpdatedAt)
	return userdate.StackText, userdate.UpdatedAt
}
