package postgreSQL

import (
	log "alpha/config"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"time"
	// MySQL driver.
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Database struct {
	Api *gorm.DB
	//Docker    *gorm.DB
}

var DB *Database

func openDB(username, password, host, name string, port int) *gorm.DB {
	config := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s",
		host,
		port,
		username,
		name,
		password,
	)

	db, err := gorm.Open("postgres", config)
	if err != nil {
		log.Logger.Error("Database connection failed. Database",
			zap.String("name", name),
		)
	}

	// set for db connection
	setupDB(db)

	return db
}

func setupDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))
	db.DB().SetMaxOpenConns(viper.GetInt("postgress.max_open_conns"))     // 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	db.DB().SetMaxIdleConns(viper.GetInt("postgress.wet_max_idle_conns")) // 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	t := time.Duration(viper.GetInt("postgress.conn_max_lifetime"))
	db.DB().SetConnMaxLifetime(time.Second * t) // 连接超时时间
}

// used for cli
func InitSelfDB() *gorm.DB {
	return openDB(
		viper.GetString("postgress.username"),
		viper.GetString("postgress.password"),
		viper.GetString("postgress.host"),
		viper.GetString("postgress.name"),
		viper.GetInt("postgress.port"),
	)
}

func GetSelfDB() *gorm.DB {
	return InitSelfDB()
}

func InitDockerDB() *gorm.DB {
	return openDB(
		viper.GetString("postgress.username"),
		viper.GetString("postgress.password"),
		viper.GetString("postgress.host"),
		viper.GetString("postgress.name"),
		viper.GetInt("postgress.port"),
	)
}

func GetDockerDB() *gorm.DB {
	return InitDockerDB()
}

func (db *Database) Init() {
	DB = &Database{
		Api: GetSelfDB(),
		//Docker: GetDockerDB(),
	}
}

func (db *Database) Close() {
	DB.Api.Close()
	//DB.Docker.Close()
}
