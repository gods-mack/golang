package main
import ( "fmt" )

func main() {
	fmt.Println("Hello, Go!")
	var x int = 23
	var y int = 42
	sum := x + y
	fmt.Println("sum: ",sum)

	// array
	var arr [5]int
	fmt.Println(arr)
	arr[2] = 9
	fmt.Println(arr)
	
	// array
	arr1 := []int{1,2,4,4}
	fmt.Println(arr1)
	arr1 = append(arr1, 8)
	fmt.Println(arr1)


	// Map / Hashing
	var mpp = make(map[string]int)
	mpp["manis"] = 2423
	mpp["neeraj"] = 2424
	mpp["chachi"] = 2425
	fmt.Println(mpp["manis"],mpp)

	// for loop
	for i:=0; i < 5; i++ {
		fmt.Println(i)
	}
	//while loop
	i:=0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// iterate map
	for key,value := range mpp {
		fmt.Println("key: ",key," value: ",value)
	}

	// pointer and reference
	n := 4
	incr(&n)
	fmt.Println("x: ",n)
	
	// passing array as reference
	parr(&arr1)
        fmt.Println(arr1)

	// working with string
	var s string = "mugambo"
	fmt.Println(s[2])
	for i,j := range s{
		fmt.Printf("%d %c\n",i,j)
	}
}

func incr(num *int){
	*num++
}

func parr(x *[]int) {
	(*x)[2] = 1222
	fmt.Println(*x)
	for i,j := range *x {
		fmt.Println(i,j)
	}
}



// Please enter the commit message for your changes. Lines starting
