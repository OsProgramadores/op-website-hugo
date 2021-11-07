// Program to validate PARTICIPANTES.MD file.
// Marcelo Pinheiro - mpinheir@gmail.com  - @mpinheir.
// Marco Paganini - paganini@paganini.net - @mpaganini.

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"

	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

// checkParticipantes checks the PARTICIPANTES.md file for common errors.
func checkParticipantes(r io.Reader) error {
	const (
		githubPrefix    = "https://github.com/"
		separatorPrefix = "| ----"
	)

	var (
		colindex  []int
		err       error
		inData    bool
		line      string
		mailRegex = regexp.MustCompile(`(?i)[a-z0-9.-_%]+@.+\..*`)
		partcount int
		prevName  string
	)

	// We need a collator to compare order, or accents will break comparison.
	// Loose sets the collator to ignore diacritics, case and weight.
	cl := collate.New(language.English, collate.Loose)

	inreader := bufio.NewReader(r)

	for linecounter := 1; ; linecounter++ {
		line, err = inreader.ReadString('\n')
		// ReadString sets io.EOF if if reaches EOF without finding the delimiter.
		if err == io.EOF {
			break
		}
		if err != nil && err != io.EOF {
			return fmt.Errorf("line %d: Error reading file: %v", linecounter, err)
		}

		// Reject DOS formatted files (\x0d anywhere in the file).
		if strings.ContainsRune(line, '\x0d') {
			return fmt.Errorf("line %d: Found CR character. Please format file as UNIX, not DOS/Windows: %q", linecounter, line)
		}

		// Reject tabs anywhere in the file.
		if strings.ContainsRune(line, '\x09') {
			return fmt.Errorf("line %d: Found tab characters. Only spaces allowed: %q", linecounter, line)
		}

		// Reject lines ending in spaces.
		if strings.HasSuffix(line, " ") {
			return fmt.Errorf("line %d: Extra space at end of line: %q", linecounter, line)
		}

		// Ignore everything until we find a line starting with separatorPrefix.
		// This indicates the end of our headers (start processing next line).
		if strings.HasPrefix(line, separatorPrefix) {
			inData = true
			colindex = indexall(line, '|')
			continue
		}
		if !inData {
			continue
		}

		columns := strings.Split(line, "|")

		// Must have five columns.
		if len(columns) != 5 {
			return fmt.Errorf("line %d: incorrect number of rows. Want 5, got %d", linecounter, len(columns))
		}

		// Make sure columns align. This is UTF-8 aware.
		lineindex := indexall(line, '|')
		if !equalSlice(colindex, lineindex) {
			return fmt.Errorf("line %d: Column markers are not aligned: %q", linecounter, line)
		}

		// Name, email, and github page must start and end with a space (padding
		// between column contents and column separator). We skip the first and
		// last columns, as they're always blank (our table starts and ends
		// with the separator.)
		for _, v := range columns[1 : len(columns)-1] {
			if !strings.HasPrefix(v, " ") {
				return fmt.Errorf("line %d: need a space between %q and the left column separator", linecounter, v)
			}
			if !strings.HasSuffix(v, " ") {
				return fmt.Errorf("line %d: need a space between %q and the right column separator", linecounter, v)
			}
		}

		name := strings.TrimSpace(columns[1])
		email := strings.TrimSpace(columns[2])
		github := strings.TrimSpace(columns[3])

		if prevName == "" {
			prevName = name
		}

		// Name cannot be blank.
		if name == "" {
			return fmt.Errorf("line %d: name cannot be blank: %q", linecounter, line)
		}

		// Name must be alphabetically above or equal to the last one.
		if cl.CompareString(name, prevName) < 0 {
			return fmt.Errorf("line %d: names not sorted properly: %q is in the wrong location; should be %q instead", linecounter, name, prevName)
		}

		// Email must match foo@bar.something or be "No Email"
		if !mailRegex.MatchString(email) && email != "No Email" {
			return fmt.Errorf("line %d: invalid email: %q", linecounter, email)
		}

		// Github page must start with https://github.com/ and have something after the slash.
		if !strings.HasPrefix(github, githubPrefix) || strings.HasSuffix(github, githubPrefix) {
			return fmt.Errorf("line %d: inconsistent github url: %q", linecounter, github)
		}

		partcount++
		prevName = name
	}

	// In a properly formatted file, the last line ends with '\n'. Since
	// ReadString reads the lines up to '\n', io.EOF should always be set with
	// an empty line (nothing after the '\n').  Anything in 'lines' here means
	// there's an entire line without a new line at the end of the file.
	if line != "" {
		return fmt.Errorf("Last line in the file MUST end in a LF (0x10) character (Some editors remove it automatically): %q", line)
	}

	// Make sure we have at least one valid participant.
	if partcount == 0 {
		return fmt.Errorf("no valid participant lines found in input file")
	}

	return nil
}

// indexall returns a slice of ints containing the visible positions of runes
// on the passed string.
func indexall(str string, rch rune) []int {
	var (
		pos int
		ret []int
	)

	// We cannot use pos in the loop directly, as a runes may have width > 1 and
	// we want to count the number of "visible" positions.
	for _, r := range str {
		if r == rch {
			ret = append(ret, pos)
		}
		pos++
	}
	return ret
}

// Return true if both slices are equal, false otherwise.
func equalSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: checkparticipantes <file_name>")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Error opening input file %q: %v\n", os.Args[1], err)
	}
	defer file.Close()

	if err := checkParticipantes(file); err != nil {
		log.Fatalln(err)
	}
}
