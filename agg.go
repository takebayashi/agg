package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var mu = new(sync.Mutex)
var sum = map[string]int{}

func update(key string, value int) {
	mu.Lock()
	if _, exists := sum[key]; !exists {
		sum[key] = 0
	}
	sum[key] += value
	mu.Unlock()
}

func aggregate() {
	mu.Lock()
	for k, v := range sum {
		fmt.Println(k + "\t" + strconv.Itoa(v))
		delete(sum, k)
	}
	mu.Unlock()
}

func main() {
	i := flag.Int("i", 1, "interval for aggregation")
	flag.Parse()
	t := time.NewTicker(time.Duration(*i) * time.Second)
	go func() {
		for {
			select {
			case <-t.C:
				aggregate()
			}
		}
	}()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "\t")
		k := s[0]
		v, _ := strconv.Atoi(s[1])
		update(k, v)
	}
	aggregate()
}
