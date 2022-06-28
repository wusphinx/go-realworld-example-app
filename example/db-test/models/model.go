package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func (p *Product) TableName() string {
	return "products"
}

func (p *Product) GetByID(db *gorm.DB, id int) (*Product, error) {
	var pt Product
	if err := db.First(&pt, id).Error; err != nil {
		return nil, err
	}

	return &pt, nil
}
