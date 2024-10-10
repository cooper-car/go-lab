package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// 用併發改寫
func main() {
	//fmt.Println(fileLine("./bufferchannel-1.go"))
	dir := "."
	total := 0
	var wg sync.WaitGroup

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".go" || !info.IsDir() {

			wg.Add(1)
			go func() {
				cnt := fileLine2(path)
				total += cnt
				wg.Done()
			}()
		}
		return nil
	})

	wg.Wait()

	fmt.Println("total lines: ", total)
}

func fileLine2(path string) int {

	cnt := 0
	file, err := os.Open(path)
	if err != nil {
		log.Print(err)
		return cnt
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		ctx, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		txt := strings.TrimSpace(string(ctx))
		if txt == "" || strings.HasPrefix(txt, "//") {
			continue
		}

		cnt++
	}

	return cnt
}
