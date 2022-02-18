package widgets

import (
	"github.com/rivo/tview"
)

// Console shows our server console and input box.
func Console(app *tview.Application, nextSlide func()) (title string, content tview.Primitive) {
	// Set up our console text view
	consoleText := tview.NewTextView().
		SetWordWrap(true).
		SetChangedFunc(func() {
			app.Draw()
		})
	consoleText.SetBorder(true).SetTitle("Console")

	// Set the console input field
	consoleInput := tview.NewInputField()
	consoleInput.SetBorder(true)

	// Pack our console elements in a Flex element
	consoleBox := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(consoleText, 0, 1, false).
		AddItem(consoleInput, 3, 1, true)

	return "Console", consoleBox
}
