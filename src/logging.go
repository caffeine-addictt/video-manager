package src

import "fmt"

func Info(message string) {
	if debug || verbose {
		fmt.Printf("INFO: %s\n", message)
	}
}

func Debug(message string) {
	if debug {
		fmt.Printf("DEBUG: %s\n", message)
	}
}
