package pretty

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// CombineStrings - Combines an array of strings, treating each string as a separate paragraph.
//
// p:    Main Pretty struct containing all your settings.
// strs: An array of strings to be combined.
func (p Pretty) CombineStrings(strs ...string) string {
	var ret string

	for i, s := range strs {
		ret += s
		if i < len(strs)-1 && !strings.HasSuffix(s, p.Newline) {
			ret += p.Paragraph
		} else if strings.HasSuffix(s, p.Newline) && !strings.HasSuffix(s, p.Paragraph) {
			ret += p.Newline
		}
	}

	if p.UseTrailingNewLine && !strings.HasSuffix(ret, p.Newline) {
		ret += p.Newline
	}

	return ret
}

// GetTermSize - Calculates the size of the current terminal window and saves it to Pretty's Terminal struct.
func (p Pretty) GetTermSize() error {
	if strings.ToLower(runtime.GOOS) == "linux" {
		size := runCommand("stty", "size")
		size = strings.Trim(size, "\n")
		sizes := strings.Split(string(size), " ")
		h, err := strconv.ParseInt(sizes[0], 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		w, err := strconv.ParseInt(sizes[1], 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		p.Terminal.Height = int(h)
		p.Terminal.Width = int(w)
	} else if strings.ToLower(runtime.GOOS) != "windows" {
		sizes := runCommand("mode", "con")
		for _, s := range strings.Split(sizes, "\r\n") {
			if strings.Contains(s, "Columns") {
				size := strings.TrimFunc(s, func(r rune) bool {
					return !unicode.IsNumber(r)
				})
				w, err := strconv.ParseInt(size, 10, 32)
				if err != nil {
					log.Fatal(err)
				}
				p.Terminal.Width = int(w)
			} else if strings.Contains(s, "Lines") {
				size := strings.TrimFunc(s, func(r rune) bool {
					return !unicode.IsNumber(r)
				})
				h, err := strconv.ParseInt(size, 10, 32)
				if err != nil {
					log.Fatal(err)
				}
				p.Terminal.Height = int(h)
			}
		}
	} else {
		return fmt.Errorf("Sorry, there currently doesn't seem to be support for %v. Please add a bug/feature request on github! (%v)", runtime.GOOS, bugReportURL)
	}

	return nil
}

func runCommand(cmd string, args ...string) string {
	exe := exec.Command(cmd, args...)
	exe.Stdin = os.Stdin
	out, err := exe.Output()
	if err != nil {
		log.Fatal(err)
	}

	return string(out)
}

func getLongestStringLength(list []string) int {
	var longest int

	for _, s := range list {
		if utf8.RuneCountInString(s) > longest {
			longest = len(s)
		}
	}

	return longest
}

func repeatChar(char string, amount int) string {
	var ret string

	for i := 0; i < amount; i++ {
		ret += char
	}

	return ret
}
