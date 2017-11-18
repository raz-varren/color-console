[![GoDoc](https://godoc.org/github.com/raz-varren/color-console?status.svg)](https://godoc.org/github.com/raz-varren/color-console "Color-Console documentation")

Color-Console
=============

Color-Console is a simple library for printing ANSI colored text to an io.Writer, usually stdout.

#### Example:
```go
package main

import(
	"github.com/raz-varren/color-console"
	"os"
)

func main(){
	cc.Println(cc.ACRed, "this is red", "as is this")
	cc.Printf(cc.ACBlue, "this is %s red\n", "not")
	
	cc.Print(cc.ACLightGreen, "colors")
	cc.Print(cc.ACLightPurple, " are")
	cc.Print(cc.ACYellow, " fun\n")
	
	f, err := os.OpenFile("/tmp/colortest.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	
	cc2 := cc.NewCC(f)
	cc2.Println(cc.ACRed, "this is still red,", "as is this")
	cc2.Printf(cc.ACBlue, "this is %s %s red\n", "still", "not")
	
	cc2.Print(cc.ACLightGreen, "colors")
	cc2.Print(cc.ACLightPurple, " are")
	cc2.Print(cc.ACCyan, " still")
	cc2.Print(cc.ACYellow, " fun\n")
	
	cc.Fprint(cc.ACYellow, f, "frpint test\n")
	cc.Fprintf(cc.ACRed, f, "fprinf %s\n", "test")
	cc.Fprintln(cc.ACLightGreen, f, "fprinln test")
}
```