// This section of code was created in part or in full using Google Gemini on 28 September 2026.
// Manually checked (and modified) by Jonas Neubürger
// Start AI

package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	gocron "github.com/gritGmbH/go-cron/go-cron"
)

func main() {
	// 1. Setup flags for schedule and port
	schedule := flag.String("s", "* * * * * *", "Cron schedule")
	port := flag.String("p", "18080", "HTTP server port")
	flag.Parse()

	// 2. Parse the command and its arguments
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: go-cron -s \"<schedule>\" -p <port> <command> [args...]")
		os.Exit(1)
	}
	command := args[0]
	cmdArgs := args[1:]

	// 3. Create the cron job using the exported function from go-cron.go
	c, wg := gocron.Create(*schedule, command, cmdArgs)

	// 4. Start the scheduler
	gocron.Start(c)

	// 5. Run the health check HTTP server in a background goroutine
	go gocron.Http_server(*port)

	// 6. Listen for termination signals (like Ctrl+C or Docker stop)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch

	// 7. Stop the scheduler gracefully
	gocron.Stop(c, wg)
}

// End AI