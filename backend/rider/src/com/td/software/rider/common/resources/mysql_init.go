package resources

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Db() error {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/?charset=utf8mb4&parseTime=True&loc=Local&timeout=10s"

	Database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}
