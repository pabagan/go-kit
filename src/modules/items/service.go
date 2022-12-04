package items

import (
	"errors"
)

type ItemService interface {
	Save(item Item) (Item, error)
}

type itemService struct{}

func (itemService) Save(item Item) (Item, error) {
	if item.Name == "" {
		return item, ErrEmpty
	}
	return item, nil
}

var ErrEmpty = errors.New("empty string")
