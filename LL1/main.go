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

var cons []string

func main() {
	express := getExpression("eLL.l")

	firsts := make(map[LLesp][]string)
	for _, v := range express {
		esps := removeElement(v, express)

		bs := strings.Split(v.Body, " ")
		// fmt.Println(v, esps)
		cons = make([]string, 0)
		searchfirst(bs[0], esps)
		firsts[v] = removeElementByMap(cons)

	}

	// searchfirst("T", express)

	// cons2 := removeElementByMap(cons)
	for k, v := range firsts {
		fmt.Println(k, v)
	}

}

func removeElement(ele LLesp, slc []LLesp) []LLesp {
	rel := make([]LLesp, 0)
	for _, v := range slc {
		if ele != v {

			rel = append(rel, v)
		}
	}

	return rel
}

//利用map给数组去重和去除$
func removeElementByMap(slc []string) []string {
	temp := make(map[string]int)
	rel := make([]string, 0)
	for k, v := range slc {
		lens := len(temp)
		temp[v] = k
		if len(temp) != lens {
			if v == "$" {
				continue
			}
			rel = append(rel, v)
		}
	}

	return rel
}

func searchfirst(input string, esps []LLesp) {
	for _, v := range esps {
		if v.Head == input {
			bs := strings.Split(v.Body, " ")
			for _, b := range bs {
				if !isProducer(b, esps) {
					cons = append(cons, b)
					// fmt.Println(v, b)
					break
				} else if b == input {
					break
				} else {
					// fmt.Println("searchfirst", v, b)
					searchfirst(b, esps)
				}
			}
		} else {
			cons = append(cons, input)
		}
	}

}

//是非终结符
func isProducer(input string, esps []LLesp) bool {
	for _, v := range esps {
		if v.Head == input {
			return true
		}
	}
	return false
}

//遍历非终结符,得到所有终结符
func searchProducer(esp LLesp) {

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
