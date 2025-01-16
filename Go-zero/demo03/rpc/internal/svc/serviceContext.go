package svc

import (
	"demo03/common/init_db"
	"demo03/models"
	"demo03/rpc/internal/config"

	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqldb := init_db.InitGorm(c.Mysql.DataSource)
	sqldb.AutoMigrate(&models.User{})
	return &ServiceContext{
		Config: c,
		DB:     sqldb,
	}
}
