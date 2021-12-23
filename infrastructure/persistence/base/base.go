package base

import (
	"log"
	"test_ms_beer/domain/entity/test_ms_beer_entity"
	"test_ms_beer/infrastructure/persistence/base/db"

	"gorm.io/gorm"
)

// Persistencias
type Persistence struct {
	DbTestMSBeer *gorm.DB
}

// Iniciación db | repos
func NewPersistence() (*Persistence, error) {

	// TestMSBeer engine
	TestMSBeerEngine, errorConnection := db.NewTestMSBeerDB()
	if errorConnection != nil {
		log.Fatal(errorConnection)
	}

	return &Persistence{
		DbTestMSBeer: TestMSBeerEngine.DB,
	}, nil
}

// Close DB
func (p *Persistence) Close() error {
	posDB, errQ := p.DbTestMSBeer.DB()
	if errQ != nil {
		return errQ
	}
	errDbClose := posDB.Close()
	if errDbClose != nil {
		return errDbClose
	}

	return nil
}

// AutoMigración
func (p *Persistence) Automigrate() {
	p.DbTestMSBeer.AutoMigrate(test_ms_beer_entity.Beer{})

}
