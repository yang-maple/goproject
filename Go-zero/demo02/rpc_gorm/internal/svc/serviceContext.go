package svc

import (
	"demo02/common/init_db"
	"demo02/models"
	"demo02/rpc_gorm/internal/config"
	"fmt"

	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDb := init_db.InitGorm(c.Mysql.DataSource)
	mysqlDb.AutoMigrate(&models.User{})
	fmt.Println(mysqlDb.Migrator().CurrentDatabase())
	return &ServiceContext{
		Config: c,
		DB:     mysqlDb,
	}
}
