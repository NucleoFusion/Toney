package enums

type TaskStatus int

const (
	Pending = iota
	Started
	Abandoned
	Complete
)
