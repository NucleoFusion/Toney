package daily

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/SourcewareLab/Toney/internal/enums"
	"github.com/charmbracelet/bubbles/list"
)

type Tasks struct {
	Recurring []Task `json:"recurring"`
	Unique    []Task `json:"unique"`
}

type Task struct {
	TaskTitle string           `json:"title"`
	TaskDesc  string           `json:"desc"`
	Status    enums.TaskStatus `json:"status"`
}

func (m Task) Title() string       { return m.TaskTitle }
func (m Task) Description() string { return m.TaskDesc }
func (m Task) FilterValue() string { return m.TaskTitle }

func (m Tasks) ItemsAsList() []list.Item {
	fmt.Println(m)

	list := make([]list.Item, 0, len(m.Recurring)+len(m.Unique))

	for _, v := range m.Recurring {
		list = append(list, v)
	}

	for _, v := range m.Unique {
		list = append(list, v)
	}

	return list
}

func GetItems() Tasks {
	path := GetPath()

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		_, err2 := os.Create(path)
		if err2 != nil {
			fmt.Println(err2.Error())
		}
	} else if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	content, _ := os.ReadFile(path)

	tasks := Tasks{}

	err = json.Unmarshal(content, &tasks)
	if err != nil {
		fmt.Println(err.Error())
	}

	return tasks
}

func WriteItems(tasks Tasks) {
	path := GetPath()

	data, _ := json.Marshal(tasks)

	os.WriteFile(path, data, 0o644)
}

func GetPath() string {
	home, _ := os.UserHomeDir()
	date := time.Now().Format("2006-01-02")

	return filepath.Join(home, ".toney", ".daily", date)
}
