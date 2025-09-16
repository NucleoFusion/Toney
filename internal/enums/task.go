package enums

type TaskStatus int

var TaskStatusMap map[TaskStatus]string = map[TaskStatus]string{
	Pending:   "Pending",
	Abandoned: "Abandoned",
	Complete:  "Complete",
	Started:   "Started",
}

const (
	Pending = iota
	Started
	Abandoned
	Complete
)
