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

	m.Tasks.Unique = append(m.Tasks.Unique, task)

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
}

func (m Daily) EditTask(msg messages.TaskPopupMessage) {
}
