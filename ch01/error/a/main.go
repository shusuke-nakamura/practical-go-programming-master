package main

import (
	"bufio"
	"os"
)

func main() {
	// error-intro
	f, err := os.Open("important.text")
	if err != nil {
		// エラーハンドリング
	}
	r := bufio.NewReader(f)
	l, err := r.ReadString('\n')
	if err != nil {
		// エラーハンドリング
	}
	// ここではエラーが発生しない
	// error-intro
	_ = l
}
