package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

func DoRequest(n int, url string, ch chan<- string, expected_len int) {
	start := time.Now()
	resp, _ := http.Get(url)

	req_time := time.Since(start).Seconds()

	body, _ := ioutil.ReadAll(resp.Body)

	body_len := len(body)

	body_dbg := make([]byte, 0, 0)
	if (expected_len > 0 && expected_len != body_len) {
		body_dbg = body
	}

	ch <- fmt.Sprintf(
		"[%3d] req to url %s took %.2f secs; body len: %d%s",
		n,
		url,
		req_time,
		len(body),
		body_dbg,
	)
}

func main() {

	args_num := len(os.Args)

	if args_num < 2 {
		fmt.Printf("\ntoo few arguments. At least 2 expected, %d given\n", args_num)
		os.Exit(1)
	}

	expected_len := 0

	url := os.Args[1]

	ch := make(chan string)

	reps := 1
	for i := range os.Args {
		if i == 2 {
			reps, _ = strconv.Atoi(os.Args[i])
		}

		if i == 3 {
			expected_len, _ = strconv.Atoi(os.Args[i])
		}
	}

	fmt.Printf("\ntesting %s %d time(s)\n", url, reps)

	for n := 0; n < reps; n++ {
		go DoRequest(n, url, ch, expected_len)
	}

	for nn := 0; nn < reps; nn++ {
		fmt.Println(<-ch)
	}
}
