package main

import (
	"fmt"
	"strconv"

	"github.com/DataDrake/cli-ng/v2/cmd"
	"github.com/EbonJaeger/minecraftcp/internal/widgets"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Run(root *cmd.Root, c *cmd.Sub) {
	app := tview.NewApplication()

	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}

	// Create our server information pane
	infoTextView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetScrollable(true).
		SetWordWrap(true).
		SetChangedFunc(func() {
			app.Draw()
		})
	infoTextView.SetText("Server Info\n\nStatus: [red]Offline[white]")

	// Control area
	slides := []widgets.Slide{
		widgets.Console,
		widgets.Quit,
	}

	pages := tview.NewPages()

	controls := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false).
		SetHighlightedFunc(func(added, removed, remaining []string) {
			pages.SwitchToPage(added[0])
		})

	// Create the pages for all slides.
	previousSlide := func() {
		slide, _ := strconv.Atoi(controls.GetHighlights()[0])
		slide = (slide - 1 + len(slides)) % len(slides)
		controls.Highlight(strconv.Itoa(slide)).
			ScrollToHighlight()
	}
	nextSlide := func() {
		slide, _ := strconv.Atoi(controls.GetHighlights()[0])
		slide = (slide + 1) % len(slides)
		controls.Highlight(strconv.Itoa(slide)).
			ScrollToHighlight()
	}

	for index, slide := range slides {
		title, primitive := slide(app, nextSlide)
		pages.AddPage(strconv.Itoa(index), primitive, true, index == 0)
		fmt.Fprintf(controls, `%d ["%d"][darkcyan]%s[white][""]  `, index+1, index, title)
	}
	controls.Highlight("0")

	grid := tview.NewGrid().
		SetRows(1, 0, 1).
		SetColumns(0, 0, 30).
		SetBorders(true).
		AddItem(newPrimitive("Minecraft Control Panel"), 0, 0, 1, 3, 0, 0, false).
		AddItem(controls, 2, 0, 1, 3, 0, 0, false)

	// Layout for screens wider than 100 cells
	grid.AddItem(pages, 1, 0, 1, 3, 0, 0, false)
	grid.AddItem(infoTextView, 0, 0, 0, 0, 0, 0, false)

	// Layout for screens wider than 100 cells
	grid.AddItem(pages, 1, 0, 1, 2, 0, 100, false)
	grid.AddItem(infoTextView, 1, 2, 1, 1, 0, 100, false)

	// Shortcuts to navigate the slides.
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlN {
			nextSlide()
			return nil
		} else if event.Key() == tcell.KeyCtrlP {
			previousSlide()
			return nil
		}
		return event
	})

	// Run the application
	if err := app.SetRoot(grid, true).EnableMouse(true).SetFocus(controls).Run(); err != nil {
		panic(err)
	}
}
