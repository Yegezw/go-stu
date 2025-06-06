package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func test3() {
	fmt.Println("----------------------test3----------------------")

	c := &SafeCounter{v: make(map[string]int)}

	// 1000 Go 程并发写
	for i := 0; i < 1000; i++ {
		go c.Inc("apple")
	}
	time.Sleep(10 * time.Millisecond)
	fmt.Println(c.Value("apple"))

	// 等待所有 goroutine 完成
	var group sync.WaitGroup
	group.Add(2)
	go myGo(&group)
	go myGo(&group)
	group.Wait()

	// 获取单例配置
	var config *Config
	config = GetConfig()
	config = GetConfig()
	fmt.Println(config)

	// 并发安全 Map
	m := sync.Map{}
	m.Store("name", "zs")
	m.Store("age", 18)

	value, _ := m.Load("age")
	fmt.Println("age:", value.(int))

	m.Range(
		func(k, v any) bool {
			fmt.Println("key:", k, "value:", v)
			return true // 返回 true 继续迭代
		},
	)

	m.Delete("name")
	value, ok := m.Load("name")
	fmt.Println("value:", value, "ok:", ok)

	actual, loaded := m.LoadOrStore("name", "zs")
	fmt.Println("actual:", actual, "loaded:", loaded)

	// atomic 包的使用
	num := int32(0)
	group = sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			atomic.AddInt32(&num, 1)
		}()
	}
	group.Wait()
	fmt.Println("num:", num)

	// 支持任意类型进行原子操作
	var v atomic.Value

	v.Store("Hello")
	fmt.Println(v.Load())

	v.Swap("World")
	fmt.Println(v.Load())

	v.CompareAndSwap("World", "Hello, World!")
	fmt.Println(v.Load())

	// 对象池 sync.Pool
	poolDemo()
}

// ----------------------------------------------------------

// SafeCounter 是并发安全的
type SafeCounter struct {
	m sync.Mutex   // 互斥锁
	l sync.RWMutex // 读写锁
	v map[string]int
}

// Inc 会执行 v[key]++
func (c *SafeCounter) Inc(key string) {
	c.m.Lock()
	c.v[key]++
	c.m.Unlock()
}

// Value 返回 c.v[key]
func (c *SafeCounter) Value(key string) int {
	c.m.Lock()
	defer c.m.Unlock()
	return c.v[key]
}

// ----------------------------------------------------------

func myGo(group *sync.WaitGroup) {
	defer group.Done()
	time.Sleep(10 * time.Millisecond)
	fmt.Println("myGo")
}

// ----------------------------------------------------------

type Config struct{}

var once sync.Once
var instance *Config

// GetConfig 返回单例配置
func GetConfig() *Config {
	// 确保只执行一次
	once.Do(
		func() {
			instance = &Config{}
			fmt.Println("Config instance created")
		},
	)
	return instance
}

// ----------------------------------------------------------

type Student struct {
	Name string
	Age  int
}

func (s *Student) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", s.Name, s.Age)
}

func poolDemo() {
	var student *Student
	pool := &sync.Pool{
		New: func() any {
			return &Student{
				Name: "",
				Age:  -1,
			}
		},
	}

	// 从池中获取对象
	student = pool.Get().(*Student)
	fmt.Printf("Address of student: %p -> {%v}\n", student, student)

	// 池化 student 对象
	pool.Put(student)
	student.Name = "zs"
	student.Age = 18
	student = pool.Get().(*Student)
	fmt.Printf("Address of student: %p -> {%v}\n", student, student)
}
