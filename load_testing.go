package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

func DoRequest(n int, url string, ch chan <- string) {
	start := time.Now()
	resp, _ := http.Get(url)

	req_time := time.Since(start).Seconds()

	body, _ := ioutil.ReadAll(resp.Body)

	ch <- fmt.Sprintf("[%3d] req to url %s took %.2f secs; body len: %d", n, url, req_time, len(body))
}

func main() {

	args_num := len(os.Args)

	if args_num < 2 {
		fmt.Printf("\ntoo few arguments. At least 2 expected, %d given\n", args_num)
		os.Exit(1)
	}

	url := os.Args[1]

	ch := make(chan string)

	reps := 1
	for i := range os.Args {
		if i == 2 {
			reps, _ = strconv.Atoi(os.Args[i])
		}
	}

	fmt.Printf("\ntesting %s %d time(s)\n", url, reps)

	for n := 0; n < reps; n++ {
		go DoRequest(n, url, ch)
	}

	for nn := 0; nn < reps; nn++ {
		fmt.Println(<-ch)
	}
}
