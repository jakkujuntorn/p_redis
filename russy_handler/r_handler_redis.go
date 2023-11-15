package russyhandler

import (
	"context"
	"encoding/json"
	"goredis/russy_services"
	_ "goredis/services"
	"time"

	"fmt"

	"github.com/go-redis/redis/v9"
)

type RedisData struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type I_RedisHandler interface {
	GetDataRedis() ([]RedisData, error)
}

type redisProductHandler struct {
	productService russyservices.IRus_service
	redisClient    *redis.Client
}

// func New paramiter  ต้องเรียงตาม struct ด้วยว่าอะรก่อน และหลัง ***
func NewHandler_Russy_redis( productService russyservices.IRus_service, redisClient *redis.Client) I_RedisHandler {
	return redisProductHandler{productService,redisClient}
}

// GetDataRedis implements I_RedisHandler
func (r redisProductHandler) GetDataRedis() (products []RedisData, err error) {
	key := "repostitory::GetProducts"

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

	productService, errGetService := r.productService.Custom_GetProduct()
	if errGetService != nil {
		return nil, errGetService
	}

	for _, p := range productService {
		products = append(products, RedisData{
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
