package main

import (
	"context"
	"strconv"
	"time"

	"github.com/willdot/gomacosnotify"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
	remaining int
	stopCountdown chan bool
}
 
// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		stopCountdown: make(chan bool),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Countdown returns a countdown from the given number
func (a *App) Countdown(number int) {
	a.remaining = number
	a.RemainingCountdown()
}

// StopCountdown stops the current countdown
func (a *App) StopCountdown() {
	// runtime.LogInfo(a.ctx, "Stopping countdown")
	a.stopCountdown <- true
}

// RemainingCountdown returns the remaining countdown time after each iteration
func (a *App) RemainingCountdown() {
    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-a.stopCountdown:
			// Exit the loop and stop the countdown
			// runtime.LogInfo(a.ctx, "Countdown stopped")
			a.remaining = 0
			runtime.EventsEmit(a.ctx, "countdown", "0h 0m 0s")
			a.Notify("Timer Stopped!", "Countdown was interrupted")
			return
        case <-ticker.C:
            if a.remaining <= 0 {
                runtime.EventsEmit(a.ctx, "countdown", "Time's up!")
				a.Notify("Time's up!", "Countdown finished")
                return
            }
			// runtime.LogInfo(a.ctx, "Countdown: "+strconv.Itoa(a.remaining))
			// convert remaining time in seconds to hours, minutes and seconds
			remainingTime := strconv.Itoa(a.remaining/3600) + "h " + strconv.Itoa((a.remaining%3600)/60) + "m " + strconv.Itoa((a.remaining%3600)%60) + "s"
            // remainingTime := strconv.Itoa(a.remaining)
            runtime.EventsEmit(a.ctx, "countdown", remainingTime)
            a.remaining--
        }
    }
}

// Notify sends a notification to the user
func (a *App) Notify(title, message string) {
	notifier, err := notify.New()
	if err != nil {
		panic(err)
	}

	notification := notify.Notification{
    Title:        "Time Ticker",
    SubTitle:     title,
    Message:      message,
    ContentImage: "src/assets/images/bell.png",
}
	_, err = notifier.Send(notification)
	if err != nil {
		panic(err)
	}
}
