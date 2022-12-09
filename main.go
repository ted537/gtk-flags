package main

import (
	"flag"
	"os/exec"
	"regexp"
	"sort"
	"strings"
)

func uniq(strings []string) (ret []string) {
	return []string{}
}

func authors() []string {
	if b, err := exec.Command("git", "log").Output(); err == nil {
		lines := strings.Split(string(b), "\n")

		var a []string
		r := regexp.MustCompile(`^Author:\s*([^ <]+).*$`)
		for _, e := range lines {
			ms := r.FindStringSubmatch(e)
			if ms == nil {
				continue
			}
			a = append(a, ms[1])
		}
		sort.Strings(a)
		var p string
		lines = []string{}
		for _, e := range a {
			if p == e {
				continue
			}
			lines = append(lines, e)
			p = e
		}
		return lines
	}
	return []string{"Yasuhiro Matsumoto <mattn.jp@gmail.com>"}
}

var nameFlag = flag.String("name", "example name", "choose a name!!!")
var reverseFlag = flag.Bool("reverse", false, "should we reverse?")

func main() {
	RunGui()
}
