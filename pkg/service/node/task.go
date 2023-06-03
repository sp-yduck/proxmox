package node

import (
	"fmt"
	"time"

	"github.com/pkg/errors"

	"github.com/sp-yduck/proxmox/pkg/api"
)

// EnsureTaskDone waits/gets task and check if task status is ok
func (c *Node) EnsureTaskDone(upid string) error {
	task, err := c.WaitTask(upid)
	if err != nil {
		return err
	}
	if !task.IsStatusOK() {
		return errors.New(task.Status)
	}
	return nil
}

func (c *Node) Tasks() ([]*Task, error) {
	var tasks []*Task
	if err := c.Client.Get(fmt.Sprintf("/nodes/%s/tasks", c.Node), &tasks); err != nil {
		return nil, err
	}
	for _, t := range tasks {
		t.Client = c.Client
	}
	return tasks, nil
}

func (c *Node) Task(upid string) (*Task, error) {
	var tasks []*Task
	if err := c.Client.Get(fmt.Sprintf("/nodes/%s/tasks", c.Node), &tasks); err != nil {
		return nil, err
	}
	for _, t := range tasks {
		if t.UPID == upid {
			t.Client = c.Client
			return t, nil
		}
	}
	return nil, api.ErrNotFound
}

func (c *Node) WaitTask(upid string) (*Task, error) {
	fmt.Println(upid)
	for i := 0; i < 10; i++ {
		task, err := c.Task(upid)
		if api.IsNotFound(err) {
			time.Sleep(time.Second * 1)
			continue
		}
		return task, err
	}
	return nil, errors.New("task wait deadline exceeded")
}

func (c *Task) IsStatusOK() bool {
	return c.Status == "OK"
}
