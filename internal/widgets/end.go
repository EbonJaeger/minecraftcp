package widgets

import (
	"github.com/rivo/tview"
)

// Quit shows the slide with the quit dialog.
func Quit(app *tview.Application, nextSlide func()) (title string, content tview.Primitive) {
	modal := tview.NewModal().
		SetText("Do you want to quit?").
		AddButtons([]string{"Yes", "No"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Yes" {
				app.Stop()
			} else {
				nextSlide()
			}
		})

	return "Quit", modal
}
