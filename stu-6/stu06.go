package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func test6() {
	pool := NewPool(3, 10)

	for i := 0; i < 10; i++ {
		i := i
		task := NewTask(
			func() {
				fmt.Println("Task executed", i)
			},
		)
		pool.AddTask(task)
	}

	fmt.Println(pool.GetCurRoutines())
	time.Sleep(time.Second)
}

// ----------------------------------------------------------

type Task struct {
	run func()
}

func NewTask(run func()) *Task {
	return &Task{run: run}
}

type Pool struct {
	maxRoutines int64      // 最大协程数
	curRoutines int64      // 当前协程数
	tasks       chan *Task // 任务队列
	lock        sync.Mutex // 互斥锁
}

func NewPool(maxRoutines int64, maxTasks int64) *Pool {
	return &Pool{
		maxRoutines: maxRoutines,
		curRoutines: 0,
		tasks:       make(chan *Task, maxTasks),
	}
}

func (pool *Pool) GetMaxRoutines() int64 {
	return pool.maxRoutines
}

func (pool *Pool) GetCurRoutines() int64 {
	return atomic.LoadInt64(&pool.curRoutines)
}

func (pool *Pool) AddTask(task *Task) {
	pool.tasks <- task

	pool.lock.Lock()
	defer pool.lock.Unlock()
	if pool.GetCurRoutines() < pool.GetMaxRoutines() {
		pool.run()
	}
}

func (pool *Pool) incRoutine() {
	atomic.AddInt64(&pool.curRoutines, 1)
}

func (pool *Pool) decRoutine() {
	atomic.AddInt64(&pool.curRoutines, -1)
}

func (pool *Pool) run() {
	pool.incRoutine()
	go func() {
		defer pool.decRoutine()
		for task := range pool.tasks {
			task.run()
		}
	}()
}
