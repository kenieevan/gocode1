package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {

	start := time.Now()
	secs := time.Since(start).Seconds()
	for true {
		secs = time.Since(start).Seconds()
		if secs > 4 {
			break
		}
	}
	// if strings.HasPrefix(url, "http://") != true {
	//	 url = "http://" + url
	// }
	// resp, err := http.Get(url)
	// if err != nil {
	//    ch <- fmt.Sprint(err)
	//	return
	//	}
	//
	// nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	// resp.Body.Close()
	// if err != nil {
	//	  ch <- fmt.Sprintf("while reading %s, %v", url, err)
	//    return
	// }
	//
	//ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
	ch <- fmt.Sprintf("%.2fs", secs)
}
