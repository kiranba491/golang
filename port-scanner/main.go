package main

import (
	"flag"
	"log"
	"net"
	"strconv"
	"sync"
)

func checkPort(target string, wg *sync.WaitGroup) {
	defer wg.Done()
	conn, err := net.Dial("tcp", target)
	if err != nil {
		return
	}
	log.Printf("%s is open\n", target)
	conn.Close()
}

func main() {
	var target_host = flag.String("target", "", "Provide valid target host address.")
	flag.Parse()
	if *target_host == "" {
		log.Fatal("--target <string> option required.")
		return
	}
	var wg sync.WaitGroup
	for port := 0; port < 2<<16; port++ {
		wg.Add(1)
		go checkPort(*target_host+":"+strconv.Itoa(port), &wg)
	}
	wg.Wait()
}
