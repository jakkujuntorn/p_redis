package repositories

import (
	"gorm.io/gorm"
)

// ใช้ตัวเล็กเพราะ จะไม่ใครเข้าถึงโดยตรง เดียวจะไม่ใส db เข้ามา ให้ใช้ผ่าน New
type productRepositoryDB struct {
	db *gorm.DB
}

func NewProductRepositoryDB(dbGorm *gorm.DB) I_ProducrRepository {

	// dbGorm.AutoMigrate(&product{})
	// ****** mock data *****
	// mockData(db)

	return productRepositoryDB{db: dbGorm}
}

// struct comfrom ตาม I_ProducrRepository
func (r productRepositoryDB) GetProducts() (products []product, err error) {
	// err ถูกประกาศด้านบนแล้ว
	// r.db.Find(&products) // แบบนี้จะได้ข้อมูลทั้งหมด ไม่ควจเพราะอาจ คอมค้างถ้ามีเยอะมาก
	// err = r.db.Order("quantity desc").Limit(30).Find(&products).Error  // แบบนี้จะได้ err ออกไปเลย *****

	// เรียงลำดับ และ ลิมิต
	// tx := r.db.Order("quantity desc").Limit(30).Find(&products)

	// ใช้แบบนี้ก็ได้ เพราะ ประกาศ err ไว้ด้านบนแล้ว แล้ว หลัง DB มี .Error(มันได้ค่า error)
	err = r.db.Order("quantity desc").Limit(30).Find(&products).Error

	// ****** ตัวแปรถูกสร้างจากด้านบนแล้ว เลยไม่ต้องสร้างใหม่ ******
	// return products, tx.Error
	
	return products, err
}
