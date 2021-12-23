package repository

import "test_ms_beer/domain/entity/test_ms_beer_entity"

// Repositorio central del dominio / entidad
type BeerRepository interface {
	GetAllBeersRepository() []test_ms_beer_entity.Beer
	CreateNewBeerRepository(beer test_ms_beer_entity.Beer) (test_ms_beer_entity.Beer, error)
	GetBeerByIDRepository(beerID int64) (test_ms_beer_entity.Beer, error)
	GetBoxPriceBeersRepository(beerID int64, quantity int64, currency string) (test_ms_beer_entity.BoxPriceBeer, error)
}
