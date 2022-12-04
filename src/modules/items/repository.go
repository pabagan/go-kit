package items

import (
	"github.com/go-kit/kit/log"
	"gorm.io/gorm"
	"time"
)

type ItemRepository struct {
	DB     *gorm.DB
	logger *log.LoggerFunc
}

func NewItemRepository(db *gorm.DB, logger *log.LoggerFunc) *ItemRepository {
	return &ItemRepository{
		DB:     db,
		logger: logger,
	}
}

func (p *ItemRepository) Migrate() (err error) {
	defer func(begin time.Time) {
		_ = p.logger.Log(
			"method", "migrate",
			"input", p,
			"output", nil,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	return p.DB.AutoMigrate(&Item{})
}

func (p *ItemRepository) Save(item *Item) (err error) {
	defer func(begin time.Time) {
		_ = p.logger.Log(
			"method", "Save",
			"input", p,
			"output", nil,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	return p.DB.Create(*item).Error
}
