// Program to validate PARTICIPANTES.MD file.
// Author: Marcelo Pinheiro - mpinheir@gmail.com  - @mpinheir.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Processfile Loads PARTICIPANTES file and returns slice names containing PARTICIPANTES' names.
func processfile(fileName string) (participantes []string, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	// names slice will hold participant names that were read from file from file PARTICIPANTES.md.
	names := []string{}

	// Reads file line by line and process each line to extract participant's name.
	scanner := bufio.NewScanner(file)

	for linecounter := 1; scanner.Scan(); linecounter++ {
		line := scanner.Text()

		// Ignores first 10 lines.
		if linecounter < 11 {
			continue
		}

		// Splits line contents into hopefully three columns.
		columns := strings.Split(line, "|")

		// Checks if line contains expected number of columns.
		if len(columns) != 5 {
			return []string{}, fmt.Errorf("Line # %d, incorrect number of rows. Want 5, got %d", linecounter, len(columns))
		}

		// Appends participant name to list of names.
		names = append(names, strings.TrimSpace(columns[1]))
	}

	if err := scanner.Err(); err != nil {
		return []string{}, err
	}

	return names, nil
}

func main() {
	// Checks if name of file to be processed was passed as a parameter.
	if len(os.Args) != 2 {
		log.Fatalf("Usage: checkparticipantes <file_name>")
	}

	// Extracts names of participants from PARTICIPANTES.MD file and checks err for potential error.
	names, err := processfile(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	// Checks if there are names to process in slice names.
	if len(names) < 1 {
		log.Fatalf("File: %q does not appear to contain any participant names.\n", os.Args[1])
	}

	// Checks if names are listed in the proper order.
	for i := 0; i < len(names)-1; i++ {
		if names[i] > names[i+1] {
			// Prints error message and exits in case a given name is not in the proper order.
			log.Fatalf("Names not sorted properly: %q > %q\n", names[i], names[i+1])
		}
	}

}
