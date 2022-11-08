package model

import (
	"encoding/json"
	"time"

	"database/sql/driver"
	"gorm.io/gorm"
)

type GormList []string

func (g GormList) Value() (driver.Value, error) {
	return json.Marshal(g)
}

// Scan 实现sql.Scanner接口,Scan将value扫描至Jsonb
func (g *GormList) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &g)
}

type BaseModel struct {
	ID        int32     `gorm:"primarykey;type:int"` // why is int32
	CreatedAt time.Time `gorm:"column:add_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`
	DeleteAt  gorm.DeletedAt
	IsDeleted bool `gorm:"column:is_deleted"`
}
