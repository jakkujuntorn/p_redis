package repositories

import (
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type product struct {
	ID       int
	Name     string
	Quantity int
}

type I_ProducrRepository interface {
	GetProducts() ([]product, error)
}


// ***** Mock Data ****
func mockData(db *gorm.DB) error {
	// เช็คก่อนว่ามีข้อมูลรึยัง
	var count int64
	db.Model(&product{}).Count(&count)
	if count > 0 {
		return nil
	}

	products := []product{}
	// ทำ random number
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	// ใส Mock Data 5000 record
	for i := 0; i < 5000; i++ {
		products = append(products, product{
			Name:     fmt.Sprintf("Product%v", i+1),
			Quantity: random.Intn(100),
		})
	}

	// return DB ที่สร้างเสร็จแล้วออกมา
	// create หลาย record ได้ ให้สร้างเป็น slice แล้วส่งเข้าไปที่เดียวเลย
	return db.Create(&products).Error 
}