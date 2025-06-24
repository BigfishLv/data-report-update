package dependencies

import (
	"data-report-update/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQLClient(config *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&loc=%s&parseTime=True",
		config.MySQL.Username, config.MySQL.Password, config.MySQL.Host, config.MySQL.Port, config.MySQL.Database, config.MySQL.Charset, config.MySQL.TimeZone)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Connect mysql on %s:%s failed", config.MySQL.Host, config.MySQL.Port))
		return nil
	}
	return db
}
