package timer_task

import (
	"fmt"
	"time"
)

type Queue struct {
	tasks  *Task
	length int
}

func NewQueue() *Queue {
	return &Queue{
		tasks:  nil,
		length: 0,
	}
}

type Task struct {
	name    string // 任务名字
	exeTime int64  // 任务执行时间
	f       func() // 要执行的方法
	next    *Task
}

func NewTask(name string, et time.Time, f func()) *Task {
	return &Task{
		name:    name,
		exeTime: et.Unix(),
		f:       f,
		next:    nil,
	}
}

func (q *Queue) Insert(qt *Task) {
	if q.length == 0 {
		q.tasks = qt
	} else {
		cur := q.tasks
		var curPre *Task = nil
		for {
			if cur.exeTime < qt.exeTime {
				if cur.next == nil {
					cur.next = qt
					break
				} else {
					tmp := cur
					cur = tmp.next
					curPre = tmp
				}
			} else {
				qt.next = cur
				if curPre == nil {
					q.tasks = qt
				} else {
					curPre.next = qt
				}
				break
			}
		}
	}
	q.length++
}

func (q *Queue) GetAll() {
	if q.length == 0 {
		return
	}
	cur := q.tasks
	for {
		if cur == nil {
			break
		}
		fmt.Println(cur.name)
		fmt.Println(cur.exeTime)
		cur = cur.next
	}
	fmt.Println(q.length)
}

func (q *Queue) Pop() *Task {
	if q.length == 0 {
		return nil
	}
	o := q.tasks
	q.tasks = o.next
	o.next = nil
	q.length--
	return o
}
