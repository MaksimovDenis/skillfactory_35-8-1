package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

const addr = "localhost:8080"
const network = "tcp4"

var data = [19]string{"Don't communicate by sharing memory, share memory by communicating.",
	"Concurrency is not parallelism.",
	"Channels orchestrate; mutexes serialize.",
	"The bigger the interface, the weaker the abstraction.",
	"Make the zero value useful.",
	"interface{} says nothing.",
	"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.",
	"A little copying is better than a little dependency.",
	"Syscall must always be guarded with build tags.",
	"Cgo must always be guarded with build tags.",
	"Cgo is not Go.",
	"With the unsafe package there are no guarantees.",
	"Clear is better than clever.",
	"Reflection is never clear.",
	"Errors are values.",
	"Don't just check errors, handle them gracefully.",
	"Design the architecture, name the components, document the details.",
	"Documentation is for users.",
	"Don't panic."}

func main() {

	conn, err := net.Dial(network, addr)
	if err != nil {
		log.Fatalf("Failed to connect %v", err)
	}
	defer conn.Close()

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			_, err = conn.Write([]byte(randomProverbs()))
			if err != nil {
				log.Fatalf("Failed to write messsage %v", err)
			}

			reader := bufio.NewReader(conn)
			b, err := reader.ReadBytes('\n')
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Ответ от сервера:", string(b))

		}
	}
}

func randomProverbs() string {

	idx := rand.Intn(19)

	return data[idx] + "\n"
}
