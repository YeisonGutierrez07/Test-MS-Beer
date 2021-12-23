package main

import (
	"log"
	"os"
	"test_ms_beer/infrastructure/interfaces"
	"test_ms_beer/infrastructure/interfaces/middleware"
	"test_ms_beer/infrastructure/persistence/base"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// en este caso incluire en el repo el .env pero en produccion no es recomendable, las variables deben estar directamente seteadas en el entorno (aws, heroku etc.)
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}
}

func main() {

	// Iniciamos los repositorios - conexion a la DB
	p, err := base.NewPersistence()
	if err != nil {
		panic(err)
	}

	// Defer para cerrar conexiones postConexion
	defer p.Close()

	// Automigraciones necesarias
	p.Automigrate()

	// Iniciamos las interfaces que usaremos en los endpoints
	beerServices := interfaces.NewBeer(p)

	// Configuracion gin por default
	r := gin.Default()
	// mobile services
	beers := r.Group("/beers")
	beers.Use(middleware.CORSMiddleware()) //CORS

	////Routes

	// Servicio para Lista todas las cervezas
	beers.GET("", beerServices.GetAllBeers)
	// Servicio para crear una nueva cerveza
	beers.POST("", beerServices.CreateNewBeer)
	// Servicio para buscar una cerveza por ID
	beers.GET("/:beerID", beerServices.GetBeerByID)
	// Servicio para consultar el precio de una caja de cervaza
	beers.GET("/:beerID/boxprice", beerServices.GetBoxPriceBeers)

	//Starting the application
	appPort := os.Getenv("PORT")
	if appPort == "" {
		appPort = "5000" //localhost
	}
	log.Fatal(r.Run(":" + appPort))
}
