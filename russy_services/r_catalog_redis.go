package russyservices

import (
	"context"
	"encoding/json"
	"goredis/repositories"
	"os"
	"time"

	"fmt"

	"github.com/go-redis/redis/v9"
)

type RedisProduct struct {
	ID       int
	Name     string
	Quantity int
}

type I_Redis_Service interface {
	GetRedisData() ([]RedisProduct, error)
}

type productServiceRedis struct {
	productRepo repositories.I_ProducrRepository
	redisClient *redis.Client
}

func NewR_Catalog_Service_Redis(redisRepo repositories.I_ProducrRepository, redisClient *redis.Client) I_Redis_Service {
	return productServiceRedis{redisRepo, redisClient}
}

// GetRedisData implements I_Redis
func (r productServiceRedis) GetRedisData() (products []RedisProduct, err error) {
	// key := "repostitory::GetProducts"
	key:=os.Getenv("REDIS_KEY")

	// Redis Get เพื่อเช็คว่ามีค่าใน  redis หรือไม่
	productsJson, err := r.redisClient.Get(context.Background(), key).Result()
	//****** ถ้า err == nil(ไม่มี err) แสดงว่า อ่านใน redis ได้  ถ้ามี err จะผ่านไปอ่านที่ Database  ******
	if err == nil {
		// ******** อ่านค่าใน redis ได้ *****
		// ******** แปลงค่า **********
		err = json.Unmarshal([]byte(productsJson), &products)
		// Unmarshal ไม่มี err
		if err == nil {
			fmt.Println("********* Form Redis **********")
			return products, nil
		}
	}

	productRepo, errGetRepor := r.productRepo.GetProducts()
	if errGetRepor != nil {
		return nil,errGetRepor
	}

	
	for _,p := range productRepo {
		products = append(products, RedisProduct{
			ID:       p.ID,
			Name:     p.Name,
			Quantity: p.Quantity,
		})
	}

	// Set Redis  แบบย่อ
	if data, err := json.Marshal(products); err == nil {
		r.redisClient.Set(context.Background(), key, string(data), time.Second*30)
	}

	fmt.Println(" ******** Form Database **********")
	return products, nil
}
