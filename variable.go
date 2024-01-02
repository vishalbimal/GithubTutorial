/* Types of Variable
- int : stores integers (123,-123)
- float32: stores floating point numbers, with decimals 
such as 19.99 or -19.99
- string : stores texts
- bool : stores values with two states, true or false.

## Declaring Variables
*/
package main
import ("fmt")

func main() {
  var a bool 
  a= true    // Boolean
  var b int  
  b= 25    // Integer
  var c float32 
  c=25.989  // Floating point number
  
  d:= "Boolean" // String

  fmt.Println("Boolean: ", a)
  fmt.Println("Integer: ", b)
  fmt.Println("Float:   ", c)
  fmt.Println("String:  ", d)
}