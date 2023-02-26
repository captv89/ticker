package main

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
	remaining int
}
 
// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}


// Countdown returns a countdown from the given number
func (a *App) Countdown(number string) {
	num, err := strconv.Atoi(number)
	if err != nil {
		println("Error:", err.Error())
	}
	a.remaining = num
	a.RemainingCountdown()
}

// RemainingCountdown returns the remaining countdown time after each iteration
func (a *App) RemainingCountdown() {
    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()

    for a.remaining > 0 {
        <-ticker.C
        a.remaining--
        remainingTime := strconv.Itoa(a.remaining)
        runtime.EventsEmit(a.ctx, "countdown", remainingTime)
    }

    runtime.EventsEmit(a.ctx, "countdown", "Time's up!")
}
