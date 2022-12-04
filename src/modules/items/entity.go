package items

import (
	"fmt"
)

const POSTGRES_TABLENAME = "items"

type Table interface {
	TableName() string
}

func (Item) GetName() string {
	return fmt.Sprintf("%s.persons", POSTGRES_TABLENAME)
}

//func (p *Item) BeforeCreate(tx *gorm.DB) (err error) {
//	return
//}
