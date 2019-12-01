package main

import (
	"fmt"
	"io/ioutil"
	"strconv"

	log "github.com/donnie4w/go-logger/logger"

	"strings"
)

type CToken struct {
	m_iKind     int
	m_szContent string
	m_iRow      int
}

var (
	m_szLexTbl     [50][130]int
	m_szSource     string
	m_szFileName   string
	m_KeywordTbl   map[string]int
	m_szBuffer     string
	m_iRow         int
	m_iPos         int
	m_iNonterminal int
	m_pTokenList   []CToken
)

func init() {
	LexInit()
}

func LexInit() {
	m_pTokenList = make([]CToken, 0)
	m_KeywordTbl = make(map[string]int)
	SetLexTbl(FileToString("lex.txt"))
	SetKeywords(FileToString("KEYWORDS.txt"))
	// for i := 0; i < 50; i++ {
	// 	fmt.Println(m_szLexTbl[i])
	// }

	// fmt.Println(m_KeywordTbl)
}

func main() {

	GetToken(FileToString("test.p"))
	log.Debug(m_pTokenList)
}

func FileToString(filename string) string {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(bs)

}
func SetLexTbl(szStr string) {

	iTmp := 0
	// szs := strings.TrimSpace(szStr)
	szs := strings.Replace(szStr, "\r\n", "", -1)

	for iRow := 0; iRow <= 36; iRow++ {
		for iCol := 0; iCol <= 128; iCol++ {

			t, _ := strconv.Atoi(szs[iTmp:(iTmp + 3)])
			// fmt.Println(iTmp, iTmp+3, szs[iTmp:(iTmp+3)])

			m_szLexTbl[iRow][iCol] = t
			iTmp = iTmp + 3
		}
	}

}

func SetKeywords(szSource string) {
	var szTmp string
	cnt := 0

	for i := 0; i < len(szSource); i++ {

		if szSource[i] != '\n' {

			szTmp += string(szSource[i])
		} else {

			if szTmp != "" {

				m_KeywordTbl[szTmp] = cnt
				szTmp = ""
				cnt++
			}
		}
	}
	if szTmp != "" {

		m_KeywordTbl[szTmp] = cnt
	}

}

func GetToken(zsStr string) bool {
	bTag := true
	m_iPos = 0

	zss := zsStr + ""

	m_iRow = 1
	TmpPos := 0

	for (m_iPos < len(zss)) && (bTag) {

		if string(zss[m_iPos]) == "\n" && TmpPos != m_iPos {

			m_iRow++
			TmpPos = m_iRow
		}
		m_szBuffer += string(zss[m_iPos])

		col := zss[m_iPos]
		if zss[m_iPos] >= 128 {
			col = 128
		}

		bTag = Process(m_szLexTbl[m_iNonterminal][col])
		if !bTag {
			fmt.Println(m_iRow, ":", "词法分析错误，请检查单词")
			return false
		}
		m_iPos++
	}
	return bTag
}

func Process(iTag int) bool {
	iTmp := 0

	if iTag == -99 {

		return false
	}
	if iTag < 0 {
		// log.Debug(m_szBuffer)
		m_szBuffer = m_szBuffer[:len(m_szBuffer)-1]
		m_iPos--
		m_szBuffer = strings.TrimSpace(m_szBuffer)
		// log.Debug(m_szBuffer)
		if iTag == -1 {

			m_szBuffer = strings.ToUpper(m_szBuffer)

			if SearchKeyword(m_szBuffer) {
				EmitToken(iTmp+40, "", m_iRow)
			} else {
				if m_szBuffer == "TRUE" || m_szBuffer == "FALSE" {
					EmitToken(3, m_szBuffer, m_iRow)
				} else {
					EmitToken(1, m_szBuffer, m_iRow)
				}
			}
		}
		if iTag >= -6 && iTag <= -2 {
			EmitToken(-iTag, m_szBuffer, m_iRow)
		}
		if iTag >= -15 && iTag <= -7 {
			EmitToken(-iTag, m_szBuffer, m_iRow)
		}
		if iTag >= -28 && iTag <= -16 {
			EmitToken(-iTag, m_szBuffer, m_iRow)
		}
		if iTag == -42 {
			m_szBuffer := m_szBuffer[:len(m_szBuffer)-1]
			m_iPos--
			EmitToken(3, m_szBuffer, m_iRow)
		}
		m_szBuffer = ""
		m_iNonterminal = 0
	} else {

		m_iNonterminal = iTag

	}
	return true

}

func SearchKeyword(szKey string) bool {
	if _, ok := m_KeywordTbl[szKey]; ok != true {
		return false
	} else {
		return true
	}

}

func EmitToken(iKind int, iContent string, iRow int) {
	ct := CToken{
		m_iKind:     iKind,
		m_szContent: iContent,
		m_iRow:      iRow,
	}
	m_pTokenList = append(m_pTokenList, ct)

}
