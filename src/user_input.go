

package main
import ( "fmt" )

func main() {
	// input n elements in an array
	var n int
	fmt.Scan(&n)
	arr := make([]int, n) 
	for i:=0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println(arr)


	// make array of all unique number
	mp := make(map[int]bool)

	for i:=0; i < n; i++ {
		mp[arr[i]] = true

	}

	// all unique elements
	fmt.Println("all unique elements")
	for key := range mp {
		fmt.Println(key)
	}

}
