package main

import (
	"flag"
	"fmt"
	"github.com/t3rm1n4l/go-batchrun"
	"os"
	"os/exec"
	"path"
	"time"
)

func CommandFn(cmd, logfile string) error {
	out, err := exec.Command(cmd).CombinedOutput()
	log, _ := os.Create(logfile)
	log.Write(out)
	log.Close()

	return err
}

func main() {
	var (
		logdir  = flag.String("logdir", ".", "Log directory")
		concurr = flag.Int("concurrency", 1, "Concurrency")
		help    = flag.Bool("help", false, "Help")
	)

	var Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage %s: prog1 prog2 prog3...\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if flag.NArg() < 1 || *help {
		Usage()
		os.Exit(1)
	}

	runner := batch.New()
	runner.SetConcurrency(*concurr)
	for i := 0; i < flag.NArg(); i++ {
		cmd := flag.Arg(i)
		name := path.Base(cmd)
		name2 := fmt.Sprintf("%s.%d", name, i)
		logfile := path.Join(*logdir, fmt.Sprintf("%s.log", name2))
		fn := func() {
			c := cmd
			n := name2
			l := logfile
			fmt.Println("Starting task : ", n)
			t1 := time.Now()
			CommandFn(c, l)
			diff := time.Now().Sub(t1)
			fmt.Printf("Completed task : %s in %s\n", n, diff)
		}

		runner.Add(name, fn)
	}

	runner.Start()
	runner.Wait()

}
