package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"bufio"
)

var field Field

func main() {
	game(0, -1, -1)
}

func game(n, y, x int) {
	printField(field)
	time.Sleep(500*time.Millisecond)
	resY, resX := fetch(calcTurn(n), y, x)
	field[resY][resX] = calcColor(n)
	game(n+1, resY, resX)
}

func calcTurn(n int) string {
	if n % 2 == 0 {
		fmt.Println("A")
		return os.Args[1]
	} else {
		fmt.Println("B")
		return os.Args[2]
	}
}

func calcColor(n int) State {
	if n % 2 == 0 {
		return Black
	} else {
		return White
	}
}

func fetch(url string, y, x int) (int, int) {
	fmt.Println(url)
	fmt.Printf("%d %d\n", y, x)
	res, err := http.Post(url, "text/plain", strings.NewReader(fmt.Sprintf("%d %d\n", y, x)))
	if err != nil {
		fmt.Println(err)
		return 0, 0
	}

	line, _, err := bufio.NewReader(res.Body).ReadLine()
	if err != nil {
		fmt.Println(err)
		return 0, 0
	}
	var resY, resX int
	fmt.Sscanf(string(line), "%d %d", &resY, &resX)
	return resY, resX
}
