package demo

import (
	"github.com/go-kit/kit/log"
	"gorm.io/gorm"
	"time"
)

type PersonRepository struct {
	DB     *gorm.DB
	logger *log.Logger
}

func NewPersonRepository(db *gorm.DB, logger *log.Logger) *PersonRepository {
	return &PersonRepository{
		DB:     db,
		logger: logger,
	}
}

func (p *PersonRepository) Migrate() (err error) {
	defer func(begin time.Time) {
		_ = p.logger.Log(
			"method", "migrate",
			"input", nil,
			"output", nil,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	return p.DB.AutoMigrate(&Person{})
}

func (p *PersonRepository) Save(person *[]Person) error {
	p.logger.("In Save function")

	return p.DB.Create(*person).Error
}
