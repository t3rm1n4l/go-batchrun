package main

import (
	"fmt"
	"github.com/t3rm1n4l/go-batchrun"
	"os"
	"os/exec"
)

func CommandFn(name, cmd string) error {
	out, _ := exec.Command(cmd).CombinedOutput()
	log, _ := os.Create(name)
	log.Write(out)
	log.Close()
	return nil
}

func main() {

	commands := map[string]string{
		"ls":     "ls",
		"whoami": "whoami",
	}

	runner := batch.New()
	runner.SetConcurrency(3)
	for name, cmd := range commands {
		n := name
		c := cmd
		fn := func() {
			fmt.Println("Starting task : ", n)
			CommandFn(n, c)
			fmt.Println("Completed task : ", n)
		}

		runner.Add(name, fn)
	}

	runner.Start()
	runner.Wait()

}
