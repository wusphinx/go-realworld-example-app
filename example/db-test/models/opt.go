package models

import "gorm.io/gorm"

type ListProductsOptions struct {
	limit  int
	offset int
	code   string
	price  uint
}

type Option interface {
	applyTo(*ListProductsOptions)
}

type WithCode string

func (r WithCode) applyTo(opts *ListProductsOptions) {
	opts.code = string(r)
}

type WithPrice uint

func (r WithPrice) applyTo(opts *ListProductsOptions) {
	opts.price = uint(r)
}

type WithPager struct {
	limit  int
	offset int
}

func (r WithPager) applyTo(opts *ListProductsOptions) {
	opts.limit = r.limit
	opts.offset = r.offset
}

func GetProducts(db *gorm.DB, opts ...Option) ([]Product, error) {
	var products []Product
	config := ListProductsOptions{limit: 10, offset: 0, code: "", price: 0}
	for _, opt := range opts {
		opt.applyTo(&config)
	}

	//TODO: filter with config
	return products, nil
}
