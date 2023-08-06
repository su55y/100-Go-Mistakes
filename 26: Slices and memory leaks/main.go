package main

import (
	"fmt"
	"runtime"
	"time"
)

var (
	messages  [][]byte = make([][]byte, 0, 1000)
	nextIndex int
)

const messagesLimit = 1000

func receiveMessage(n int) []byte {
	randomBytes := make([]byte, n)
	b := uint8(0)
	for i := 0; i < n; i++ {
		randomBytes[i] = b
		b++
	}
	return randomBytes
}

func consumeMessages(limit int) {
	for i := 0; i < limit; i++ {
		msg := receiveMessage(1_000_000)

		// storeMessageType(getMessageTypeOld(msg))
		storeMessageType(getMessageType(msg))

		printAlloc()
		time.Sleep(50 * time.Millisecond)
	}
}

// returns backing array pointer with full len
func getMessageTypeOld(msg []byte) []byte {
	return msg[:5]
}

// returns copy of first 5 elements from backing array
func getMessageType(msg []byte) []byte {
	msgType := make([]byte, 5)
	copy(msgType, msg)
	return msgType
}

func storeMessageType(msg []byte) {
	if len(messages) < messagesLimit {
		messages = append(messages, msg)
	} else {
		messages[nextIndex] = msg
	}

	nextIndex = (nextIndex + 1) % messagesLimit
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("\r\033[K%d KB", m.Alloc/1024)
}

func main() {
	fmt.Println()
	printAlloc()
	fmt.Println()
	consumeMessages(1000)
	fmt.Println()
}
