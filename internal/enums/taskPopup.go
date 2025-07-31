package enums

type TaskPopup int

const (
	CreateUnique = iota
	CreateRecurring
	Delete
	Edit
	ChangeStatus
	ClosePopup
)
