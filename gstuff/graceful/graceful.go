package graceful

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Graceful ..
type Graceful struct {
	sig        os.Signal
	Processing int
}

// New ..
func New(timer time.Duration) (g *Graceful) {
	g = &Graceful{}
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	go func() {
		g.sig = <-gracefulStop
		log.Printf("Graceful shutdown %v\n", g.sig)
		time.Sleep(timer)
		os.Exit(0)
	}()

	return
}

// CheckAlive ..
func (g *Graceful) CheckAlive() bool {
	return g.sig == nil
}

// CheckAndExit ..
func (g *Graceful) CheckAndExit() {
	for {
		time.Sleep(time.Millisecond * 100)
		if !g.CheckAlive() && g.Processing == 0 {
			os.Exit(0)
		}
	}
}
