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
	dir := "."
	total := 0
	var wg sync.WaitGroup
	var wgTotal sync.WaitGroup
	channel := make(chan int)

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".go" || !info.IsDir() {

			wg.Add(1)
			go func() {
				cnt := fileLine3(path)
				channel <- cnt

				wg.Done()
			}()
		}
		return nil
	})

	wgTotal.Add(1)
	go func() {
		for cnt := range channel {
			total += cnt
		}
		wgTotal.Done()
	}()

	wg.Wait()
	close(channel)
	wgTotal.Wait()

	fmt.Println("total lines: ", total)
}

func fileLine3(path string) int {

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
