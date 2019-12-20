package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println(getVar())
}

//得到非终结符,非终结符像变量用var命名
func getVar() []string {
	f, err := os.Open("test.l")
	if err != nil {
		panic(err)
	}
	br := bufio.NewReader(f)
	vars := make([]string, 0)
	for {
		t, _, err := br.ReadLine()
		if err == io.EOF {
			return vars
		}
		if len(t) <= 0 {
			continue
		}
		ts := strings.Trim(string(t))

		index := strings.Index(ts, "->")
		vars = append(vars, ts[:index])
	}
	return vars
}
