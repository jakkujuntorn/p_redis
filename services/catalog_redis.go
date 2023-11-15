package services

import (
	"context"
	"encoding/json"
	"fmt"
	"goredis/repositories"
	"os"
	"time"

	"github.com/go-redis/redis/v9"
)

// ****** Adaptor  Redis******

type catalogServiceRedis struct {
	productRepo repositories.I_ProducrRepository
	redisClient *redis.Client
}

func NewCatalogServiceRedis(productRepo repositories.I_ProducrRepository, redisClient *redis.Client) I_CatalogService {
	return catalogServiceRedis{productRepo, redisClient}
}

func (s catalogServiceRedis) GetProducts() (products []ProductCatalog, err error) {

	// อ่านค่าจาก redis ก่อนถ้าได้ก็แสดงผลออกมา
	// อ้่านค่าจาก redis ไม่ได้ ให้ไปปอ่านจาก repository แทน ถ้าอ่านได้
	// ให้ set ค่าลง redis ด้วย

	// Redis Get
	// key := "service::GetProducts"
	key := os.Getenv("REDIS_KEY")

	// s.redisClient.Keys() // ใช้ทำอะไร

	// 1.1 เขียนแบบย่อ อ่านค่าจาก Redis
	//****** ถ้า err == nil(ไม่มี err) แสดงว่า อ่านค่าใน redis ได้
	if productsJson, err := s.redisClient.Get(context.Background(), key).Result(); err == nil {
		// 1.2ถ้า json err == nil แสดงว่ามีค่าจาก redis
		// 1.3 แปลงค่ามาเป็น struct ด้วย json
		if json.Unmarshal([]byte(productsJson), &products); err == nil {
			// แปลงข้อมูล จาก productsJson สู่ products
			fmt.Println(" ******Service Redis *******")
			return products, nil
		}
	}

	// 2.1 ถ้าอ่านจาก Redis ไม่ได้ให้ไปอ่านที่ repository แทน
	productsDB, err := s.productRepo.GetProducts()
	if err != nil {
		return nil, err
	}

	//2.2 ปั้นข้อมูลใหม่
	for _, p := range productsDB {
		products = append(products, ProductCatalog{
			ID:       p.ID,
			Name:     p.Name,
			Quantity: p.Quantity,
		})
	}

	// 3. Set Redis  แบบย่อ
	// ถ้า err == nil แสดงว่าแปลงได้
	if data, err := json.Marshal(products); err == nil {
		// json.Marshal retun []byte
		// ไม่ต้องจัดการ error ก็ได้ เพราะถ้า set ไม่ได้ มันจะไปอ่านจาก DB เอง
		// er := s.redisClient.Set(context.Background(), key, string(data), time.Second*10)
		// _ = er

		s.redisClient.Set(context.Background(), key, string(data), time.Second*10)
		// er.Err()
	}

	fmt.Println(" ****** Service Repository *******")
	return products, nil
}
