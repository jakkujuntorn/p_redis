package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"goredis/services"

	"time"

	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"
)

type catalogHandlerRedis struct {
	catalogSrv  services.I_CatalogService
	redisClient *redis.Client
}

func NewCatalogHandlerRedis(catalogSrv services.I_CatalogService, redisClient *redis.Client) ICatalogHandler {
	return catalogHandlerRedis{catalogSrv, redisClient}
}

func (h catalogHandlerRedis) GetProducts(c *fiber.Ctx) error {
	// 1 อ่านค่าจาก redis
	// 2 อ่นค่าจาก DB
	// 3 set ค่า ให้ redis
	keyRedis := "handler::GetProducts"

	// Redis Get และ เช็ค error
	if responseJson, err := h.redisClient.Get(context.Background(), keyRedis).Result(); err == nil {

		fmt.Println("Redis")

		// ข้อมูลที่ได้เป็น json อยู่แล้วเพราะ service แปลงมาให้แล้ว
		// ข้อมูลเป็น json อยู่แล้ว เลย return ออกไปได้เลย
		// แต่ต้องเซต ่header ใหม่ ไม่งั้นจะเป็น text plan ********
		c.Set("Content-Type", "application/json")
		return c.SendString(responseJson)
	}

	// ถ้าอ่านจาก Redis ไม่ได้ให้อ่านจาก Services แทน
	products, err := h.catalogSrv.GetProducts()
	if err != nil {
		return err
	}

	response := fiber.Map{
		"status":   "ok",
		"products": products,
	}

	// Redis Set
	// ส่ง response  เข้าไปเลย เพราะต้องการ status ด้วย ใช้ตอน return
 	if data, err := json.Marshal(response); err == nil {
		h.redisClient.Set(context.Background(), keyRedis, string(data), time.Second*10)
	}

	fmt.Println("Data Base")

	return c.JSON(response)
}
