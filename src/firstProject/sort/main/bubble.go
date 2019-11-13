package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	var n int
	var max int
	rand.Seed(time.Now().UnixNano())
	fmt.Println("请输入数组的长度")
	fmt.Scanf("%d" , &n )
	fmt.Println("请输入范围大小")
	fmt.Scanf("%d" , &max)
	arr := make([]int , n)
	for i := 0; i < n ; i++ {
		arr[i] = rand.Intn(max)
	}
	bubble(arr)


	fmt.Println("输出数组")
	fmt.Println(arr)

}


func bubble(arr []int)  {
	length := len(arr)
	for i := 0 ; i < length - 1 ; i++ {
		for j := 0 ; j < length - i - 1 ; j++ {
			if arr[j] < arr[j+1] {
				arr[j] ,arr[j+1] = arr[j+1] , arr[j]
			}
		}
	}

}
