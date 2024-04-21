package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

type executer interface {
	execute() (string, error)
}

func main() {
	proj := flag.String("p", "", "Project directory.")
	flag.Parse()

	if err := run(*proj, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(project string, out io.Writer) error {
	if project == "" {
		return fmt.Errorf("project directory is required: %w", ErrValidation)
	}

	pipeline := make([]executer, 4)

	pipeline[0] = newStep(
		"go build",
		"go",
		"Go Build: SUCCESS",
		project,
		[]string{"build", ".", "errors"})

	pipeline[1] = newStep(
		"go test",
		"go",
		"Go Test: SUCCESS",
		project,
		[]string{"test", "-v"})

	pipeline[2] = newExceptionStep(
		"go fmt",
		"gofmt",
		"Gofmt: SUCCESS",
		project,
		[]string{"-l", "."})

	pipeline[3] = newTimeoutStep(
		"git push",
		"git",
		"Git Push: SUCCESS",
		project,
		[]string{"push", "origin", "master"},
		10*time.Second)

	for _, s := range pipeline {
		msg, err := s.execute()
		if err != nil {
			return err
		}

		_, err = fmt.Fprintln(out, msg)
		if err != nil {
			return err
		}
	}

	return nil
}
