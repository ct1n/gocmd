// go build -ldflags="-Hwindowsgui" dterm

package main

import (
	"os"
	"os/exec"
	"strings"
)

const dterm = "C:\\Users\\cchirvasuta\\Downloads\\drawterm.exe"

func main() {
	args := make([]string, 0)
	if len(os.Args) > 1 {
		u := ""
		h := os.Args[1]
		if i := strings.IndexRune(h, '@'); i >= 0 {
			u = h[:i]
			if (i < len(h)-1) {
				h = h[i+1:]
			}
		}
		if len(u) > 0 {
			args = append(args, "-u", u)
		}
		if len(h) > 0 {
			args = append(args, "-c", h, "-a", h)
		}
	}
	c := exec.Command(dterm, args...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		println(err.Error())
	}
}
