package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		x11      int
		y11      int
		z11      int
		c11      bool
		python11 bool
		java11   bool
	)
	fmt.Println("preview", x11, y11, z11, c11, python11, java11)

	//
	a := "Hello World"
	fmt.Printf("a => %c\n", a[1])

	//
	b := 2
	fmt.Printf("b => %d\n", b)

	//
	for c := 0; c < 3; c++ {
		fmt.Printf("c => %d\n", c)
	}

	//
	d := 0
	for d < 3 {
		d++
		fmt.Printf("d => %d\n", d)
	}

	// Switch 是比 if 更有閱讀性的控制結構，跟 C 語言不同的是他不需要「 break 」來跳出
	e := 5
	choose(e)

	//
	var f [5]int
	f[4] = 100
	fmt.Printf("f => %d\n", f[4])

	//
	var g_list [5]float64
	g_list[2] = 999
	var g float64 = 0
	for _, value := range g_list {
		g += value
	}
	fmt.Printf("g => %f\n", float64(g))

	//
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println("h => ", slice1, slice2)

	//
	i := make(map[string]int)
	i["key"] = 10
	fmt.Println("i => ", i["key"])
	delete(i, "key")
	fmt.Println("i => ", i["key"])

	//
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
	}
	if el, ok := elements["H"]; ok {
		fmt.Println("j => ", el["name"], el["state"])
	}

	//
	k1, k2 := getMultiValues(1, 3)
	fmt.Println("k => k1:", k1, " k2:", k2)

	go funcForDemoGorountine(3)
	time.Sleep(time.Second * 1) // 暫停一秒鐘

	messages := make(chan string)
	go func() { messages <- "ping" }()
	msg := <-messages
	fmt.Println("l => ", msg)

}

func choose(i int) {
	switch i {
	case 1:
		fmt.Println("e => 1")
	case 2, 3:
		fmt.Println("e => 2 or 3")
	case 4:
		fallthrough
	case 5:
		fmt.Println("e => 5")
	}
}

func getMultiValues(i, j int) (int, int) {
	return i, j
}

func funcForDemoGorountine(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println("Gorountine => ", n, ":", i)
	}
}
