package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gqlgen-generator/config"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

// 初始化db
func Setup() {
	var dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		config.CONF_INSTANCE.DbConf.Host,
		config.CONF_INSTANCE.DbConf.Username,
		config.CONF_INSTANCE.DbConf.Password,
		config.CONF_INSTANCE.DbConf.Dbname,
		config.CONF_INSTANCE.DbConf.Port,
		config.CONF_INSTANCE.DbConf.SslMode,
		config.CONF_INSTANCE.DbConf.Timezone)

	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // Disable color
		},
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic(err)
	}

}
