package Carl_Banner

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var Banner = "Modules/Banner/Banner.txt"

func Out() {
	f, x := os.Open(Banner)
	if x != nil {
		log.Fatal(x)
	} else {
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}
}
