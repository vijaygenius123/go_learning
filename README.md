# Introduction To Go

## GO Charcacteristics

- Compiled
- Statically Typed
- Object Oriented
- No Inheritance
- No Method Or Operator Overloading
- No Implicit Numeric Conversions

## Syntax Rules

- Case Sensitive
- Variables & Package Names Are Lower And Mixed Case
- Exported Functions Have Initial Uppercase 
- Semicolons Arent Required
- Code Blocks Within Braces

## Strcture Of Go Program

```go
package main <- Name Of Package

import (    
	"fmt"     <- Import Required Packages
)

func main(){     <- Function Header main Is The function That Gets Called
	fmt.Println("Hello World")   <- Function Body 
}
```

```go
fmt.Println("Hello World") -> From Fmt Package Use Println To Print On Console
```

To Run The Program
```bash
go run <FileName.go> <- To Run The Program
go build <FileName.go> <- To Build A Executable
./FileName <- Run The Executable
```