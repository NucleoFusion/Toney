package daily

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/SourcewareLab/Toney/internal/config"
	"github.com/SourcewareLab/Toney/internal/enums"
	"github.com/charmbracelet/bubbles/list"
	"github.com/jszwec/csvutil"
)

type Task struct {
	TaskTitle string           `csv:"title"`
	TaskDesc  string           `csv:"desc"`
	Status    enums.TaskStatus `csv:"status"`
	// We will not store these data in the file
	ID       int            `csv:"-"` // Point to index in the respective type array
	TaskType enums.TaskType `csv:"-"`
}

func (m Task) Title() string       { return m.TaskTitle }
func (m Task) Description() string { return m.TaskDesc }
func (m Task) FilterValue() string { return m.TaskTitle }

func TaskToItems(tasks []Task) []list.Item {
	list := make([]list.Item, 0)
	for i, v := range tasks {
		v.ID = i + 1
		v.TaskType = enums.RecurringTask
		list = append(list, v)
	}
	return list
}

func GetItems() []Task {
	path := GetPath()

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		f, err2 := os.Create(path)
		if err2 != nil {
			fmt.Println(err2.Error())
		}

		content, err2 := os.ReadFile(GetYesterdayPath())
		if err2 != nil {
			fmt.Println(err2.Error())
		}

		tasks := make([]Task, 0)
		csvutil.Unmarshal(content, &tasks)

		for k := range tasks {
			tasks[k].ID = k + 1
			tasks[k].Status = enums.Pending
		}

		data, err2 := csvutil.Marshal(tasks)
		if err2 != nil {
			fmt.Println(err2.Error())
		}

		f.Write(data)
	} else if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	content, _ := os.ReadFile(path)

	tasks := make([]Task, 0)
	err = csvutil.Unmarshal(content, &tasks)
	if err != nil {
		fmt.Println(err.Error())
	}

	return tasks
}

func WriteItems(tasks []Task) {
	path := GetPath()

	data, _ := csvutil.Marshal(tasks)

	os.WriteFile(path, data, 0o644)
}

func GetPath() string {
	home, _ := os.UserHomeDir()
	date := time.Now().Format("2006-01-02")

	return filepath.Join(home, config.AppConfig.General.NotesDir, ".daily", date)
}

func GetYesterdayPath() string {
	home, _ := os.UserHomeDir()
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	return filepath.Join(home, config.AppConfig.General.NotesDir, ".daily", yesterday)
}
