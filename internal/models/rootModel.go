package models

import (
	"fmt"

	"github.com/SourcewareLab/Toney/internal/enums"
	"github.com/SourcewareLab/Toney/internal/messages"
	"github.com/SourcewareLab/Toney/internal/models/daily"
	"github.com/SourcewareLab/Toney/internal/models/diary"
	filepopup "github.com/SourcewareLab/Toney/internal/models/filePopup"
	homemodel "github.com/SourcewareLab/Toney/internal/models/homeModel"
	"github.com/SourcewareLab/Toney/internal/models/menu"

	"github.com/SourcewareLab/Toney/internal/colors"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type RootModel struct {
	Width         int
	Height        int
	Page          enums.Page
	Home          *homemodel.HomeModel
	Menu          *menu.Menu
	Daily         *daily.Daily
	Diary         *diary.Diary
	CurrentPage   enums.Page
	ShowPopup     bool
	FilePopupType enums.PopupType
	FilePopup     *filepopup.FilePopup
	isLoading     bool
}

func NewRoot() *RootModel {
	return &RootModel{
		Page:      enums.MenuPage,
		ShowPopup: false,
		isLoading: true,
	}
}

func (m RootModel) Init() tea.Cmd {
	return nil
}

func (m *RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case messages.TaskPopupMessage:
		if m.CurrentPage == enums.DailyPage {
			dailypage, cmd := m.Daily.Update(msg)
			if d, ok := dailypage.(*daily.Daily); ok { // Type matching, cause I cant assign it straightaway
				m.Daily = d
				return m, cmd
			}
			return m, cmd
		}
	case messages.ChangePage:
		switch msg.Page {
		case enums.MenuPage:
			m.CurrentPage = enums.MenuPage
		case enums.HomePage:
			m.CurrentPage = enums.HomePage
		case enums.DailyPage:
			m.CurrentPage = enums.DailyPage
		case enums.DiaryPage:
			m.CurrentPage = enums.DiaryPage
		case enums.Quit:
			return m, tea.Quit
		}
	case messages.ShowLoader:
		m.isLoading = true
		return m, nil
	case messages.HideLoader:
		m.isLoading = false
		return m, nil
	case messages.ShowPopupMessage:
		m.FilePopup = filepopup.NewPopup(msg.Type, msg.Curr)
		m.ShowPopup = true
	case messages.HidePopupMessage:
		m.ShowPopup = false
	case messages.RefreshFileExplorerMsg:
		m.Home.FileExplorer.Update(msg)
		return m, nil
	case messages.EditorClose:
		if m.CurrentPage == enums.DiaryPage {
			m.Diary.Update(msg)
			return m, nil
		}
		if msg.Err != nil {
			fmt.Println(msg.Err.Error())
		}
		m.Home.FileExplorer.Update(msg)
		m.Home.Viewer.Update(msg)

		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height

		if m.Home != nil { // Checking whether this is an app resize or app open
			m.Menu.Update(msg)
			m.Home.Update(msg)
			m.Diary.Update(msg)
			m.Daily.Update(msg)
		} else {
			m.Home = homemodel.NewHome(msg.Width, msg.Height)
			m.Menu = menu.NewMenu(msg.Width, msg.Height)
			m.Daily = daily.NewDaily(msg.Width, msg.Height)
			m.Diary = diary.NewDiary(msg.Width, msg.Height)
		}

		m.isLoading = false

		return m, nil
	}

	var cmd tea.Cmd

	if m.ShowPopup {
		_, cmd = m.FilePopup.Update(msg)
	} else {
		switch m.CurrentPage {
		case enums.MenuPage:
			_, cmd = m.Menu.Update(msg)
		case enums.HomePage:
			_, cmd = m.Home.Update(msg)
		case enums.DailyPage:
			_, cmd = m.Daily.Update(msg)
		case enums.DiaryPage:
			_, cmd = m.Diary.Update(msg)
		default:
			fmt.Printf("UNHANDLED MSG: %#v\n", msg)
		}
	}

	return m, cmd
}

func (m *RootModel) View() string {
	if m.isLoading {
		return lipgloss.NewStyle().Render("Loading...")
	}

	if m.ShowPopup && m.FilePopup != nil {
		return lipgloss.Place(m.Width+2, m.Height+2, lipgloss.Center, lipgloss.Center, m.FilePopup.View())
	}

	switch m.CurrentPage {
	case enums.HomePage:
		return lipgloss.NewStyle().Background(colors.ColorPalette().Background).Render(m.Home.View())
	case enums.MenuPage:
		return m.Menu.View()
	case enums.DailyPage:
		return m.Daily.View()
	case enums.DiaryPage:
		return m.Diary.View()
	default:
		return lipgloss.NewStyle().Background(colors.ColorPalette().Background).Render(m.Home.View())
	}
}
