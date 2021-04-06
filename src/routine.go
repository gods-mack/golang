package main
import ("fmt"
	"time"
)


func main() {
//	fmt.Println("hello")
	go producer(1) 
	
	go producer(2)

	go func () {
		for i:=0; i < 3; i++ {
			fmt.Println("Foo() ")
		}
		
		go producer(3)

	}()


	time.Sleep(time.Millisecond)
}

func producer(id int) {
	for i:=0; i < 5; i++ {
		fmt.Printf("id : %d i: %d\n",id,i)
	}
}
