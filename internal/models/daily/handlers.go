package daily

import (
	"slices"

	"github.com/SourcewareLab/Toney/internal/enums"
	"github.com/SourcewareLab/Toney/internal/messages"
)

func (m Daily) CreateTask(msg messages.TaskPopupMessage) {
	task := Task{
		TaskTitle: msg.Title,
		TaskDesc:  msg.Desc,
		Status:    msg.Status,
	}

	m.Tasks.Unique = append(m.Tasks.Unique, task) // Seperate Task Input for recurring / unique tasks

	WriteItems(m.Tasks)
}

func (m Daily) DeleteTask(msg messages.TaskPopupMessage) {
	if !msg.IsDeleted {
		return
	}

	item := m.List.SelectedItem()

	task, ok := item.(Task)
	if !ok { // Making sure that item is of type Task
		return
	}

	switch task.TaskType {
	case enums.RecurringTask:
		m.Tasks.Recurring = slices.Delete(m.Tasks.Recurring, task.Index, task.Index+1)
	case enums.UniqueTask:
		m.Tasks.Unique = slices.Delete(m.Tasks.Unique, task.Index, task.Index+1)
	}

	WriteItems(m.Tasks)
}

func (m Daily) StatusChangeTask(msg messages.TaskPopupMessage) {
	item := m.List.SelectedItem()

	task, ok := item.(Task)
	if !ok { // Making sure that item is of type Task
		return
	}

	task.Status = msg.Status

	switch task.TaskType {
	case enums.RecurringTask:
		m.Tasks.Recurring[task.Index] = task
	case enums.UniqueTask:
		m.Tasks.Unique[task.Index] = task
	}

	WriteItems(m.Tasks)
}

func (m Daily) EditTask(msg messages.TaskPopupMessage) {
	task := Task{
		TaskTitle: msg.Title,
		TaskDesc:  msg.Desc,
		Status:    msg.Status,
	}

	item := m.List.SelectedItem()

	oldTask, ok := item.(Task)
	if !ok { // Making sure that item is of type Task
		return
	}

	task.Status = oldTask.Status

	switch oldTask.TaskType {
	case enums.RecurringTask:
		m.Tasks.Recurring[oldTask.Index] = task
	case enums.UniqueTask:
		m.Tasks.Unique[oldTask.Index] = task
	}

	WriteItems(m.Tasks)
}
