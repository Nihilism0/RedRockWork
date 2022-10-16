package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var nice = make(map[string]int, 100)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("number%02d", i)
		value := rand.Intn(100)
		nice[key] = value
	}
	var keys = make([]string, 0, 100)
	for key := range nice {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println(key, nice[key])
	}
}
