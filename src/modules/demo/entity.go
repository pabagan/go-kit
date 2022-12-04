package demo

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

const POSTGRES_TABLENAME = "person"

type Person struct {
	Name      string `gorm:"column:name" json:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Table interface {
	TableName() string
}

func (Person) GetName() string {
	return fmt.Sprintf("%s.persons", POSTGRES_TABLENAME)
}

func (p *Person) BeforeCreate(tx *gorm.DB) (err error) {
	//sads
	return
}
