package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	for i := 1; i < 65535; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			addres := fmt.Sprintf("192.168.101.101:%d", n)
			conn, err := net.Dial("tcp", addres)
			if err != nil {
				//fmt.Printf("%s close...\n", addres)
				return
			}
			conn.Close()
			fmt.Printf("%s open...\n", addres)
		}(i)
	}
	wg.Wait()
	elapsed := time.Since(start) / 1e9
	fmt.Printf("\n\n%d seconds", elapsed)
}
