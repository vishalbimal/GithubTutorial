/* There are two types of constants
1. Typed Constant 
2. Untyped Constant

# TYPED CONSTANT
- Typed Constant are declared with a defined type:*/


// package main 
// import "fmt"

// const A int = 10

// func main(){
// 	fmt.Println(A)
// }

/*

# UNTYPED CONATANT
- Untyped constant is declared without a type
*/
// package main
// import "fmt"
// const A =1

// func main(){
// 	fmt.Println(A)
// }

/*Multiple Constants Declaration
- Multiple constants can be grouped together into 
a block for readability*/

package main
import "fmt"

const (
	A int = 1
	B =3.24
	C = "Hi!"
)
func main(){
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}

