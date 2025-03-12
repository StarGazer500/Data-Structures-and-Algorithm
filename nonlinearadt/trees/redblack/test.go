package redblack

import (
	"fmt"
	"strconv"
)


type Integer int
func (i Integer) String() string {
    return strconv.Itoa(int(i))
}

func TestRedBlack(){
	myTree := NewTree[Integer](10)
    myTree.Insert(20)
    myTree.Insert(4)
    myTree.Insert(15)
    myTree.Insert(17)
    myTree.Insert(40)
    myTree.Insert(50)
    myTree.Insert(60)
    myTree.Insert(70)
    myTree.Insert(35)
    myTree.Insert(38)
	fmt.Println(myTree.Search(31))
	fmt.Println(myTree.Search(35))

	myTree.Delete(35)
	fmt.Println(myTree.Search(35))

}