package singleton

import "sync"

var once sync.Once

type singleton struct {
}

var instance *singleton

// 用于测试
var count int

//GetInstance return singleton
//revive:disable:unexported-return
func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
		count++
	})
	return instance
}
