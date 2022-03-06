package main

import "fmt"

func main() {
	a1 := [3]string{}
	a1[0] = "Red"
	a1[1] = "Blue"
	a1[2] = "Green"
	pokes := [...]string{"Pika", "Zeni", "Hito", "Fushi"}

	fmt.Println(a1)
	fmt.Println(pokes)

	// Slice 少しパフォーマンスは落ちるらしい？　動的配列的な？
	monsters := []string{}
	monsters = append(monsters, "スライム")
	monsters = append(monsters, "ドラゴン")
	monsters = append(monsters, "ドラキー")
	fmt.Println(monsters)

	a := []int{}

	for i := 0; i < 10; i++ {
		a = append(a, i)
		fmt.Println(len(a), cap(a))
	}

	// map 連想配列的な map[key type]value type
	mapvar := map[string]int{
		"x": 100,
		"y": 200,
	}

	fmt.Println(mapvar)
	fmt.Println("mapvar[x] is ", mapvar["x"])

	mapvar["z"] = 300
	fmt.Println("mapvar[z] is ", mapvar["z"])

	fmt.Println("mapvar length is ", len(mapvar))

	for key, value := range mapvar {
		fmt.Printf("%s = %d\n", key, value)
	}
}
