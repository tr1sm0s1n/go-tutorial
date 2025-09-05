package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

const (
	width         = 120
	glamourGutter = 2
)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render

type slide struct {
	current  int
	contents []string
	viewport viewport.Model
	renderer *glamour.TermRenderer
}

func startSlides(cs []string) (*slide, error) {
	vp := viewport.New(width, 30)
	vp.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("60")).
		PaddingRight(2)

	glamourRenderWidth := width - vp.Style.GetHorizontalFrameSize() - glamourGutter

	rr, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(glamourRenderWidth),
	)
	if err != nil {
		return nil, err
	}

	str, err := rr.Render(cs[0])
	if err != nil {
		return nil, err
	}

	vp.SetContent(str)

	return &slide{
		current:  0,
		contents: cs,
		viewport: vp,
		renderer: rr,
	}, nil
}

func (e slide) Init() tea.Cmd {
	return nil
}

func (e slide) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return e, tea.Quit
		case "left", "right":
			var cmd tea.Cmd

			s := msg.String()
			if s == "left" {
				e.current--
			} else {
				e.current++
			}

			if e.current > len(e.contents)-1 {
				e.current = 0
			} else if e.current < 0 {
				e.current = len(e.contents) - 1
			}

			for i, c := range e.contents {
				if i == e.current {
					str, _ := e.renderer.Render(c)
					e.viewport.SetContent(str)
				}
			}

			return e, cmd
		default:
			var cmd tea.Cmd
			e.viewport, cmd = e.viewport.Update(msg)
			return e, cmd
		}
	default:
		return e, nil
	}
}

func (e slide) View() string {
	return e.viewport.View() + e.helpView()
}

func (e slide) helpView() string {
	return helpStyle("\n  ↑/↓: Navigate • ←/→: Turn • q: Quit\n")
}

func main() {
	model, err := startSlides(data)
	if err != nil {
		fmt.Println("Could not initialize Bubble Tea model:", err)
		os.Exit(1)
	}

	if _, err := tea.NewProgram(model, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Bummer, there's been an error:", err)
		os.Exit(1)
	}
}
