package timer_task

import (
	"errors"
	"time"
)

type Cronjob struct {
	queue *Queue
}

func NewCronjob() *Cronjob {
	return &Cronjob{
		queue: NewQueue(),
	}
}

func (c *Cronjob) AddTask(name string, et time.Time, f func()) error {
	if et.Unix() < time.Now().Unix() {
		return errors.New("Param error:et")
	}
	if f == nil {
		return errors.New("Param error:f")
	}

	task := NewTask(name, et, f)
	c.queue.Insert(task)
	return nil
}

func (c *Cronjob) Do() {
	for {
		if c.queue.tasks == nil {
			continue
		}
		if time.Now().Unix() > c.queue.tasks.exeTime {
			task := c.queue.Pop()
			go task.f()
		}
	}
}
