package main

import (
	"fmt"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func basicAuth(config Config) string {
	return base64.StdEncoding.EncodeToString([]byte(config.BasicAuth))
}

func doRequest(config Config, url string, n int, ch chan<- string, expected_len int) {
	start := time.Now()

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Printf("\nNewRequest error: %s\n", err)
	}

	req.Header.Add("Authorization", "Basic " + basicAuth(config))

	resp, _ := http.DefaultClient.Do(req)

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


	fmt.Printf("\nArgs: %s\n", os.Args)

	args_num := len(os.Args)

	if args_num < 2 {
		fmt.Printf("\ntoo few arguments. At least 2 expected, %d given\n", args_num)
		os.Exit(1)
	}

	expected_len := 0

	url := os.Args[1]

	configFileName := "./config.json"

	reps := 1
	for i := range os.Args {
		if i == 2 {
			reps, _ = strconv.Atoi(os.Args[i])
		}

		if i == 3 {
			expected_len, _ = strconv.Atoi(os.Args[i])
		}

		if strings.HasPrefix(os.Args[i], "--config=") && len(os.Args[i]) > 9 {
			configFileName = os.Args[i][9:]
		}
	}

	fmt.Printf("configFileName: %s ", configFileName)
	fmt.Printf("\ntesting %s %d time(s)\n", url, reps)

	config := GetConfig(configFileName)
	ch := make(chan string)

	for n := 0; n < reps; n++ {
		go doRequest(config, url, n, ch, expected_len)
	}

	for nn := 0; nn < reps; nn++ {
		fmt.Println(<-ch)
	}
}
