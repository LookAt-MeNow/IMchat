package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	Red *redis.Client
)

// viper 读取配置文件

func InitConfig() {
	viper.SetConfigName("app")      //config下的app.yml
	viper.AddConfigPath("./config") //相对路径
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config app: ", viper.Get("app"))
	fmt.Println("config mysql: ", viper.Get("mysql"))
}

func InitMySQL() {
	newLoggr := logger.New( // gorm日志配置
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{ // 日志配置
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,		// 启用彩色打印
		},
	)
	db, err := gorm.Open(mysql.Open(viper.GetString("mysql.DNS")), &gorm.Config{Logger: newLoggr}) // 配置文件读取连接MySQL
	if err != nil {
		fmt.Println(err)
	}
	DB = db // 全局变量赋值
}

func InitRedis() { // 配置文件读取连接Redis
	Red = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),    // 地址
		Password:     viper.GetString("redis.password"), // 密码
		DB:           viper.GetInt("redis.DB"),        // 数据库
		PoolSize:     viper.GetInt("redis.poolsize"),    // 连接池大小
		MinIdleConns: viper.GetInt("redis.minIdleConn"), // 最小连接数
	})
	pong, _ := Red.Ping(context.Background()).Result() // 测试连接
	fmt.Println("init redis..... ", pong)             // 打印连接结果
}

const (
	PublishKey = "websocket" // 发布消息的频道
)

// 发布消息到 Redis，channel 为频道，msg 为消息
func Publish(ctx context.Context, channel string, msg string) error {
	fmt.Println("publish msg: ", msg)
	err := Red.Publish(ctx, channel, msg).Err()
	return err
}

// 订阅消息到 Redis
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
