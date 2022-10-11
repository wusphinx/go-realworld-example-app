package singleton

import "sync"

var once sync.Once

type singleton struct {
}

var instance *singleton

// 用于测试
var count int

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
		count++
	})

	return instance
}
