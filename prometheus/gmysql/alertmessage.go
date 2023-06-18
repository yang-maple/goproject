package gmysql

import (
	"gorm.io/gorm"
	"prometheus/pstruct"
)

type AlertMessageSql struct {
	gorm.Model
	pstruct.Alerts
}

func (AlertMessageSql) TableName() string {
	return "alerts"
}
