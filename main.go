package main

// This section of code was created in part or in full using Google Gemini on 28 September 2026.
// Manually checked (and modified) by Jonas Neubürger
// Start AI

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	gocron "go-cron/go-cron"
)

func main() {
	schedule := flag.String("s", "* * * * * *", "Cron schedule")
	port := flag.String("p", "18080", "HTTP server port")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: go-cron -s \"<schedule>\" -p <port> <command> [args...]")
		os.Exit(1)
	}
	command := args[0]
	cmdArgs := args[1:]

	c, wg := gocron.Create(*schedule, command, cmdArgs)
	gocron.Start(c)

	go gocron.Http_server(*port)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch

	gocron.Stop(c, wg)
}

// End AI