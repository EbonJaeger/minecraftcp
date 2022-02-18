package widgets

import "github.com/rivo/tview"

// Slide is a function which returns the slide's main primitive and its title.
// It receives a "nextSlide" function which can be called to advance the
// presentation to the next slide.
type Slide func(app *tview.Application, nextSlide func()) (title string, content tview.Primitive)
