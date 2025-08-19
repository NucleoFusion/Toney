package daily

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/SourcewareLab/Toney/internal/config"
	"github.com/SourcewareLab/Toney/internal/enums"
	"github.com/charmbracelet/bubbles/list"
)

type Tasks struct {
	Recurring []Task       `json:"recurring"`
	Unique    []Task       `json:"unique"`
	Github    []GithubTask `json:"github"`
	Todo      []TodoTask   `json:"todo"`
}

type GithubTask struct {
	TaskTitle string           `json:"title"`
	TaskDesc  string           `json:"desc"`
	Status    enums.TaskStatus `json:"status"`
	Ref       string           `json:"ref"`
	Repo      string           `json:"repo"`
	Owner     string           `json:"owner"`
	// We will not store these data in the file
	Link     string   `json:"-"`
	Labels   []string `json:"-"`
	Assignee []string `json:"-"`
}

type TodoTask struct {
	TaskType    enums.TaskType   `json:"-"`
	Status      enums.TaskStatus `json:"status"`
	ProjectName string           `json:"projectName"`
	Path        string           `json:"-"`
	RelPath     string           `json:"relPath"`
	Text        string           `json:"text"`
	Line        int              `json:"line"`
}

type Task struct {
	TaskTitle string           `json:"title"`
	TaskDesc  string           `json:"desc"`
	Status    enums.TaskStatus `json:"status"`
	// We will not store these data in the file
	Index    int            `json:"-"` // Point to index in the respective type array
	TaskType enums.TaskType `json:"-"`
}

func (m TodoTask) Title() string {
	return fmt.Sprintf("%s ~ %s:%d", m.ProjectName, m.RelPath, m.Line)
}
func (m TodoTask) Description() string { return strings.SplitAfter(m.Text, "TODO:")[1] }
func (m TodoTask) FilterValue() string { return m.Title() }

func (m GithubTask) Title() string       { return m.TaskTitle }
func (m GithubTask) Description() string { return m.TaskDesc }
func (m GithubTask) FilterValue() string { return m.TaskTitle }

func (m Task) Title() string       { return m.TaskTitle }
func (m Task) Description() string { return m.TaskDesc }
func (m Task) FilterValue() string { return m.TaskTitle }

func (m Tasks) ItemsAsList() []list.Item {
	lst1 := TaskToItems(m.Recurring, enums.RecurringTask)
	lst2 := TaskToItems(m.Unique, enums.UniqueTask)
	lst3 := TodoTaskToItems(m.Todo)

	return append(lst1, append(lst2, lst3...)...)
}

func TaskToItems(tasks []Task, typ enums.TaskType) []list.Item {
	list := make([]list.Item, 0)
	for i, v := range tasks {
		v.Index = i
		v.TaskType = typ
		list = append(list, v)
	}
	return list
}

func TodoTaskToItems(tasks []TodoTask) []list.Item {
	list := make([]list.Item, 0)
	for _, v := range tasks {
		v.TaskType = enums.TodoTask
		list = append(list, v)
	}
	return list
}

func GetItems() (Tasks, error) {
	tasks := Tasks{}

	content, err := ReadTaskFile()
	if err != nil {
		return tasks, err
	}

	err = json.Unmarshal(content, &tasks)
	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

func ReadTaskFile() ([]byte, error) {
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

		tasks := Tasks{}
		json.Unmarshal(content, &tasks)

		tasks.Unique = make([]Task, 0)

		tasks.Todo, err2 = ReadTodos()
		if err2 != nil {
			return []byte{}, err
		}

		data, err2 := json.Marshal(tasks)
		if err2 != nil {
			return []byte{}, err
		}

		f.Write(data)
	} else if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	content, _ := os.ReadFile(path)

	return content, nil
}

func ReadTodos() ([]TodoTask, error) {
	projs := config.AppConfig.General.Todo.Projects
	result := make([]TodoTask, 0)

	for _, v := range projs {
		var tasks []TodoTask // TODO: Exclude Dir's

		cmd := exec.Command("bash", "-c", fmt.Sprintf(`
rg -n -i -P --json '(?:\/\/|#|--|/\*+)\s*TODO:?\s*(.*)' %s \
| jq -r 'select(.type=="match") | {relPath: .data.path.text, line: .data.line_number, text: .data.submatches[0].match.text}' \
| jq -s .`, v.Path))

		out, err := cmd.Output()
		if err != nil {
			return []TodoTask{}, err
		}

		err = json.Unmarshal(out, &tasks)
		if err != nil {
			return []TodoTask{}, err
		}

		for k, val := range tasks {
			val.ProjectName = v.Name
			val.Status = enums.Pending
			val.Path = v.Path
			val.RelPath = strings.Replace(val.RelPath, v.Path, ".", 1)

			tasks[k] = val
		}

		result = append(result, tasks...)
	}

	return result, nil
}

func WriteItems(tasks Tasks) {
	path := GetPath()

	data, _ := json.Marshal(tasks)

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
