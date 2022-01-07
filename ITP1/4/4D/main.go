/*
最小値、最大値、合計値
n　個の整数 ai(i=1,2,...n)を入力し、それらの最小値、最大値、合計値を求めるプログラムを作成してください。
Constraints
・0<n≤10000
・−1000000≤ai≤1000000
Sample Input
5
10 1 5 4 17
Sample Output
1 17 37
*/
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	ret, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return ret
}

func main() {
	// スペースで文字列を区切る
	sc.Split(bufio.ScanWords)
	var min float64 = 1 << 30
	var max float64 = -1 << 30
	var sum int = 0

	n := nextInt()
	for i := 0; i < n; i++ {
		m := nextInt()
		min = math.Min(min, float64(m))
		max = math.Max(max, float64(m))
		sum += m
	}
	fmt.Println(int(min), int(max), sum)
}

// よく諦めずにググりきれた！２０２２年１月４？５？日〜７日　まる３日４日かかってやっとAC
//　だいぶがんばった！！


/*
var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return i
}
*/
/*
N, K := scanInt(), scanInt()

A, F := make([]int, N), make([]int, N)
for i := 0; i < N; i++ {
	A[i] = scanInt()
}

for i := 0; i < N; i++ {
	F[i] = scanInt()
}
*/


/*
func main() {
	var n int
	fmt.Scanf("%d",&n)

	var min float64 = 1 << 30
	var max float64 = -1 << 30
	var sum int = 0
	for i := 0; i < n; i++ {
		var a int
		fmt.Scanf("%d",&a)
		min = math.Min(min, float64(a))
		max = math.Max(max, float64(a))
		sum += a
	}

	fmt.Println(int(min), int(max), sum)
}
*/
