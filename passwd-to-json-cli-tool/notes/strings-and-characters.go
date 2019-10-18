// UTF 8 by default, can use non-ASCII characters inside strings
//

package strings

import "fmt"

func main() {

}

func strings() {
	fmt.Println("Hello!")
	fmt.Println(`
	Multi 
	Line
	String
	it can contain "quotes"!`)
	fmt.Println("\u2272")

}

func numbers() {
	fmt.Println("Addition:", 1+3)
	fmt.Println("Substraction", 11-2)
	fmt.Println("Multiplication", 5*3)
	fmt.Println("Division", 50/2)
}

func vars() {
	var myInt int = 10
	var val, ok = "yes", true
	// val = yes
	// ok = true
	fmt.Println(ok, val, myInt)
}

func arrays(){
	names := [3]string{"Mike", "Lawl", "Jenkins"}
  fmt.Println(names[])
	fmt.Println(names[1])
	

}
