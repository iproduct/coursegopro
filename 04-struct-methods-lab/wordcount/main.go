package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var stopwords = []string{`ourselves`, `hers`, `between`, `yourself`, `but`, `again`, `there`, `about`, `once`,
	`during`, `out`, `very`, `having`, `with`, `they`, `own`, `an`, `be`, `some`, `for`, `do`, `its`, `yours`,
	`such`, `into`, `of`, `most`, `itself`, `other`, `off`, `is`, `s`, `am`, `or`, `who`, `as`, `from`, `him`,
	`each`, `the`, `themselves`, `until`, `below`, `are`, `we`, `these`, `your`, `his`, `through`, `don`, `nor`,
	`me`, `were`, `her`, `more`, `himself`, `this`, `down`, `should`, `our`, `their`, `while`, `above`, `both`,
	`up`, `to`, `ours`, `had`, `she`, `all`, `no`, `when`, `at`, `any`, `before`, `them`, `same`, `and`, `been`,
	`have`, `in`, `will`, `on`, `does`, `yourselves`, `then`, `that`, `because`, `what`, `over`, `why`, `so`,
	`can`, `did`, `not`, `now`, `under`, `he`, `you`, `herself`, `has`, `just`, `where`, `too`, `only`, `myself`,
	`which`, `those`, `i`, `after`, `few`, `whom`, `t`, `being`, `if`, `theirs`, `my`, `against`, `a`, `by`,
	`doing`, `it`, `how`, `further`, `was`, `here`, `than`}

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	for _, fname := range files {
		//scanner := bufio.NewScanner(os.Stdin)
		file, err := os.Open(fname)
		if err != nil {
			log.Printf("file '%s' not found", fname)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Printf("Entered: %s\n", line)
			//if line == "" {
			//	break
			//}
			words := strings.Fields(line)
			fmt.Printf("%v\n", words)
			for _, w := range words {
				counts[w]++
			}
		}
	}
	for w, n := range counts {
		fmt.Printf("%s -> %d\n", w, n)
	}
}
