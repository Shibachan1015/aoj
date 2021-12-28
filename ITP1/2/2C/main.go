// 参考元 jakenu0x5e https://onlinejudge.u-aizu.ac.jp/solutions/problem/ITP1_2_C  
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Scanner struct {
	reader *bufio.Scanner
}

func (s *Scanner) Next() string {  //メソッド　https://pkg.go.dev/bufio　参照
	s.reader.Scan()
	return s.reader.Text()
}

func (s *Scanner) Nextint() int {
	s.reader.Scan()
	val, err := strconv.Atoi(s.reader.Text())  //パッケージstrconvは、基本的なデータ型の文字列表現との間の変換を実装する
	if err != nil {                            //最も一般的な数値変換は、Atoi（文字列からint）とItoa（intから文字列）である。
		panic(err)                            // パッケージpanicはエラーハンドルを実装する
	}
	return val
}

var sc *Scanner

func init() {
	sc = &Scanner{bufio.NewScanner(os.Stdin)}
	sc.reader.Split(bufio.ScanWords)
}

func main() {
	a := sc.Nextint()
	b := sc.Nextint()
	c := sc.Nextint()

	vals := []int{a, b, c}
	sort.Ints(vals)
	fmt.Printf("%d %d %d\n", vals[0], vals[1], vals[2])
}
