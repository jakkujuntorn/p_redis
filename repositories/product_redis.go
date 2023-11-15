package repositories

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
)

type productRepositoryRedis struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewProductRepositoryRedis(db *gorm.DB, redisClient *redis.Client) I_ProducrRepository {
	// db.AutoMigrate(&product{})
	// mockData(db)

	return productRepositoryRedis{db, redisClient}
}

func (r productRepositoryRedis) GetProducts() (products []product, err error) {

	key := "repostitory::GetProducts"

	//*************************** redis *************************
	// Redis Get เพื่อเช็คว่ามีค่าใน  redis หรือไม่
	productsJson, err := r.redisClient.Get(context.Background(), key).Result()
	//****** ถ้า err == nil(ไม่มี err) แสดงว่า อ่านใน redis ได้  ถ้ามี err จะผ่านไปอ่านที่ Database  ******
	if err == nil {
		// ******** อ่านค่าใน redis ได้ *********
		// productsJson คือค่าที่ได้จาก redis *******
		// ******** แปลงค่า **********
		err = json.Unmarshal([]byte(productsJson), &products)
		// Unmarshal ไม่มี err แสดงว่าแปลงค่าได้
		if err == nil {
			fmt.Println("********* Form Redis **********")
			return products, nil
		}
	}

	// ********************* DB gorm ***************************
	err = r.db.Order("quantity desc").Limit(30).Find(&products).Error
	if err != nil {
		return nil, err
	}

	// ข้อมูลมา
	// fmt.Println(products)

	// แปลงข้อมูล จาก DB เป็น json
	// เพื่อเอาไปใช้ในการ Set Redis *******
	data, err := json.Marshal(products)
	if err != nil {
		// return errors.New("Cannot Set Marshal") // แบบนี้ขึ้นเหลือง
		// fmt.Errorf("Cannot Set Marshal")
		// message := fmt.Fscanf("Cannot Set Marshal")
		return nil, errors.New("Cannot Set Marshal")
	}

	

	// ข้อมูลมา
	// fmt.Println(string(data))

	//ใสค่าใน Redis / Redis Set  (context.Background(),คีย์,value (แปลงจาก []byte เป็น string),ระยะเวลา(ถ้าไม่ให้มดอายุใส 0) )
	// ตรง ข้อมูลเอา products ที่ได้จาก DB ลงไปเลยไม่ได้เหรอ **********************
	// ตรง data เอา products  เอามาใสเลยไม่ได้เหรอ
	errSetRedis := r.redisClient.Set(context.Background(), key, string(data), time.Second*18).Err()
	if errSetRedis != nil {
		return nil, errSetRedis
	}

	fmt.Println(" ******** Form Database **********")
	return products, nil
}
