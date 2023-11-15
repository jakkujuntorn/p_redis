package main

import (
	"goredis/handler_error"
	"goredis/handlers"
	"goredis/repositories"
	"goredis/russy_handler"
	"goredis/russy_services"
	"goredis/services"
	"log"
	"os"

	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

func main() {

	// ****** config ********
	errEnv := godotenv.Load("config.env")
	if errEnv != nil {
		log.Println("please conside environment varable : %s", errEnv)
	}

	// ****** Data base MySql  ******
	db := initDatabase()

	// ****** Redis ******
	redisClient := initRedis()

	// ************* Repo ***********
	productRepo := repositories.NewProductRepositoryDB(db)

	// **** ทำเอง น่าจะได้ ค่าของ products *****
	// ee,_:=repositories.I_ProducrRepository.GetProducts(productRepo)

	// productRepo := repositories.NewProductRepositoryRedis(db, redisClient)

	// ************* Service ****************
	// productService := services.NewCatalogService(productRepo)
	productService := services.NewCatalogServiceRedis(productRepo, redisClient)
	// ***************          ***************

	// ************* Service ทำเอง ***********
	//  DB
	customProductService := russyservices.NewR_Catalog_Service(productRepo)
	_ = customProductService
	// Redis
	customProductRedis := russyservices.NewR_Catalog_Service_Redis(productRepo, redisClient)
	// dataRedis,_:=customProductRedis.GetRedisData()
	_ = customProductRedis

	//***************** Handler ****************
	// productHandler := handlers.NewCatalogHandler(productService)
	productHandler := handlers.NewCatalogHandlerRedis(productService, redisClient)
	_ = productHandler
	// productHandler.GetProducts()

	// ************* Handler ทำเอง ***********
	// Redis
	productHandlerRussyRedis := russyhandler.NewHandler_Russy_redis(customProductService, redisClient)
	// fmt.Println(productHandlerRussy.GetDataRedis())
	_ = productHandlerRussyRedis

	// DB
	productHandlerRussy := russyhandler.NewHandler_Russy(customProductRedis)
	// fmt.Println(productHandlerRussy.GetDataHandler())

	// ********* fiber ***************
	fiberPort := os.Getenv("FIBER_PORT")

	app := fiber.New()
	app.Get("/products", productHandler.GetProducts)
	app.Get("/productsrussy", productHandlerRussy.GetDataHandler)
	app.Get("/testerro", handlererror.Error1)
	app.Listen(fiberPort)

}

func initDatabase() *gorm.DB {
	dial := mysql.Open("root:P@ssw0rd@tcp(127.0.0.1:3306)/redis")
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func initRedis() *redis.Client {
	addr := os.Getenv("REDIS_ADDR")
	return redis.NewClient(&redis.Options{
		Addr: addr,
	})
}
