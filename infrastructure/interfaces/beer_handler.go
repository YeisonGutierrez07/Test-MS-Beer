package interfaces

import (
	"net/http"
	"strconv"
	"test_ms_beer/application"
	"test_ms_beer/domain/entity"
	"test_ms_beer/domain/entity/test_ms_beer_entity"
	"test_ms_beer/infrastructure/persistence/base"
	"test_ms_beer/infrastructure/repository/beer"

	"github.com/gin-gonic/gin"
)

type Beer struct {
	BeerAppInterface application.BeerAppInterface
	Persistence      *base.Persistence
}

//constructor
func NewBeer(p *base.Persistence) *Beer {
	return &Beer{
		Persistence: p,
	}
}

func (a *Beer) GetAllBeers(c *gin.Context) {
	responseContextData := entity.ResponseContext{Ctx: c}

	//Get Repository
	a.BeerAppInterface = beer.NewBeerRepository(a.Persistence)
	data := a.BeerAppInterface.GetAllBeersRepository()

	c.JSON(http.StatusOK, responseContextData.ResponseData(entity.StatusSuccess, "Operacion exitosa", data))
}

func (a *Beer) CreateNewBeer(c *gin.Context) {
	responseContextData := entity.ResponseContext{Ctx: c}

	var newBear test_ms_beer_entity.Beer
	if err := c.ShouldBindJSON(&newBear); err != nil {
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(entity.StatusFail, "Request invalida", ""))
		return
	}

	//Get Repository
	a.BeerAppInterface = beer.NewBeerRepository(a.Persistence)
	data, err := a.BeerAppInterface.CreateNewBeerRepository(newBear)
	if err != nil {
		c.JSON(http.StatusConflict, responseContextData.ResponseData(entity.StatusFail, err.Error(), ""))
		return
	}

	c.JSON(http.StatusCreated, responseContextData.ResponseData(entity.StatusSuccess, "Operacion exitosa", data))
}

func (a *Beer) GetBeerByID(c *gin.Context) {
	responseContextData := entity.ResponseContext{Ctx: c}

	beerID, _ := strconv.ParseInt(c.Param("beerID"), 10, 64)

	//Get Repository
	a.BeerAppInterface = beer.NewBeerRepository(a.Persistence)
	data, err := a.BeerAppInterface.GetBeerByIDRepository(beerID)

	if err != nil {
		c.JSON(http.StatusNotFound, responseContextData.ResponseData(entity.StatusFail, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, responseContextData.ResponseData(entity.StatusSuccess, "Operacion exitosa", data))
}

func (a *Beer) GetBoxPriceBeers(c *gin.Context) {
	responseContextData := entity.ResponseContext{Ctx: c}

	beerID, _ := strconv.ParseInt(c.Param("beerID"), 10, 64)
	quantity, _ := strconv.ParseInt(c.Query("quantity"), 10, 64)
	currency := c.Query("currency")

	if currency == "" {
		c.JSON(http.StatusBadRequest, responseContextData.ResponseData(entity.StatusFail, "Request invalida", ""))
		return
	}

	if quantity == 0 {
		quantity = 6
	}

	//Get Repository
	a.BeerAppInterface = beer.NewBeerRepository(a.Persistence)
	data, err := a.BeerAppInterface.GetBoxPriceBeersRepository(beerID, quantity, currency)

	if err != nil {
		c.JSON(http.StatusNotFound, responseContextData.ResponseData(entity.StatusFail, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, responseContextData.ResponseData(entity.StatusSuccess, "Operacion exitosa", data))
}
