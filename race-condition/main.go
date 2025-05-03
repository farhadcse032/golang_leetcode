package main
import (
	"fmt"
	"time"
)

var mount = 0

func Add(amount int) {
	mount = mount + amount
}

func main() {
	for i := 0; i < 5; i++ {

		go Add(20)
		go Add(-10)
	}
	time.Sleep(time.Second * 1)
	fmt.Println(mount)
}
