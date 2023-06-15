package initialize

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"tank_war/server/cmd/user/config"
	model "tank_war/server/cmd/user/pkg/mysql"
	"tank_war/server/shared/consts"
)

func InitDB() *gorm.DB {
	c := config.GlobalServerConfig.MysqlInfo
	dsn := fmt.Sprintf(consts.MysqlDSN, c.User, c.Password, c.Host, c.Port, c.Name)
	//klog.Infof("mysql dsn: %s", dsn)

	//global mode
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		klog.Fatalf("init gorm failed: %s", err.Error())
	}
	db.AutoMigrate(&model.User{})
	return db
}
