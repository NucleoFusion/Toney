package daily

import (
	"slices"

	"github.com/SourcewareLab/Toney/internal/enums"
	"github.com/SourcewareLab/Toney/internal/messages"
)

func (m Daily) CreateTask(msg messages.TaskPopupMessage, isUnique bool) {
	task := Task{
		TaskTitle: msg.Title,
		TaskDesc:  msg.Desc,
		Status:    msg.Status,
	}

	if isUnique {
		m.Tasks.Unique = append(m.Tasks.Unique, task) // Seperate Task Input for recurring / unique tasks
	} else {
		m.Tasks.Recurring = append(m.Tasks.Recurring, task) // Seperate Task Input for recurring / unique tasks
	}
	WriteItems(m.Tasks)
}

func (m Daily) DeleteTask(msg messages.TaskPopupMessage) {
	if !msg.IsDeleted {
		return
	}

	item := m.List.SelectedItem()

	switch task := item.(type) {
	case Task:
		switch task.TaskType {
		case enums.RecurringTask:
			m.Tasks.Recurring = slices.Delete(m.Tasks.Recurring, task.Index, task.Index+1)
		case enums.UniqueTask:
			m.Tasks.Unique = slices.Delete(m.Tasks.Unique, task.Index, task.Index+1)
		}
	case TodoTask:
		m.Tasks.Todo = slices.Delete(m.Tasks.Todo, task.Index, task.Index+1)
	}

	WriteItems(m.Tasks)
}

func (m Daily) StatusChangeTask(msg messages.TaskPopupMessage) {
	item := m.List.SelectedItem()

	switch task := item.(type) {
	case Task:
		switch task.TaskType {
		case enums.RecurringTask:
			m.Tasks.Recurring[task.Index].Status = msg.Status
		case enums.UniqueTask:
			m.Tasks.Unique[task.Index].Status = msg.Status
		}
	case TodoTask:
		m.Tasks.Todo[task.Index].Status = msg.Status
	}

	WriteItems(m.Tasks)
}

func (m Daily) EditTask(msg messages.TaskPopupMessage) {
	item := m.List.SelectedItem()

	switch task := item.(type) {
	case Task:
		switch task.TaskType {
		case enums.RecurringTask:
			v := &m.Tasks.Recurring[task.Index]
			v.Status = msg.Status
			v.TaskTitle = msg.Title
			v.TaskDesc = msg.Desc
		case enums.UniqueTask:
			v := &m.Tasks.Unique[task.Index]
			v.Status = msg.Status
			v.TaskTitle = msg.Title
			v.TaskDesc = msg.Desc
		}
	case TodoTask: // Cant Edit Title
		v := &m.Tasks.Todo[task.Index]
		v.Status = msg.Status
		v.Text = msg.Desc
	}

	WriteItems(m.Tasks)
}
