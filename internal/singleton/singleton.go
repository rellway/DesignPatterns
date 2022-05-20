package singleton

import (
	"fmt"
	"sync"
)

var (
	once     sync.Once
	instance *Singleton // instance需要首字母小写，保证其他模块无法引用
)

type Singleton struct{}

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
		fmt.Println("first create instance")
	})
	fmt.Println("instance created")
	return instance
}
