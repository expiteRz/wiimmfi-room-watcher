package terminal

import (
	"app.rz-public.xyz/wiimmfi-room-watcher/utils"
	"app.rz-public.xyz/wiimmfi-room-watcher/web"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"log"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

var (
	// Format used for non-private room
	nonPrivateMemberFormat = []table.Column{
		{"#", 2},
		{"Friend Code", 14},
		{"Name", 20}, // allocates 20 bytes space for any 2-byte string
		{"VR", 5},    // Possibly any hackers gains over 9999 vr, so declare the width as 5
		{"BR", 5},    // Same as above
	}
	// Format used for private room and competitive room similar to Lounge
	privateMemberFormat = []table.Column{
		{"#", 2},
		{"Friend Code", 14},
		{"Team", 10},
		{"Name", 20},  // allocates 20 bytes space for any 2-byte string
		{"Points", 3}, // Obviously the lounge player can get 3 digit points. must be more allocation?
	}
)

type model struct {
	table       table.Model
	keymap      keymap
	cmdInput    textinput.Model
	toggleInput bool

	help help.Model
}

type keymap struct {
	enterCommand key.Binding
}

func initializeModel() model {
	k := keymap{
		enterCommand: key.NewBinding(
			key.WithKeys("i"),
			key.WithHelp("i", "Input command")),
	}

	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 128
	ti.Width = 20

	values := []table.Row{
		{"1", "0000-0000-0000", "no name", "5000", "5000"},
		{"2", "0000-0000-0000", "no name", "5000", "5000"},
		{"3", "0000-0000-0000", "no name", "5000", "5000"},
	}
	t := table.New(
		table.WithColumns(nonPrivateMemberFormat),
		table.WithRows(values),
		table.WithFocused(true),
		table.WithHeight(12),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("240")).BorderBottom(true).Bold(false)
	s.Selected = s.Selected.Foreground(lipgloss.Color("229")).Background(lipgloss.Color("57")).Bold(false)
	t.SetStyles(s)

	return model{t, k, ti, false, help.New()}
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
			m.toggleInput = false
		//case "q", "ctrl+c":
		//	return m, tea.Quit
		case "i":
			m.toggleInput = true
			m.cmdInput.Focus()
		case "enter":
			if m.toggleInput {
				res := tea.Printf("%s", m.cmdInput.Value())
				m.cmdInput.SetValue("")
				return m, tea.Batch(res)
			}
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.table.SelectedRow()[1]),
			)
		}
	}
	if m.toggleInput {
		m.cmdInput, cmd = m.cmdInput.Update(msg)
	} else {
		m.table, cmd = m.table.Update(msg)
	}
	return m, cmd
}

func (m model) View() string {
	switch web.RoomData.Setting.GameMode {
	case utils.ModePrivateBalloonBattle, utils.ModePrivateVS, utils.ModePrivateCoinBattle:
		m.table.SetColumns(privateMemberFormat)
	default:
		m.table.SetColumns(nonPrivateMemberFormat)
	}

	return baseStyle.Render(m.table.View()) + "\n" + m.viewHelp()
}

func (m model) viewHelp() string {
	return "\n" + m.help.ShortHelpView([]key.Binding{m.keymap.enterCommand})
}

func TerminalTest() {
	if _, err := tea.NewProgram(initializeModel()).Run(); err != nil {
		log.SetPrefix("[Terminal:Test] ")
		log.Fatalln(err)
	}
}
