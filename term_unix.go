// +build !windows

package pretty

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
	"strings"
)

// GetTermSize - Calculates the size of the current terminal window and saves it to Pretty's Terminal struct.
func (p Pretty) GetTermSize() error {
	if runtime.GOOS == "linux" || runtime.GOOS == "unix" {
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
	} else {
		return fmt.Errorf("Sorry, there currently doesn't seem to have support for %v. Please add a bug/feature request on github! (%v)", runtime.GOOS, bugReportURL)
	}

	return nil
}
