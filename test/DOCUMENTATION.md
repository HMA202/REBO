# Go Language Documentation

This file is a beginner-friendly documentation for the Go programming language.
It is useful for learning Go basics and for solving Reboot / Piscine exercises.

---

## 1. What is Go?

Go, also called Golang, is a programming language created by Google.

Go is known for being:

* Simple
* Fast
* Easy to read
* Strongly typed
* Good for backend systems
* Good for APIs and servers
* Good for command-line tools
* Good for cloud and DevOps tools

Go files end with:

```go
filename.go
```

Example:

```txt
main.go
printstr.go
atoi.go
sortintegertable.go
```

---

## 2. Basic Go Program

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")
}
```

### Explanation

```go
package main
```

Defines the package name.

`main` means this file can run as a program.

```go
import "fmt"
```

Imports the `fmt` package.

`fmt` is used for printing and formatting text.

```go
func main() {
}
```

The `main` function is the starting point of the program.

```go
fmt.Println("Hello, Go!")
```

Prints text to the terminal.

Output:

```txt
Hello, Go!
```

---

## 3. Packages

Every Go file must start with a package.

Example:

```go
package main
```

or:

```go
package piscine
```

### package main

Use `package main` when you want to run the file as a complete program.

Example:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello")
}
```

You can run it with:

```bash
go run .
```

or:

```bash
go run main.go
```

### package piscine

In Reboot / Piscine exercises, you usually use:

```go
package piscine
```

because you are writing functions for tests, not a full program.

Example:

```go
package piscine

func Add(a int, b int) int {
	return a + b
}
```

You do not run this file directly unless you create another `main.go` to call the function.

---

## 4. Imports

Imports allow you to use code from another package.

Example:

```go
import "fmt"
```

Then you can use:

```go
fmt.Println("Hello")
```

Multiple imports:

```go
import (
	"fmt"
	"strings"
)
```

In Piscine exercises, sometimes `fmt` is not allowed.
If `fmt` is not allowed, use the allowed function such as:

```go
import "github.com/01-edu/z01"
```

Then print characters using:

```go
z01.PrintRune('A')
```

---

## 5. Functions

A function is a reusable block of code.

Basic function:

```go
func SayHello() {
	fmt.Println("Hello")
}
```

Function with parameters:

```go
func Add(a int, b int) int {
	return a + b
}
```

Explanation:

```go
func Add(a int, b int) int
```

* `func` means function
* `Add` is the function name
* `a int` is the first parameter
* `b int` is the second parameter
* the last `int` means the function returns an integer

Usage:

```go
result := Add(5, 3)
fmt.Println(result)
```

Output:

```txt
8
```

### Function Without Return

```go
func PrintName(name string) {
	fmt.Println(name)
}
```

Usage:

```go
PrintName("Habib")
```

### Function With Return

```go
func Double(n int) int {
	return n * 2
}
```

Usage:

```go
x := Double(4)
fmt.Println(x)
```

Output:

```txt
8
```

---

## 6. Variables

Variables store values.

Long way:

```go
var name string = "Ali"
var age int = 20
```

Short way:

```go
name := "Ali"
age := 20
```

Important:

```go
:=
```

can only be used inside functions.

Correct:

```go
func main() {
	name := "Ali"
	fmt.Println(name)
}
```

Wrong:

```go
name := "Ali"

func main() {
	fmt.Println(name)
}
```

Outside functions, use:

```go
var name string = "Ali"
```

---

## 7. Data Types

Go has many built-in data types.
A data type tells Go what kind of value a variable can store.

---

### Common Go Data Types

```go
string  // text
int     // integer number, can be negative or positive
uint    // unsigned integer, positive number only or zero
bool    // true or false
rune    // character, alias for int32
byte    // character as number, alias for uint8
float64 // decimal number
```

Example:

```go
var name string = "Habib"
var age int = 25
var active bool = true
```

---

### string

`string` is used for text.

```go
var name string = "Habib"
```

Short way:

```go
name := "Habib"
```

Example:

```go
fmt.Println(name)
```

Output:

```txt
Habib
```

---

### bool

`bool` means true or false.

```go
var active bool = true
var finished bool = false
```

Example:

```go
if active {
	fmt.Println("User is active")
}
```

---

### int

`int` is used for normal whole numbers.

It can store negative and positive numbers.

```go
var x int = -10
var y int = 50
var z int = 0
```

Use `int` most of the time for normal numbers.

---

### uint

`uint` means unsigned integer.

Unsigned means the number cannot be negative.

```go
var age uint = 25
var score uint = 100
```

`uint` accepts:

```txt
0
1
2
100
500
```

`uint` does not accept:

```txt
-1
-10
-500
```

Wrong:

```go
var number uint = -5
```

This gives an error because `uint` cannot store negative values.

### Difference Between int and uint

```go
var a int = -10
var b uint = 10
```

`int` can be:

```txt
-10, -1, 0, 1, 10
```

`uint` can be:

```txt
0, 1, 10, 100
```

But `uint` cannot be:

```txt
-1, -10
```

### Loop With uint

If a variable is `uint`, make the loop variable `uint` too.

```go
func PrintNumbers(n uint) {
	for i := uint(0); i < n; i++ {
		fmt.Println(i)
	}
}
```

Usage:

```go
PrintNumbers(5)
```

Output:

```txt
0
1
2
3
4
```

Do not mix `int` and `uint` directly:

```go
func Test(n uint) {
	for i := 0; i < n; i++ {
		fmt.Println(i)
	}
}
```

This causes an error because:

```txt
i is int
n is uint
```

Correct:

```go
func Test(n uint) {
	for i := uint(0); i < n; i++ {
		fmt.Println(i)
	}
}
```

Rule:

```txt
int with int
uint with uint
```

---

### Other Integer Types

Go has different integer sizes.

```go
int8
int16
int32
int64
```

Unsigned versions:

```go
uint8
uint16
uint32
uint64
```

Examples:

```go
var a int8 = 100
var b int16 = 30000
var c int64 = 9000000000

var x uint8 = 255
var y uint16 = 65535
```

Important:

```go
var x uint8 = 256
```

This is wrong because `uint8` can only store from `0` to `255`.

---

### byte

`byte` is an alias for `uint8`.

```go
byte = uint8
```

It is often used for ASCII characters or raw data.

Example:

```go
var ch byte = 'A'
```

Print as number:

```go
fmt.Println(ch)
```

Output:

```txt
65
```

Print as character:

```go
fmt.Printf("%c\n", ch)
```

Output:

```txt
A
```

Why?

Because in ASCII:

```txt
'A' = 65
'B' = 66
'a' = 97
'0' = 48
```

---

### rune

`rune` is an alias for `int32`.

```go
rune = int32
```

It is used for characters, especially Unicode characters like Arabic letters.

Example:

```go
var letter rune = 'ح'
```

Print as number:

```go
fmt.Println(letter)
```

Print as character:

```go
fmt.Printf("%c\n", letter)
```

Output:

```txt
1581
ح
```

The first output is the Unicode number.
The second output is the actual character.

---

### Difference Between byte and rune

`byte` is good for simple ASCII characters.

```go
var a byte = 'A'
```

`rune` is better for Unicode characters.

```go
var ar rune = 'ح'
```

Important:

```txt
byte = uint8
rune = int32
```

For English letters, `byte` can work.
For Arabic letters, use `rune`.

---

### float32 and float64

Float means decimal number.

```go
float32
float64
```

Example:

```go
var price float64 = 10.5
var temperature float64 = 36.7
```

Usually, use `float64` because it is more accurate.

---

### complex64 and complex128

Go also supports complex numbers.

```go
complex64
complex128
```

Example:

```go
var c complex64 = 1 + 2i
fmt.Println(c)
```

Output:

```txt
(1+2i)
```

You will not use this much in beginner Piscine exercises.

---

### Print Type of Variable

Use `%T` to print the type.

```go
name := "Habib"
age := 25
price := 9.99

fmt.Printf("%T\n", name)
fmt.Printf("%T\n", age)
fmt.Printf("%T\n", price)
```

Output:

```txt
string
int
float64
```

---

## 8. If Statements

`if` is used for conditions.

Example:

```go
age := 18

if age >= 18 {
	fmt.Println("Adult")
} else {
	fmt.Println("Minor")
}
```

Output:

```txt
Adult
```

### if without else

```go
n := 5

if n > 0 {
	fmt.Println("Positive")
}
```

### if, else if, else

```go
n := 0

if n > 0 {
	fmt.Println("Positive")
} else if n < 0 {
	fmt.Println("Negative")
} else {
	fmt.Println("Zero")
}
```

---

## 9. Comparison Operators

Comparison operators are used in conditions.

```go
==  // equal
!=  // not equal
>   // greater than
<   // less than
>=  // greater than or equal
<=  // less than or equal
```

Example:

```go
x := 10
y := 20

fmt.Println(x == y)
fmt.Println(x != y)
fmt.Println(x < y)
```

Output:

```txt
false
true
true
```

---

## 10. Logical Operators

Logical operators combine conditions.

```go
&& // and
|| // or
!  // not
```

Example:

```go
age := 20
hasID := true

if age >= 18 && hasID {
	fmt.Println("Allowed")
}
```

Example:

```go
isFriday := true
isHoliday := false

if isFriday || isHoliday {
	fmt.Println("Rest day")
}
```

Example:

```go
active := false

if !active {
	fmt.Println("Not active")
}
```

---

## 11. Loops

Go uses only `for` loops.

---

### Normal loop

```go
for i := 0; i < 5; i++ {
	fmt.Println(i)
}
```

Output:

```txt
0
1
2
3
4
```

Explanation:

```go
i := 0
```

Start from 0.

```go
i < 5
```

Continue while `i` is less than 5.

```go
i++
```

Increase `i` by 1 after each loop.

---

### While-style loop

Go does not have `while`, but you can write a while-style loop using `for`.

```go
i := 0

for i < 5 {
	fmt.Println(i)
	i++
}
```

Output:

```txt
0
1
2
3
4
```

---

### Infinite loop

```go
for {
	fmt.Println("Running")
}
```

This runs forever unless you stop it or use `break`.

---

### break

`break` stops the loop.

```go
for i := 0; i < 10; i++ {
	if i == 5 {
		break
	}
	fmt.Println(i)
}
```

Output:

```txt
0
1
2
3
4
```

---

### continue

`continue` skips the current loop and moves to the next one.

```go
for i := 0; i < 5; i++ {
	if i == 2 {
		continue
	}
	fmt.Println(i)
}
```

Output:

```txt
0
1
3
4
```

---

## 12. Arrays

An array stores multiple values of the same type.
Array size is fixed.

```go
var nums [3]int

nums[0] = 10
nums[1] = 20
nums[2] = 30
```

Example:

```go
arr := [3]int{1, 2, 3}
fmt.Println(arr)
```

Output:

```txt
[1 2 3]
```

Access values:

```go
fmt.Println(arr[0])
fmt.Println(arr[1])
fmt.Println(arr[2])
```

Output:

```txt
1
2
3
```

Loop through array:

```go
for i := 0; i < len(arr); i++ {
	fmt.Println(arr[i])
}
```

---

## 13. Slices

A slice is like an array, but more flexible.
Its size can grow.

```go
nums := []int{1, 2, 3}
```

Add value:

```go
nums = append(nums, 4)
```

Print length:

```go
fmt.Println(len(nums))
```

Example:

```go
nums := []int{1, 2, 3}
nums = append(nums, 4)

fmt.Println(nums)
fmt.Println(len(nums))
```

Output:

```txt
[1 2 3 4]
4
```

Loop through slice:

```go
for i := 0; i < len(nums); i++ {
	fmt.Println(nums[i])
}
```

Using range:

```go
for index, value := range nums {
	fmt.Println(index, value)
}
```

---

## 14. Strings

A string is text.

```go
word := "hello"
```

Access character by index:

```go
fmt.Println(word[0])
```

This prints the byte value, not the letter.

Output:

```txt
104
```

To print the character:

```go
fmt.Printf("%c\n", word[0])
```

Output:

```txt
h
```

### Length of string

```go
word := "hello"
fmt.Println(len(word))
```

Output:

```txt
5
```

### Loop through string by byte

```go
word := "hello"

for i := 0; i < len(word); i++ {
	fmt.Printf("%c\n", word[i])
}
```

Output:

```txt
h
e
l
l
o
```

### Loop through string by rune

```go
word := "مرحبا"

for _, ch := range word {
	fmt.Printf("%c\n", ch)
}
```

Use `range` for Arabic or Unicode characters.

---

## 15. Runes

A rune represents a character.

```go
var ch rune = 'A'
fmt.Println(ch)
```

This prints:

```txt
65
```

To print the character:

```go
fmt.Printf("%c\n", ch)
```

Output:

```txt
A
```

Arabic example:

```go
var ch rune = 'م'
fmt.Println(ch)
fmt.Printf("%c\n", ch)
```

Output:

```txt
1605
م
```

---

## 16. ASCII

Characters have numeric values.

Example:

```txt
'A' = 65
'B' = 66
'C' = 67

'a' = 97
'b' = 98
'c' = 99

'0' = 48
'1' = 49
'2' = 50
```

Example:

```go
fmt.Println('A')
fmt.Printf("%c\n", 'A')
```

Output:

```txt
65
A
```

### Convert digit to character

```go
n := 5
fmt.Printf("%c\n", '0'+n)
```

Output:

```txt
5
```

In Piscine with `z01`:

```go
z01.PrintRune('0' + rune(n))
```

---

## 17. Pointers

A pointer stores the memory address of a variable.

```go
x := 10
p := &x
```

`&x` means address of `x`.

```go
fmt.Println(p)
fmt.Println(*p)
```

`*p` means value inside that address.

Example:

```go
x := 10
p := &x

fmt.Println(x)
fmt.Println(p)
fmt.Println(*p)
```

### Change value using pointer

```go
func PointOne(n *int) {
	*n = 1
}
```

Usage:

```go
x := 10
PointOne(&x)
fmt.Println(x)
```

Output:

```txt
1
```

Explanation:

```go
&x
```

sends the address of `x`.

```go
*n = 1
```

changes the value inside that address.

---

## 18. Example: DivMod

```go
func DivMod(a int, b int, div *int, mod *int) {
	*div = a / b
	*mod = a % b
}
```

Explanation:

* `a / b` gives division result
* `a % b` gives remainder
* `div *int` means pointer to int
* `mod *int` means pointer to int
* `*div = ...` changes the value of the variable passed by address
* `*mod = ...` changes the value of the variable passed by address

Usage:

```go
a := 13
b := 2
var div int
var mod int

DivMod(a, b, &div, &mod)

fmt.Println(div)
fmt.Println(mod)
```

Output:

```txt
6
1
```

---

## 19. PrintRune

In Piscine exercises, you often use:

```go
import "github.com/01-edu/z01"
```

Then:

```go
z01.PrintRune('A')
```

This prints one character.

Example:

```go
func PrintAlphabet() {
	for ch := 'a'; ch <= 'z'; ch++ {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
```

Output:

```txt
abcdefghijklmnopqrstuvwxyz
```

### Print newline

```go
z01.PrintRune('\n')
```

### Print space

```go
z01.PrintRune(' ')
```

### Print comma

```go
z01.PrintRune(',')
```

---

## 20. Print Numbers Without fmt

In some exercises, `fmt` is not allowed.

To print a number as a character:

```go
z01.PrintRune('0' + rune(n))
```

Example:

```go
n := 5
z01.PrintRune('0' + rune(n))
```

Output:

```txt
5
```

Why?

```txt
'0' = 48
5 = 5
48 + 5 = 53
53 = '5'
```

This works for digits from 0 to 9.

For numbers bigger than 9, you need to split the number into digits.

---

## 21. Example: PrintStr

```go
func PrintStr(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
}
```

Explanation:

```go
for _, ch := range s
```

Loops through each character in the string.

`_` means ignore the index.

`ch` is the current character.

Example:

```go
PrintStr("hello")
```

Output:

```txt
hello
```

---

## 22. Example: StrLen

```go
func StrLen(s string) int {
	count := 0

	for range s {
		count++
	}

	return count
}
```

This counts the number of characters.

Example:

```go
fmt.Println(StrLen("hello"))
```

Output:

```txt
5
```

For Arabic text, `range` is better than `len` if you want to count characters.

---

## 23. Example: Swap

```go
func Swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
```

Usage:

```go
x := 10
y := 20

Swap(&x, &y)

fmt.Println(x)
fmt.Println(y)
```

Output:

```txt
20
10
```

Explanation:

```go
temp := *a
```

Stores the value of `a` temporarily.

```go
*a = *b
```

Puts value of `b` inside `a`.

```go
*b = temp
```

Puts old value of `a` inside `b`.

---

## 24. Example: Atoi

Atoi converts string to int.

Example:

```txt
"123" -> 123
"-45" -> -45
```

Simple version:

```go
func BasicAtoi(s string) int {
	result := 0

	for i := 0; i < len(s); i++ {
		result = result*10 + int(s[i]-'0')
	}

	return result
}
```

Explanation for `"123"`:

```txt
result = 0

result = 0*10 + 1 = 1
result = 1*10 + 2 = 12
result = 12*10 + 3 = 123
```

### Why s[i] - '0'?

Characters have ASCII values.

```txt
'0' = 48
'1' = 49
'2' = 50
'3' = 51
```

So:

```txt
'3' - '0' = 51 - 48 = 3
```

This converts a digit character into a real number.

---

## 25. Sorting

Example: sort slice in ascending order.

```go
func SortIntegerTable(table []int) {
	for i := 0; i < len(table)-1; i++ {
		for j := i + 1; j < len(table); j++ {
			if table[j] < table[i] {
				table[i], table[j] = table[j], table[i]
			}
		}
	}
}
```

Input:

```go
[]int{5, 4, 3, 2, 1}
```

Output:

```go
[]int{1, 2, 3, 4, 5}
```

Explanation:

* The first loop chooses one position.
* The second loop searches the rest of the slice.
* If it finds a smaller number, it swaps them.
* At the end, the slice becomes sorted.

---

## 26. Structs

A struct groups related data together.

Example:

```go
type User struct {
	Name string
	Age  int
}
```

Usage:

```go
u := User{
	Name: "Habib",
	Age:  25,
}

fmt.Println(u.Name)
fmt.Println(u.Age)
```

Output:

```txt
Habib
25
```

Structs are useful when you want to represent objects like users, products, lawyers, bookings, etc.

---

## 27. Maps

A map stores key-value pairs.

Example:

```go
ages := map[string]int{
	"Ali":   20,
	"Habib": 25,
}
```

Access value:

```go
fmt.Println(ages["Habib"])
```

Output:

```txt
25
```

Add new value:

```go
ages["Sara"] = 22
```

Delete value:

```go
delete(ages, "Ali")
```

Check if key exists:

```go
age, exists := ages["Habib"]

if exists {
	fmt.Println(age)
}
```

---

## 28. Constants

Constants are values that do not change.

```go
const Pi = 3.14
const AppName = "My App"
```

Example:

```go
fmt.Println(Pi)
fmt.Println(AppName)
```

You cannot change a constant after creating it.

Wrong:

```go
const x = 10
x = 20
```

---

## 29. Switch

`switch` is used when you have many conditions.

Example:

```go
day := "Friday"

switch day {
case "Friday":
	fmt.Println("Weekend")
case "Saturday":
	fmt.Println("Weekend")
default:
	fmt.Println("Working day")
}
```

Output:

```txt
Weekend
```

---

## 30. Useful Terminal Commands

Run Go project:

```bash
go run .
```

Run specific file:

```bash
go run main.go
```

Format Go code:

```bash
gofmt -w file.go
```

Test package:

```bash
go test
```

Create module:

```bash
go mod init projectname
```

Check Go version:

```bash
go version
```

Build program:

```bash
go build
```

---

## 31. Common Errors

### Error: package command-line-arguments is not a main package

This happens when you try to run a file that does not have:

```go
package main
```

and:

```go
func main() {}
```

Piscine files usually use:

```go
package piscine
```

So you need to test them using another `main.go`.

---

### Error: undefined: fmt

This means you used `fmt` but did not import it.

Wrong:

```go
package main

func main() {
	fmt.Println("Hello")
}
```

Correct:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello")
}
```

---

### Error: mismatched types int and uint

This happens when you compare or calculate `int` with `uint`.

Wrong:

```go
func Test(n uint) {
	for i := 0; i < n; i++ {
		fmt.Println(i)
	}
}
```

Correct:

```go
func Test(n uint) {
	for i := uint(0); i < n; i++ {
		fmt.Println(i)
	}
}
```

---

## 32. Difference Between package main and package piscine

### package main

Used when you want to run the program directly.

```go
package main

func main() {
}
```

### package piscine

Used when you are submitting functions for exercises.

```go
package piscine

func PrintStr(s string) {
}
```

You do not run this directly.
You call it from a separate `main.go`.

---

## 33. Clean Code Rules

Good Go code should be:

* Simple
* Clear
* Formatted using `gofmt`
* Easy to read
* Split into functions
* Without unnecessary code

Example:

```go
func IsPositive(n int) bool {
	return n > 0
}
```

Bad example:

```go
func IsPositive(n int) bool {
	if n > 0 {
		return true
	} else {
		return false
	}
}
```

The first one is cleaner.

---

## 34. Important Go Keywords

```go
package
import
func
var
const
if
else
for
range
return
type
struct
interface
go
defer
select
switch
case
default
break
continue
fallthrough
map
chan
```

---

## 35. Mini Practice Exercises

1. Print alphabet
2. Print reverse alphabet
3. Print digits
4. Check if number is negative
5. Print string
6. Count string length
7. Swap two numbers
8. Convert string to number
9. Sort array
10. Count repeated letters
11. Print numbers from 0 to n
12. Reverse a string
13. Count vowels
14. Find maximum number in slice
15. Find minimum number in slice

---

## 36. Count Repeated Letters Example

This is called Run-Length Encoding.

Input:

```txt
lovvvveee you
```

Output:

```txt
1l1o4v3e 1y1o1u
```

Code:

```go
package main

import "fmt"

func CountLetters(s string) string {
	result := ""
	i := 0

	for i < len(s) {
		if s[i] == ' ' {
			result += " "
			i++
			continue
		}

		count := 1

		for i+count < len(s) && s[i+count] == s[i] {
			count++
		}

		result += fmt.Sprintf("%d%c", count, s[i])
		i += count
	}

	return result
}

func main() {
	text := "lovvvveee you"
	fmt.Println(CountLetters(text))
}
```

Explanation:

```go
result := ""
```

Stores the final text.

```go
i := 0
```

Starts from the first character.

```go
for i < len(s)
```

Loops until the end of the string.

```go
if s[i] == ' '
```

If the current character is a space, add the space without counting it.

```go
count := 1
```

Start counting the current character.

```go
for i+count < len(s) && s[i+count] == s[i]
```

Counts repeated characters after the current character.

```go
result += fmt.Sprintf("%d%c", count, s[i])
```

Creates text like:

```txt
4v
3e
1l
```

```go
i += count
```

Jumps after the repeated characters.

---

## 37. How to Test Piscine Functions

If your exercise file is:

```go
package piscine

func Add(a int, b int) int {
	return a + b
}
```

You cannot run it directly because it does not have `package main`.

Create a separate test program.

Example structure:

```txt
project/
├── go.mod
├── main.go
└── piscine/
    └── add.go
```

Example `main.go`:

```go
package main

import (
	"fmt"
	"project/piscine"
)

func main() {
	fmt.Println(piscine.Add(5, 3))
}
```

Run:

```bash
go run .
```

Output:

```txt
8
```

---

## 38. Important Piscine Notes

In Piscine exercises:

* Read the allowed functions carefully.
* If `fmt` is not allowed, do not use it.
* Use `z01.PrintRune` when required.
* File name must match the exercise requirement.
* Function name must match exactly.
* Package name is usually `piscine`.
* Run `gofmt` before submitting.

Example:

```bash
gofmt -w printstr.go
```

---

## 39. Quick Cheat Sheet

### Print

```go
fmt.Println("Hello")
fmt.Printf("%c\n", 'A')
```

### Variable

```go
name := "Habib"
age := 25
```

### Function

```go
func Add(a int, b int) int {
	return a + b
}
```

### If

```go
if age >= 18 {
	fmt.Println("Adult")
}
```

### Loop

```go
for i := 0; i < 5; i++ {
	fmt.Println(i)
}
```

### String loop

```go
for _, ch := range s {
	fmt.Printf("%c\n", ch)
}
```

### Pointer

```go
x := 10
p := &x
fmt.Println(*p)
```

### Slice

```go
nums := []int{1, 2, 3}
nums = append(nums, 4)
```

### Map

```go
ages := map[string]int{
	"Ali": 20,
}
```

---

## 40. Notes

Go is powerful because it is:

* Fast
* Simple
* Strongly typed
* Good for servers and APIs
* Easy to read
* Good for beginners and professionals

For Piscine, focus first on:

```txt
functions
variables
if
for loops
strings
runes
ASCII
pointers
slices
```

These are the most important topics for beginner Go exercises.
