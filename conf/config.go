package conf

import (
	"Tigang/repository/db/model"
	"context"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type AppConfig struct {
	AppModel string
	HttpPort string
}

type MySQLConfig struct {
	Db string
	DbHost string
	DbPort string
	DbUser string
	DbPassword string
	DbName string
}

type RedisConfig struct {
	RedisDb     string
	RedisAddr   string
	RedisPw     string
	RedisDbName int
}

type EmailConfig struct {
	SmtpPort int
	SmtpHost string
	SmtpEmail string
	SmtpPass string
}

var appConfig AppConfig
var mysqlConfig MySQLConfig
var redisConfig RedisConfig
var emailConfig EmailConfig
var db *gorm.DB

func InitConfig(){
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")

	if err := viper.ReadInConfig(); err != nil{
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	viper.Sub("service").Unmarshal(&appConfig)
	viper.Sub("mysql").Unmarshal(&mysqlConfig)
	viper.Sub("redis").Unmarshal(&redisConfig)
	viper.Sub("email").Unmarshal(&emailConfig)
}

func InitMySQL(cfg MySQLConfig){
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	cfg.DbUser,
	cfg.DbPassword,
	cfg.DbHost,
	cfg.DbPort,
	cfg.DbName,
)
	fmt.Println(dsn)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		panic(fmt.Errorf("fatal error db connection: %w", err))
	}

	err = db.AutoMigrate(
		&model.User{},
		&model.Achievement{},
		&model.Record{},
		&model.UserAchievement{},
	)

	if err != nil{
		panic(fmt.Errorf("fatal error db migration: %w", err))
	}

	fmt.Println("Database connected")

}

func InitRedis() RedisConfig{
	return redisConfig
}

func InitEmail() EmailConfig{
	return emailConfig
}


func GetDB(ctx context.Context) *gorm.DB{
	DB := db
	return DB.WithContext(ctx)
}

func Init(){
	InitConfig()
	InitMySQL(mysqlConfig)
}