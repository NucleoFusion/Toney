package enums

type TaskType int

var TaskTypeMap map[TaskType]string = map[TaskType]string{
	RecurringTask: "Recurring",
	UniqueTask:    "Unique",
}

const (
	UniqueTask = iota
	RecurringTask
)
