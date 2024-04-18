package main

import (
	"fmt"
	"io"
)

func main() {

}

func run(project string, out io.Writer) error {
	if project == "" {
		return fmt.Errorf("project directory is required")
	}

	return nil
}
