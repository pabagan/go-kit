package items

import (
	"time"
)

type Item struct {
	Name      string `gorm:"column:name" json:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
