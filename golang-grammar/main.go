/*
A Go program can consist of multiple tokens, which can be:
  * keywords
		break	  default	    func	 interface	   select
		case	  defer	        go	     map	       structure
		chan	  else	        go to	 package	   switch
		const	  fallthrough	if	     range	type
		continue  for	        import	 return	var

  * identifiers
	Identifiers are used to name program entities such as variables and types.
	An identifier is actually a sequence of one or more letters (A~Z and a~z), numbers (0~9), and underscore_,
	but the first character must be a letter or an underscore instead of a number.

  * predefined identifiers
	append	bool	byte	 cap	    close	 complex	complex64	complex128	uint16
	copy	false	float32	 float64	imag	 int	    int8	    int16	    uint32
	int32	int64	iota	 len	    make     new	    nil	        panic	    uint64
	print	println	real	 recover	string	 true	    uint	    uint8	    uintptr
  * constants
  * strings
  * symbols
*/

package main

import (
	"fmt"
	"os"
	"strings"
)

type point struct {
	x, y int
}

func main() {
	fmt.Println("Google" + ".com") // In a Go program, a line represents the end of a statement.
	fmt.Println("Hello, World!")

	name := "John Doe"
	age := 27

	fmt.Printf("hello, %s\n", name)
	fmt.Printf("hello, %s\n", strings.ToUpper(name))
	fmt.Printf("%s is %d years old\n", name, age)

	p := point{1, 2}
	fmt.Printf("%v\n", p)  // Output by value's native value
	fmt.Printf("%+v\n", p) // On the basis of %v, expand the structure field name and value
	fmt.Printf("%#v\n", p) // Output the value in Go language syntax format
	fmt.Printf("%T\n", p)  // Output types and values ​​in Go language syntax format
	fmt.Printf("%t\n", true)
	fmt.Printf("%d\n", 123)
	fmt.Printf("%b\n", 14)
	fmt.Printf("%c\n", 33)
	fmt.Printf("%x\n", 456)
	fmt.Printf("%f\n", 78.9)
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)
	fmt.Printf("%s\n", "\"string\"")
	fmt.Printf("%q\n", "\"string\"")
	fmt.Printf("%x\n", "hex this")
	fmt.Printf("%p\n", &p) // pointer
	fmt.Printf("|%6d|%6d|\n", 12, 345)
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
	fmt.Printf("|%6s|%6s|\n", "foo", "b")
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error")

}
