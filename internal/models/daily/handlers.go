package daily

import (
	"slices"

	"github.com/SourcewareLab/Toney/v2/internal/enums"
	"github.com/SourcewareLab/Toney/v2/internal/messages"
)

func (m Daily) CreateTask(msg messages.TaskPopupMessage, isUnique bool) {
	task := Task{
		TaskTitle: msg.Title,
		TaskDesc:  msg.Desc,
		Status:    msg.Status,
	}

	if isUnique {
		task.TaskType = enums.UniqueTask
	} else {
		task.TaskType = enums.RecurringTask
	}

	m.Tasks = append(m.Tasks, task)

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

	m.Tasks = slices.Delete(m.Tasks, task.ID, task.ID+1)

	WriteItems(m.Tasks)
}

func (m Daily) StatusChangeTask(msg messages.TaskPopupMessage) {
	item := m.List.SelectedItem()

	task, ok := item.(Task)
	if !ok { // Making sure that item is of type Task
		return
	}

	task.Status = msg.Status

	m.Tasks[task.ID] = task

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

	m.Tasks[oldTask.ID] = task

	WriteItems(m.Tasks)
}
