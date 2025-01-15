package gmysql

import (
	"prometheus/pstruct"
	"time"
)

type AlertMessageSql struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	pstruct.Alerts
}

func (AlertMessageSql) TableName() string {
	return "alerts"
}
