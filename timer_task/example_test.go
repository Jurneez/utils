package timer_task

import (
	"fmt"
	"testing"
	"time"
)

/*
var (
	// 5s之后开始
	task1 = NewTask("猫叫", time.Now().Add(time.Second*5), func() {
		fmt.Println("猫咪开始叫。。。。")
		time.Sleep(10 * time.Second)
		fmt.Println("猫咪叫了10s。。。。")
	})
	// 10s之后开始
	task2 = NewTask("狗叫", time.Now().Add(time.Second*10), func() {
		fmt.Println("狗狗开始叫。。。。")
		time.Sleep(10 * time.Second)
		fmt.Println("狗狗叫了10s。。。。")
	})
	// 10s之后开始
	task21 = NewTask("二号狗叫", time.Now().Add(time.Second*10), func() {
		fmt.Println("二号狗狗也开始叫。。。。")
		time.Sleep(10 * time.Second)
		fmt.Println("二号狗狗叫了10s。。。。")
	})
	// 2s之后开始
	task3 = NewTask("鸡叫", time.Now().Add(time.Second*2), func() {
		fmt.Println("小鸡开始叫。。。。")
		time.Sleep(5 * time.Second)
		fmt.Println("小鸡叫了5s。。。。")
	})
	// 7s之后开始
	task4 = NewTask("吃鱼", time.Now().Add(time.Second*7), func() {
		fmt.Println("我要开始吃鱼了。。。。")
		time.Sleep(8 * time.Second)
		fmt.Println("吃鱼用了8s。。。。")
	})
	// 9s之后开始
	task5 = NewTask("鸭子", time.Now().Add(time.Second*9), func() {
		fmt.Println("鸭子开始叫。。。。")
		time.Sleep(5 * time.Second)
		fmt.Println("鸭子叫了5s。。。。")
	})
	// 20s之后开始
	task6 = NewTask("鸽子", time.Now().Add(time.Second*20), func() {
		fmt.Println("鸽子飞走了")
		time.Sleep(1 * time.Second)
		fmt.Println("鸽子飞走用了1s。。。。")
	})

	taskQueue = NewQueue()
)

func Test_insert(t *testing.T) {
	taskQueue.Insert(task1)
	taskQueue.Insert(task2)
	taskQueue.Insert(task21)
	taskQueue.Insert(task3)
	taskQueue.Insert(task4)
	taskQueue.Insert(task5)
	taskQueue.Insert(task6)
	taskQueue.GetAll()
}

func Test_Pop(t *testing.T) {
	o := taskQueue.Pop()
	taskQueue.GetAll()

	o.f()
}
*/

func Test_Conjob(*testing.T) {
	cronjob := NewCronjob()
	// 5s之后开始
	cronjob.AddTask("猫叫", time.Now().Add(time.Second*5), func() {
		fmt.Println("猫咪开始叫。。。。")
		time.Sleep(10 * time.Second)
		fmt.Println("猫咪叫了10s。。。。")
	})
	// 10s之后开始
	cronjob.AddTask("狗叫", time.Now().Add(time.Second*10), func() {
		fmt.Println("狗狗开始叫。。。。")
		time.Sleep(10 * time.Second)
		fmt.Println("狗狗叫了10s。。。。")
	})
	// 10s之后开始
	cronjob.AddTask("二号狗叫", time.Now().Add(time.Second*10), func() {
		fmt.Println("二号狗狗也开始叫。。。。")
		time.Sleep(10 * time.Second)
		fmt.Println("二号狗狗叫了10s。。。。")
	})
	// 2s之后开始
	cronjob.AddTask("鸡叫", time.Now().Add(time.Second*2), func() {
		fmt.Println("小鸡开始叫。。。。")
		time.Sleep(5 * time.Second)
		fmt.Println("小鸡叫了5s。。。。")
	})
	// 7s之后开始
	cronjob.AddTask("吃鱼", time.Now().Add(time.Second*7), func() {
		fmt.Println("我要开始吃鱼了。。。。")
		time.Sleep(8 * time.Second)
		fmt.Println("吃鱼用了8s。。。。")
	})
	// 9s之后开始
	cronjob.AddTask("鸭子", time.Now().Add(time.Second*9), func() {
		fmt.Println("鸭子开始叫。。。。")
		time.Sleep(5 * time.Second)
		fmt.Println("鸭子叫了5s。。。。")
	})
	// 20s之后开始
	cronjob.AddTask("鸽子", time.Now().Add(time.Second*20), func() {
		fmt.Println("鸽子飞走了")
		time.Sleep(1 * time.Second)
		fmt.Println("鸽子飞走用了1s。。。。")
	})
	cronjob.Do()
	time.Sleep(time.Minute)
}

func Test_Conjob2(*testing.T) {
	cron := NewCronjob()
	go cron.Do()
	cron.AddTask("闹铃", time.Now().Add(time.Second*1), func() {
		fmt.Println("闹铃")
	})
	cron.AddTask("闹铃e", time.Now().Add(time.Second*3), func() {
		fmt.Println("闹铃e")
	})

	time.Sleep(time.Second * 5)
}
