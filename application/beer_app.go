package application

import (
	"test_ms_beer/domain/entity/test_ms_beer_entity"
	"test_ms_beer/domain/repository"
)

// Este archivo es la interface que sera usada por el handler
type BeerApp struct {
	repository repository.BeerRepository
}

type BeerAppInterface interface {
	GetAllBeersRepository() []test_ms_beer_entity.Beer
	CreateNewBeerRepository(beer test_ms_beer_entity.Beer) (test_ms_beer_entity.Beer, error)
	GetBeerByIDRepository(beerID int64) (test_ms_beer_entity.Beer, error)
	GetBoxPriceBeersRepository(beerID int64, quantity int64, currency string) (test_ms_beer_entity.BoxPriceBeer, error)
}
