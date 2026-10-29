package app

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func Basic(site string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(site)
	client := http.Client{
		Timeout: 3 * time.Second,
	}
	resp, err := client.Head(site)

	if err != nil {
		fmt.Println("http error")
		return
	} else {
		x := resp.Status
		fmt.Println("Http Status:" + x)
	}
}
