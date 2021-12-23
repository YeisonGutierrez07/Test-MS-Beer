package beer

import (
	"errors"
	"test_ms_beer/domain/entity/test_ms_beer_entity"
	"test_ms_beer/domain/repository"
	"test_ms_beer/infrastructure/persistence/base"
	"test_ms_beer/infrastructure/utils"
)

// Repo que conecta con la DB
type BeerRepo struct {
	p *base.Persistence
}

func NewBeerRepository(p *base.Persistence) repository.BeerRepository {
	return &BeerRepo{p}
}

func (e *BeerRepo) GetAllBeersRepository() []test_ms_beer_entity.Beer {
	beersArray := []test_ms_beer_entity.Beer{}

	e.p.DbTestMSBeer.Find(&beersArray)

	return beersArray
}

func (e *BeerRepo) CreateNewBeerRepository(beer test_ms_beer_entity.Beer) (test_ms_beer_entity.Beer, error) {
	beerData := test_ms_beer_entity.Beer{}

	e.p.DbTestMSBeer.Where("id=?", beer.ID).First(&beerData)

	if beerData.ID == 0 {
		if errSave := e.p.DbTestMSBeer.Save(&beer).Error; errSave != nil {
			return beer, errSave
		}
	} else {
		return beer, errors.New("El ID de la cerveza ya existe")
	}

	return beer, nil
}

func (e *BeerRepo) GetBeerByIDRepository(beerID int64) (test_ms_beer_entity.Beer, error) {
	beer := test_ms_beer_entity.Beer{}

	err := e.p.DbTestMSBeer.Where("id=?", beerID).First(&beer).Error

	if err != nil {
		return beer, errors.New("El Id de la cerveza no existe")
	}

	return beer, err
}

func (e *BeerRepo) GetBoxPriceBeersRepository(beerID int64, quantity int64, currency string) (test_ms_beer_entity.BoxPriceBeer, error) {
	beer := test_ms_beer_entity.Beer{}
	boxPriceBeer := test_ms_beer_entity.BoxPriceBeer{}

	err := e.p.DbTestMSBeer.Where("id=?", beerID).First(&beer).Error

	if err != nil {
		return boxPriceBeer, errors.New("El Id de la cerveza no existe")
	}

	dataCurrencyLayer, err := utils.GetCurrencyData()
	if err != nil {
		return boxPriceBeer, err
	}

	priceInUSD := (beer.Price * float64(quantity)) / utils.GetCurrencyValueData(beer.Currency, dataCurrencyLayer)
	boxPriceBeer.PriceTotal = priceInUSD * utils.GetCurrencyValueData(currency, dataCurrencyLayer)
	boxPriceBeer.Bear = beer

	return boxPriceBeer, err
}

// Debe empatar con el repositorio del dominio y cumplir todos sus metodos
var _ repository.BeerRepository = &BeerRepo{}
