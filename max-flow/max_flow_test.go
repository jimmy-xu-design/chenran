package max_flow

import (
	"sync/atomic"
	"testing"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

type Task struct {
	TaskID		string		`json:"taskID"`
	Cap			int32		`json:"capacity"`
}

type Resource struct {
	ResourceID	string		`json:"resourceID"`
	AllowTask	[]string	`json:"allow_task"`
	Cap			int32		`json:"capacity"`
}

type TestData struct {
	Task_list		[]Task		`json:"task_list"`
	Resource_list	[]Resource	`json:"resource_list"`
}

func (t *Task) ID() string {
	return t.TaskID
}

func (t *Task) Capacity() int {
	return int(atomic.LoadInt32(&t.Cap))
}

func (t *Task) Match(p Node) bool {
	resource := p.(*Resource)
	if resource == nil {
		return false
	}

	for _, v := range resource.AllowTask {
		if v == t.TaskID {
			return true
		}
	}

	return false
}

func (t *Task) AddFlow(num int) {
	atomic.AddInt32(&t.Cap, int32(-num))
}

func (r *Resource) ID() string {
	return r.ResourceID
}

func (r *Resource) Capacity() int {
	return int(atomic.LoadInt32(&r.Cap))
}

func (r *Resource) Match(p Node) bool {
	task := p.(*Task)

	if task == nil {
		return false
	}

	for _, v := range r.AllowTask {
		if v == task.TaskID {
			return true
		}
	}

	return false
}

func (r *Resource) AddFlow(num int) {
	atomic.AddInt32(&r.Cap, int32(-num))
}

func TestMaxMatch(t *testing.T) {
	bytes, err := ioutil.ReadFile("./data.json")
	if err != nil {
		fmt.Printf("read file %s failed\n", "data.json")
		return
	}

	var test_data TestData

	err = json.Unmarshal(bytes, &test_data)
	if err != nil {
		fmt.Printf("parse business config file failed, %v\n", err)
		return
	}

	var resource_list []Node
	for i, _ := range test_data.Resource_list {
		resource_list = append(resource_list, &test_data.Resource_list[i])
	}

	var task_list []Node
	for i, _ := range test_data.Task_list {
		task_list = append(task_list, &test_data.Task_list[i])
	}

	max_flow := MaxFlow{}

	flow_list := max_flow.Match(task_list, resource_list)
	for _, v := range flow_list {
		if v.Count > 0 {
			fmt.Println(v.From.ID() + "->" + v.To.ID())
		}
	}
}