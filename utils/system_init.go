package utils

import (
	"fmt"
	"log"
	"os"
	"time"
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
	Red *redis.Client
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config app: ", viper.Get("app"))
	fmt.Println("config mysql: ", viper.Get("mysql"))
}

func InitMySQL() {
	newLoggr := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	db, err := gorm.Open(mysql.Open(viper.GetString("mysql.DNS")), &gorm.Config{Logger: newLoggr})
	if err != nil {
		fmt.Println(err)
	}
	DB = db
}

func InitRedis() {
	Red = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.poolsize"),
		MinIdleConns: viper.GetInt("redis.minIdleConn"),
	})
	pong, _ := Red.Ping(context.Background()).Result()
	fmt.Println("init redis..... ", pong)
}

const(
	PublishKey = "websocket"
)

//发布消息到 Redis
func Publish(ctx context.Context, channel string, msg string) error {
	fmt.Println("publish msg: ", msg)
	err := Red.Publish(ctx, channel, msg).Err()
	return err
}

//订阅消息到 Redis
func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := Red.Subscribe(ctx, channel)
	msg, err := sub.ReceiveMessage(ctx)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println("sub msg: ", msg.Payload)
	return msg.Payload, err
}


