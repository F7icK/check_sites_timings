package types

import (
	"time"

	"gorm.io/gorm"
)

type Site struct {
	ID            string         `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	CreatedAt     time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at" swaggertype:"string" format:"date-time"`
	Name          string         `json:"name" gorm:"column:name"`
	RequestTimeMs int64          `json:"request_time_ms" gorm:"column:request_time_ms"`
	StatusCode    int            `json:"status_code" gorm:"column:status_code"`
}

type History struct {
	ID        string         `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at" swaggertype:"string" format:"date-time"`
	Endpoint  string         `json:"endpoint" gorm:"column:endpoint"`
}

type Statistics struct {
	PathEndpoint string `json:"endpoint" gorm:"column:endpoint"`
	CountUse     int    `json:"count_use" gorm:"column:count_use"`
}
