package messages

import (
	"github.com/SourcewareLab/Toney/v2/internal/enums"
	filetree "github.com/SourcewareLab/Toney/v2/internal/fileTree"
)

type (
	ErrorMsg struct {
		Title string
		Msg   string
		Locn  string
	}

	CloseError struct{}

	FzfSelection struct {
		Selection string
		Exited    bool
	}

	TaskPopupMessage struct {
		Type      enums.TaskPopup
		IsDeleted bool
		Title     string
		Desc      string
		Status    enums.TaskStatus
	}

	ChangePage struct {
		Page enums.Page
	}

	ShowLoader struct{}

	HideLoader struct{}

	EditorClose struct {
		Err      error
		IsNative bool
		Value    string
	}

	ShowPopupMessage struct {
		Type enums.PopupType
		Curr *filetree.Node
	}

	RefreshFileExplorerMsg struct{}

	NvimDoneMsg struct {
		Err string
	}

	ChangeFileMessage struct {
		Path string
	}

	RefreshView struct {
		Content string
		Path    string
	}
	HidePopupMessage struct{}
)
