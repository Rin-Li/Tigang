package conf

import (
	"Tigang/model"
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

var appConfig AppConfig
var mysqlConfig MySQLConfig
var db *gorm.DB

func InitConifg(){
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")

	if err := viper.ReadInConfig(); err != nil{
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	viper.Sub("service").Unmarshal(&appConfig)
	viper.Sub("mysql").Unmarshal(&mysqlConfig)
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

func Init(){
	InitConifg()
	InitMySQL(mysqlConfig)
}