package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type sets struct {
	Fists  []string
	Fllows []string
}

var LLTable map[string]*sets

func init() {
	LLTable = make(map[string]*sets)
}

func main() {
	fmt.Println(getVar())
}

//从产生式获取firsts
func getFirsts(expres []string, vars []string) []string {

}

//从产生式获取fellows
func getFellows(expres []string, vars []string) []string {

}

//得到非终结符和产生式,非终结符像变量用var命名
func getVar() ([]string, []string) {
	f, err := os.Open("test.l")
	if err != nil {
		panic(err)
	}
	br := bufio.NewReader(f)
	vars := make([]string, 0)
	expresion := make([]string, 0)
	for {
		t, _, err := br.ReadLine()
		if err == io.EOF {
			return expresion, vars
		}
		if len(t) <= 0 {
			continue
		}
		ts := strings.Replace(string(t), " ", "", -1)
		expresion = append(expresion, ts)
		index := strings.Index(ts, "->")
		vars = append(vars, ts[:index])
	}
	return expresion, vars
}
