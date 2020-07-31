package main

import "fmt"

func printSlice(str string, slice []int) {
	fmt.Printf("str:%s, len=%d, cap=%d %v\n",
		str, len(slice), cap(slice), slice)

}

func main()  {
	p := []int{2, 4, 5, 22, 1, 123}
	fmt.Println("p:", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d]:%d\n", i, p[i])
	}
	// left close, right open
	fmt.Println("p[:3]:", p[:3]) // 0 1 2
	fmt.Println("p[4:]:", p[4:]) // 4 5

	// slice 的几种定义方式，和 std::vector 差不多
	a := make([]int, 5) // 自动填充零值，len=cap=5
	printSlice("a", a)
	b := make([]int, 0, 5) // 仅定义容量
	printSlice("b", b)
	c := b[:2] // 从其他 slice 中截取 (左闭右开)
	printSlice("c", c)
	d := c [2:5] // 从其他 slice 中截取 (左闭右开)
	printSlice("d", d)
}
