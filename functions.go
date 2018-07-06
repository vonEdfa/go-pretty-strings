package pretty

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
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

func abort(funcname string, err error) {
	panic(fmt.Sprintf("%s failed: %v", funcname, err))
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

func runCommand(cmd string, args ...string) string {
	exe := exec.Command(cmd, args...)
	exe.Stdin = os.Stdin
	out, err := exe.Output()
	if err != nil {
		var c []string
		c = append(c, cmd)
		c = append(c, args...)
		abort("Command ('"+strings.Join(c, " ")+"')", err)
		//log.Fatal(err)
	}

	return string(out)
}
