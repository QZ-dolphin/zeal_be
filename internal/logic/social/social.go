package social

import (
	"context"
	v1 "zeal_be/api/social/v1"
	"zeal_be/internal/dao"
	"zeal_be/internal/service"
	"zeal_be/utility"

	"github.com/gogf/gf/v2/frame/g"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type sSocial struct{}

func init() {
	service.RegisterSocial(New())
}

func New() service.ISocial {
	return &sSocial{}
}

func (s *sSocial) SearchId(ctx context.Context, in v1.SearchReqInfo) (user_list []v1.User_info, total int) {
	id_like := "%" + in.ID_desc + "%"
	user_list = make([]v1.User_info, 0)

	dao.Userdata.Ctx(ctx).Wheref(`user_name like ? or email_address=?`, id_like, in.ID_desc).
		WhereNot("email_address", g.RequestFromCtx(ctx).GetCtxVar("Email").String()).
		Page(in.Page, in.PageSize).
		ScanAndCount(&user_list, &total, false)

	for i := range user_list {
		if len(user_list[i].Avatar_url) != 0 {
			user_list[i].Avatar_url = g.Cfg().MustGet(ctx, "base_url").String() + user_list[i].Avatar_url
		}
		user_list[i].IsFollower = s.IsFollower(ctx, user_list[i].User_id)
		user_list[i].IsFollowing = s.IsFollowing(ctx, user_list[i].User_id)
	}
	return
}

func (s *sSocial) IsFollowing(ctx context.Context, user_id int) bool {
	r := g.RequestFromCtx(ctx)
	user_email := r.GetCtxVar("Email").String()
	mongoClient, _ := utility.NewMongoClient()
	collection := mongoClient.Database("zeal").Collection("social")
	filter := bson.D{
		{Key: "user_email", Value: user_email},
		{Key: "following", Value: bson.D{{Key: "$elemMatch", Value: bson.D{{Key: "user_id", Value: user_id}}}}},
	}
	count, _ := collection.CountDocuments(ctx, filter)
	return count > 0
}

func (s *sSocial) IsFollower(ctx context.Context, user_id int) bool {
	r := g.RequestFromCtx(ctx)
	user_email := r.GetCtxVar("Email").String()
	mongoClient, _ := utility.NewMongoClient()
	collection := mongoClient.Database("zeal").Collection("social")
	filter := bson.D{
		{Key: "user_email", Value: user_email},
		{Key: "followers", Value: bson.D{{Key: "$elemMatch", Value: bson.D{{Key: "user_id", Value: user_id}}}}},
	}
	count, _ := collection.CountDocuments(ctx, filter)
	return count > 0
}

func (s *sSocial) IsInSocial(ctx context.Context) bool {
	r := g.RequestFromCtx(ctx)
	user_email := r.GetCtxVar("Email").String()
	mongoClient, _ := utility.NewMongoClient()
	collection := mongoClient.Database("zeal").Collection("social")
	filter := bson.D{{Key: "user_email", Value: user_email}}
	count, _ := collection.CountDocuments(ctx, filter)
	return count > 0
}

func (s *sSocial) CreateSocialAccount(ctx context.Context) {
	r := g.RequestFromCtx(ctx)
	user_email := r.GetCtxVar("Email").String()
	user_data := &struct {
		UserId          int    `json:"user_id"`
		User_name       string `json:"user_name"`
		Email_address   string `json:"email_address"`
		Created_time    string `json:"created_time"`
		Last_login_time string `json:"last_login_time"`
	}{}
	g.Model("userdata").Safe().Where("email_address", user_email).Scan(&user_data)
	mongoClient, _ := utility.NewMongoClient()
	collection := mongoClient.Database("zeal").Collection("social")
	doc := bson.D{
		{Key: "user_id", Value: user_data.UserId},
		{Key: "user_email", Value: user_data.Email_address},
		{Key: "user_name", Value: user_data.User_name},
		{Key: "created_time", Value: user_data.Created_time},
		{Key: "last_login_time", Value: user_data.Last_login_time},
		{Key: "following", Value: []bson.D{}},
		{Key: "followers", Value: []bson.D{}},
	}
	collection.InsertOne(ctx, doc)
}

func (s *sSocial) IsInSocialv2(ctx context.Context, user_id int) bool {
	mongoClient, _ := utility.NewMongoClient()
	collection := mongoClient.Database("zeal").Collection("social")
	filter := bson.D{{Key: "user_id", Value: user_id}}
	count, _ := collection.CountDocuments(ctx, filter)
	return count > 0
}

func (s *sSocial) CreateSocialAccountV2(ctx context.Context, user_id int) {
	user_data := &struct {
		UserId          int    `json:"user_id"`
		User_name       string `json:"user_name"`
		Email_address   string `json:"email_address"`
		Created_time    string `json:"created_time"`
		Last_login_time string `json:"last_login_time"`
	}{}
	g.Model("userdata").Safe().Where("user_id", user_id).Scan(&user_data)
	mongoClient, _ := utility.NewMongoClient()
	collection := mongoClient.Database("zeal").Collection("social")
	doc := bson.D{
		{Key: "user_id", Value: user_data.UserId},
		{Key: "user_email", Value: user_data.Email_address},
		{Key: "user_name", Value: user_data.User_name},
		{Key: "created_time", Value: user_data.Created_time},
		{Key: "last_login_time", Value: user_data.Last_login_time},
		{Key: "following", Value: []bson.D{}},
		{Key: "followers", Value: []bson.D{}},
	}
	collection.InsertOne(ctx, doc)
}

func (s *sSocial) GetSocialInfo(ctx context.Context) (following, followers int) {
	r := g.RequestFromCtx(ctx)
	user_email := r.GetCtxVar("Email").String()
	mongoClient, _ := utility.NewMongoClient()
	collection := mongoClient.Database("zeal").Collection("social")
	// Aggregate pipeline to only get the length of the following and followers arrays
	pipeline := bson.A{
		bson.D{{Key: "$match", Value: bson.D{{Key: "user_email", Value: user_email}}}},
		bson.D{{Key: "$project", Value: bson.D{
			{Key: "following_count", Value: bson.D{{Key: "$size", Value: "$following"}}},
			{Key: "followers_count", Value: bson.D{{Key: "$size", Value: "$followers"}}},
			{Key: "_id", Value: 0},
		}}},
	}

	cur, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return 0, 0
	}
	defer cur.Close(ctx)

	var result struct {
		FollowingCount int `bson:"following_count"`
		FollowersCount int `bson:"followers_count"`
	}

	if cur.Next(ctx) {
		err = cur.Decode(&result)
		if err != nil {
			return 0, 0
		}
		return result.FollowingCount, result.FollowersCount
	}

	// If no document is found, return 0 for both counts
	return 0, 0
}

func (s *sSocial) ToFollowing(ctx context.Context, user_id int, user_name string, nick_name string, avatar_url string) bool {
	r := g.RequestFromCtx(ctx)
	user_email := r.GetCtxVar("Email").String()
	user_data := &struct {
		UserId     int    `json:"user_id"`
		User_name  string `json:"user_name"`
		Avatar_url string `json:"avatar_url"`
	}{}

	g.Model("userdata").Safe().Where("email_address", user_email).Scan(&user_data)

	mongoClient, _ := utility.NewMongoClient()
	collection := mongoClient.Database("zeal").Collection("social")
	filter := bson.D{{Key: "user_email", Value: user_email}}
	update := bson.D{
		{Key: "$addToSet", Value: bson.D{
			{Key: "following", Value: bson.D{
				{Key: "user_id", Value: user_id},
				{Key: "user_name", Value: user_name},
				{Key: "nick_name", Value: nick_name},
				{Key: "avatar_url", Value: avatar_url},
			}},
		}},
	}
	_, err := collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return false
	}

	// Update the other user's followers list
	if !s.IsInSocialv2(ctx, user_id) {
		s.CreateSocialAccountV2(ctx, user_id)
	}

	filter2 := bson.D{{Key: "user_id", Value: user_id}}
	update2 := bson.D{
		{Key: "$addToSet", Value: bson.D{
			{Key: "followers", Value: bson.D{
				{Key: "user_id", Value: user_data.UserId},
				{Key: "user_email", Value: user_email},
				{Key: "user_name", Value: user_data.User_name},
				{Key: "nick_name", Value: user_data.User_name},
				{Key: "avatar_url", Value: user_data.Avatar_url},
			}},
		}},
	}
	_, err2 := collection.UpdateOne(ctx, filter2, update2)
	return err2 == nil
}

func (s *sSocial) NotFollowing(ctx context.Context, user_id int) bool {
	r := g.RequestFromCtx(ctx)
	user_email := r.GetCtxVar("Email").String()
	mongoClient, _ := utility.NewMongoClient()
	collection := mongoClient.Database("zeal").Collection("social")
	filter := bson.D{{Key: "user_email", Value: user_email}}
	update := bson.D{
		{Key: "$pull", Value: bson.D{
			{Key: "following", Value: bson.D{
				{Key: "user_id", Value: user_id},
			}},
		}},
	}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return false
	}
	// Update the other user's followers list
	filter2 := bson.D{{Key: "user_id", Value: user_id}}
	update2 := bson.D{
		{Key: "$pull", Value: bson.D{
			{Key: "followers", Value: bson.D{
				{Key: "user_email", Value: user_email},
			}},
		}},
	}
	_, err2 := collection.UpdateOne(ctx, filter2, update2)
	return err2 == nil
}

func (s *sSocial) GetFollowerList(ctx context.Context, pageSize, pageIndex int) []v1.Userdata {
	r := g.RequestFromCtx(ctx)
	user_email := r.GetCtxVar("Email").String()
	mongoClient, _ := utility.NewMongoClient()
	collection := mongoClient.Database("zeal").Collection("social")
	skip := pageSize * (pageIndex - 1)
	filter := bson.D{{Key: "user_email", Value: user_email}}
	// utility.Clog(pageSize, pageIndex, skip)
	projection := bson.D{
		{Key: "followers", Value: bson.D{
			{Key: "$slice", Value: bson.A{skip, pageSize}},
		}},
	}
	var result struct {
		Followers []v1.Userdata `bson:"followers"`
	}
	err := collection.FindOne(ctx, filter, options.FindOne().SetProjection(projection)).Decode(&result)
	if err != nil {
		// 处理错误，这可能是因为没有找到文档或解码失败
		utility.Clog("Decode error:", err)
		return nil
	}

	for i := range result.Followers {
		result.Followers[i].IsFollowing = s.IsFollowing(ctx, result.Followers[i].UserId)
		result.Followers[i].IsFollower = true
		if len(result.Followers[i].Avatar_url) != 0 {
			result.Followers[i].Avatar_url = g.Cfg().MustGet(ctx, "base_url").String() + result.Followers[i].Avatar_url
		}
	}
	// utility.Clog(result)
	return result.Followers
}

func (s *sSocial) GetFollowingList(ctx context.Context, pageSize, pageIndex int) []v1.Userdata {
	r := g.RequestFromCtx(ctx)
	user_email := r.GetCtxVar("Email").String()
	mongoClient, _ := utility.NewMongoClient()
	collection := mongoClient.Database("zeal").Collection("social")
	skip := pageSize * (pageIndex - 1)
	filter := bson.D{{Key: "user_email", Value: user_email}}
	projection := bson.D{
		{Key: "following", Value: bson.D{
			{Key: "$slice", Value: bson.A{skip, pageSize}},
		}},
	}
	var result struct {
		Following []v1.Userdata `bson:"following"`
	}
	collection.FindOne(ctx, filter, options.FindOne().SetProjection(projection)).Decode(&result)
	for i := range result.Following {
		result.Following[i].IsFollower = s.IsFollower(ctx, result.Following[i].UserId)
		result.Following[i].IsFollowing = true
		// result.Following[i].Avatar_url = g.Cfg().MustGet(ctx, "base_url").String() + result.Following[i].Avatar_url
		// 因为关注列表的avatar_url是从前端返回的，所以这里不需要再次拼接base_url
	}
	// utility.Clog(result)
	return result.Following
}
