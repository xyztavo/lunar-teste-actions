package main

import (
	"context"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func runCommand(ctx context.Context, wg *sync.WaitGroup, name, cmd string, args ...string) {
	defer wg.Done()
	c := exec.Command(cmd, args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	log.Printf("Starting %s", name)
	if err := c.Start(); err != nil {
		log.Printf("%s failed to start: %v", name, err)
		return
	}
	done := make(chan error, 1)
	go func() {
		done <- c.Wait()
	}()
	select {
	case err := <-done:
		log.Printf("%s exited with error: %v", name, err)
	case <-ctx.Done():
		log.Printf("Shutting down %s gracefully", name)
		if err := c.Process.Signal(syscall.SIGTERM); err != nil {
			log.Printf("Failed to send SIGTERM to %s: %v", name, err)
			c.Process.Kill()
		} else {
			select {
			case <-done:
				log.Printf("%s shut down gracefully", name)
			case <-time.After(5 * time.Second):
				log.Printf("Force killing %s", name)
				c.Process.Kill()
				<-done
			}
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		log.Println("Received signal, shutting down...")
		cancel()
	}()

	var wg sync.WaitGroup

	// Templ watcher
	wg.Add(1)
	go runCommand(ctx, &wg, "templ", "templ", "generate", "--watch")

	// Tailwind watcher
	wg.Add(1)
	go runCommand(ctx, &wg, "tailwind", "tailwindcss", "-i", "public/input.css", "-o", "public/output.css", "--watch")

	// Air
	wg.Add(1)
	go runCommand(ctx, &wg, "air", "air")

	wg.Wait()
	log.Println("All processes finished")
}
