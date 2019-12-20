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

type LLesp struct {
	Head string
	Body string
}

var LLTable map[string]*sets

func init() {
	LLTable = make(map[string]*sets)

	// extendLLWriteToFile("eLL.l", decodeLL("LL.l"))
}

func main() {
	express := getExpression("eLL.l")
	for _, v := range express {
		fmt.Printf("%v\n", v)
	}
}

//第一次查找
func searchfirst(input string, esps []LLesp) interface{} {
	for _, v := range esps {
		if v.Head == input {
			return v
		}
	}
	return input
}

//遍历非终结符,得到所有终结符
func searchProducer(esp LLesp) {
	esps := getExpression("eLL.l")
	firsts := make(map[LLesp][]string)
	for _, v := range esps {
		bs := strings.Split(v.Body, " ")
		re := searchfirst(bs[0], esps)
		for {
			switch re.(type) {
			case string:
				firsts[v] = append(firsts[v], re.(string))
				break
			case LLesp:
				re = searchfirst(re.(LLesp).Body[], esps)
			}
		}
	}

}

//从产生式获取firsts
func getFirsts(expres []string, vars []string) []string {
	firsts := make([]string, 0)
	return firsts
}

//从产生式获取fellows
func getFellows(expres []string, vars []string) []string {
	fellows := make([]string, 0)
	return fellows
}

//展开文法表达式
func decodeLL(llname string) []string {
	f, err := os.Open(llname)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	br := bufio.NewReader(f)
	llstr := make([]string, 0)
	for {
		t, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		if len(t) <= 0 {
			continue
		}

		ts := strings.Replace(string(t), " ", "", -1)
		llstr = append(llstr, ts)
	}

	outstr := make([]string, 0)
	for _, v := range llstr {
		index := strings.Index(v, "->")
		head := v[:index+2]
		body := v[index+2:]
		ts := strings.Split(body, "|")
		ss := make([]string, 0)
		for _, m := range ts {
			ss = append(ss, head+m)
		}

		outstr = append(outstr, ss...)
	}
	return outstr
}

//展开的文法写入文件
func extendLLWriteToFile(oname string, eLL []string) {
	os.Remove("eLL.l")
	f, err := os.OpenFile(oname, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for _, v := range eLL {
		n, _ := f.Seek(0, os.SEEK_END)
		_, err = f.WriteAt([]byte(v+"\n"), n)

	}

}

//从文件得到非终结符和产生式
func getExpression(ellname string) []LLesp {
	f, err := os.Open(ellname)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	br := bufio.NewReader(f)
	expression := make([]LLesp, 0)
	for {
		t, _, err := br.ReadLine()
		if err == io.EOF {
			return expression
		}
		if len(t) <= 0 {
			continue
		}
		ts := string(t)
		index := strings.Index(ts, "->")

		expression = append(expression, LLesp{ts[:index], ts[index+2:]})

	}
	return expression
}
