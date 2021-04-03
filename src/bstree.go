
package main
import ("fmt")

type node struct {
	data int
	left *node
	right *node
}

type bst struct {
	root *node
}

	

func insert(x int, root **node) {
	//fmt.Println("inserting..")

	if (*root) == nil {
		(*root) = &node{data:x, left:nil, right:nil}
	}

	if (*root).data > x {
		insert(x, &(*root).left)
	} else if (*root).data < x {
		insert(x, &(*root).right)
	}	
}


/*
func (t *bst) insert(x int) {
	//fmt.Println("inserting..")

	if t.root == nil {
		t.root = &node{data:x, left:nil,right:nil}
	} else if t.root.data > x {
		t.root.left.insert(x)
	} else if t.root.data < x {
		t.root.right.insert(x)
	}
}
*/

func inorder(tmp *node) {
	if tmp == nil {
		return
	}
	inorder(tmp.left)
	fmt.Println(tmp.data)
	inorder(tmp.right)
}

func main() {

    tree := bst{root:nil}
   
	
	insert(10, &tree.root)
	insert(19,&tree.root)
	insert(14, &tree.root)
	insert(7, &tree.root)
	insert(5, &tree.root)
	
	
	//t.insert(10)
	//t.insert(19)
	//t.insert(14)


	inorder(tree.root)

}