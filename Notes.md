# Learning Golang

### `Reference Material`: https://www.youtube.com/playlist?list=PL4cUxeGkcC9gC88BEo9czgyS72A3doDeM

## Lesson 1 - Your First Go File
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

```
To run file: `$ go run main.go`

## Lesson 2 - Variables, String & Numbers
```go
package main

import "fmt"

func main() {

	// strings
	// var variable_name Type
	var nameOne string = "Tim"
	var nameTwo = "Joe" // Go implies the type
	var nameThree string //no value assigned

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "Jim"
	nameThree = "Max"

	fmt.Println(nameOne, nameTwo, nameThree)

	// initialize without var keyword
	nameFour := "Tom"
	fmt.Println(nameFour)

	// integers
	var ageOne int = 30
	var ageTwo = 20
	ageThree := 10

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory
	var numOne int8 = 25 // 8 bit int (-128 to 127)
	var numTwo int16 = 250 // 16 bit int (-32768 through 32767)
	var numThree uint = 25 // unsigned int(can't have negative numbers 0 to 255)
	
	// floats
	var scoreOne float32 = 7.25
	var scoreTwo float64 = 433398.35 // using this most
	scoreThree := 1.5 // defaults to float64

	// Documentation: https://go.dev/ref/spec#Numeric_types
}

```

## Lesson 3 - Printing & Formatting Strings
```go
package main

import "fmt"

func main() {

	age := 25
	name := "Tim"

	// Print
	fmt.Print("Hello, ")
	fmt.Print("World!")
	// Output: Hello, World!
	

	//Println
	fmt.Println("Heeloo")
	fmt.Println("Goodbyeee")

	fmt.Println("my name is", name, "and my age is", age)
	// Output: my name is Tim and my age is 25


	// Printf (Formatted Strings) %_ = format specifier
	fmt.Printf("my name is %v and my age is %v \n", name, age) //%v variable in order

	fmt.Printf("my name is %q and my age is %q \n", name, age) //%q put quotes on variables. (works only on strings)

	fmt.Printf("age is of type %T \n", age) //%T gets the data type of the variable

	fmt.Printf("you scored %f points! \n", 225.5) //%f for floats

	fmt.Printf("you scored %0.1f points! \n", 225.55) //%0.{x}f where {x} is the number of decimals


	// Sprintf (save formated strings)
	var str = fmt.Sprintf("my name is %v and my age is %v \n", name, age)
	fmt.Println("Saved string is:", str)

}

```

## Lesson 4 - Arrays & Slices
```go
package main

import "fmt"

func main() {
	// Arrays (are fixed size)
	var ages [3]int = [3]int{13, 23, 37}
	var ages2 = [3]int{20, 25, 30}

	fmt.Println(ages)
	fmt.Println(len(ages))

	names := [3]string{"Tim", "John", "Joe"}
	names[1] = "Jane"

	fmt.Println(names)
	fmt.Println(len(names))

	// Slices (use arrays under the hood)
	// Slices are not fixed size
	var scores = []int{100, 50, 70}
	scores[0] = 17
	append(scores, 85) // returns new slice rather than change slice
	scores = append(scores, 75) // replaces original slice

	fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)

	rangeOne = append(rangeOne, "Max")
	fmt.Println(rangeOne)
}

```

## Lesson 5 - The Standard Library
```go
package main

import "fmt"

func main() {
}

```

## Lesson 6 - Loops
```go
package main

import "fmt"

func main() {
}

```

## Lesson 7 - Booleans & Conditionals
```go
package main

import "fmt"

func main() {
}

```

## Lesson 8 - Using Functions
```go
package main

import "fmt"

func main() {
}

```

## Lesson 9 - Multiple return values
```go
package main

import "fmt"

func main() {
}

```

## Lesson 10 - Package scope
```go
package main

import "fmt"

func main() {
}

```

## Lesson 11 - Maps
```go
package main

import "fmt"

func main() {
}

```

## Lesson 12 - Pass by Value
```go
package main

import "fmt"

func main() {
}

```

## Lesson 13 - Pointers
```go
package main

import "fmt"

func main() {
}

```

## Lesson 14 - Structs and Custom Types
```go
package main

import "fmt"

func main() {
}

```

## Lesson 15 - Receiver Functions
```go
package main

import "fmt"

func main() {
}

```

## Lesson 16 - Receiver Functions with Pointers
```go
package main

import "fmt"

func main() {
}

```

## Lesson 17 - User Inputs
```go
package main

import "fmt"

func main() {
}

```

## Lesson 18 - Switch Statements
```go
package main

import "fmt"

func main() {
}

```

## Lesson 19 - Parsing Floats
```go
package main

import "fmt"

func main() {
}

```

## Lesson 20 - Saving Files
```go
package main

import "fmt"

func main() {
}

```

## Lesson 21 - Interfaces
```go
package main

import "fmt"

func main() {
}

```