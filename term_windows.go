// +build windows

package pretty

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"unicode"
)

func init() {
	if major, _ := getWindowsVersion(); major >= 10 {
		EnableWindowsANSI()
	}
}

// EnableWindowsANSI - Enables ANSI escape codes and formatting in Windows terminals. Works in Win10 and later.
func EnableWindowsANSI() error {
	var err error

	// Eable ANSI Formatting on Windows
	// https://github.com/sirupsen/logrus/issues/496
	if runtime.GOOS == "windows" {
		handle := syscall.Handle(os.Stdout.Fd())
		kernel32DLL := syscall.NewLazyDLL("kernel32.dll")
		setConsoleModeProc := kernel32DLL.NewProc("SetConsoleMode")
		_, _, err = setConsoleModeProc.Call(uintptr(handle), 0x0001|0x0002|0x0004)
	}

	return err
}

// GetTermSize - Calculates the size of the current terminal window and saves it to Pretty's Terminal struct.
func (p Pretty) GetTermSize() error {
	if runtime.GOOS == "windows" {
		sizes := runCmdCommand("mode", "con")
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
		return fmt.Errorf("Sorry, there currently doesn't seem to have support for %v. Please add a bug/feature request on github! (%v)", runtime.GOOS, bugReportURL)
	}

	return nil
}

func getWindowsVersion() (int, int) {
	long := runCmdCommand("ver")
	short := regexp.MustCompile("[0-9]+").FindAllString(long, 2)
	major, _ := strconv.Atoi(short[0])
	minor, _ := strconv.Atoi(short[1])

	return major, minor
}

func runCmdCommand(cmd string, args ...string) string {
	a := strings.Join(args, " ")
	out := runCommand("cmd", "/C", cmd, a)

	return out
}
