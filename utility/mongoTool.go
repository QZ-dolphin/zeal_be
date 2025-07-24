package utility

import (
	"sync"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	Address  string // MongoDB server address in URI format
	Database string // Target database name
}

var (
	mongoClient *mongo.Client
	once        sync.Once
)

func NewMongoClient() (*mongo.Client, error) {
	var err error
	once.Do(func() {
		var (
			ctx    = gctx.GetInitCtx()
			config *MongoConfig
		)
		// Load MongoDB configuration from config.yaml
		err = g.Cfg().MustGet(ctx, "mongo").Scan(&config)
		if err != nil {
			return
		}
		if config == nil {
			err = gerror.New("mongo config not found")
			return
		}
		g.Log().Debugf(ctx, "Mongo Config: %s", config)

		// Initialize MongoDB client with the loaded configuration
		clientOptions := options.Client().ApplyURI(config.Address)
		mongoClient, err = mongo.Connect(ctx, clientOptions)
	})
	return mongoClient, err
}
