package dao

//import _ "github.com/go-sql-driver/mysql"
import "github.com/jinzhu/gorm"

var (
	DB *gorm.DB
)

// InitMysql 连接数据库
func InitMysql() (err error) {
	dsn := "test:123456@tcp(127.0.0.1:3306)/myblog?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}
