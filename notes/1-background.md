# GoLang

## Background/ Install stuff

- Workspace used to have three main folders: bin, pkg and src
  
- Method has now changed, instead of using env variables like gopath, use go modules [link](https://stackoverflow.com/questions/10838469/how-to-compile-go-program-consisting-of-multiple-files/61793820#61793820)

### Go commands

- go `fmt` formats code to the conventions of golang. Can do for all files: `go fmt ./..`
- go `run` runs .go files, like python (well, it builds and runs, unlike python ~~ )
- go `build` can build an executable
- go `install` - not sure distinction from build... but compiles a named executable

### Go modules

- default way to manage code, good outline of how to setup: [link](https://blog.golang.org/using-go-modules)
- modules seem similar to python and others where similar functions can be grouped into modules/packages in different folders, or node
- may have whole project in one module, or project pulling in from lots of modules.

### Dependencies

- When add in a dependency, both direct and indirect dependencies may be added. One module may require others etc.
- command `go list -m all` lists current module and all dependencies
- the `go.sum` file stores the cryptographic hashes of the module versions used as dependencies (in the current module). This is useful as ensures anyone else running code can get the exact version you are working with. For this reason, this should also be checked into version control (alongside the go.mod)
- Dependencies can be updated to the most recent version or a specific version. Updating an indirect dependency can cause issues, as the updated version could cause issues in the direct dependency module being used. updated indirect modules will have a comment tagging them as `//indirect`
  - To update a module use `go get xxx`
  - List available versions of a module using `go list -m -versions xxx`
  - Can get a specific version by denoting it using '@': e.g.`rsc.io/sampler@v1.3.1`
  - Two major versions of the same module can be used in the same code. A major version is V1.x vs. V2.x etc.
    - Major versions of Go modules have different module paths.
    - Two minor versions of the same module cannot be used.
    - Allowing usage of two major versions allows developers to migrate incrementally to newer versions (e.g. not mess up tons of legacy code by switching)
- `go doc xxx` is useful for exploring a modules methods
- `go mod tidy` removes any unused dependencies

## Language details

- Like java, each program has a main function. Also each program should have a package main. Main function is entry point of program and controls the flow~~
- You can't declare variables and not use them in Go. If a method returns two values, but only want one, can throw away a value by assigning it to a special var - underscore `_` e.g. `a, _ := myFunc()`
- A 'variatic' parameter is one which can take any number of variables as inputs. In the docs will be denoted as `a ... <type>`
- The type `interface{}` is the empty interface. Everything is of type `interface{}`, so strings, chars, ints, etc. Appears in docs when a function can take any type as an argument

### Packages (inbuilt / standard library)

- like java, python and others, a lot of useful pre-written code for handling different tasks is stored in the packages of the standard library
  - A good website to explore the documentation of packages is [godoc.org](https://pkg.go.dev/std)

### Identifiers and keywords

- like python first letter of an identifier cannot be a number. Anything else other than a reserved keyword should be fine.
- Naming convention is camelCase rather than using underscores
  - Capitalised identifiers allow functions, types or fields to be exported and accessed from other packages [ref](https://golang.org/ref/spec#Exported_identifiers). (As see later with JSON)

### Operators

- "short declaration operator": `:=` . In Golang there are long and short ways of declaring variables. A variable needs to be declared the first time it is used. After declaration a variable can be updated using the normal assignment operator `=`.
  - The short declaration operator both declares and assigns a value to a variable in one line.
- Operators fairly similar to java, using `||` and `&&`, including `++` and `--`, though `++` and `--` are statements. Will have to dive into details see if they have odd behaviour like with Java.
- Golang also has binary/bitwise operators, which perform bit manipulation on the binary versions of numbers. There's binary and `&`, or `|`, XOR `^`and shift left `<<`and shift right `>>`
- Seem to be a few more assignment operators. There's a modulus assignment operator `%=` as well as bitwise assignment operators.

### Declaring and assigning

- As mentioned, can use the short declaration operator.

- Can declare using the `var` keyword. Variables assigned this way can be done so outside of a function body e.g. globally. Generally global variables are bad, so use the short declaration operator instead.

  - Can denote the type when declaring a variable using var. If no value is provided when declaring a variable in this way, the "zero" value is used e.g. `var i int` would assign the value "0" to `i`.

- Generally do not need to specify type (though can be helpful). The type will be inferred from the value of the variable: numbers without decimal points will be ints, those with them will be float64s, etc.

  - Once a variable is declared and a type assigned, you cannot assign a value of a different type to the variable.

    ```go
    a := 23
    a = "clouds are good" // this is invalid
    ```

- **Constants** can be declared. They cannot be declared with the short declarator `:=`. They also have some interesting properties with both **typed** and **untyped** constants: [link](https://blog.golang.org/constants)

  - An untyped constant allows the same value to be assigned to different types (sometimes) e.g.

    ```go
    const a = 1
    
    var b int = a
    var c float64 = a
    
    fmt.Println(b/2, c/2)	// 0 0.5
    ```

    - Even untyped constants do sort of have a type, `a` above is of type `untyped int` and can be converted to several numerical types, however could not necessarily be cast as a string or other types

  - When declaring a set of constants (in parenthesis) the `iota` value can be used. `iota` will take the value of successive integers, `0 1 2` etc. `iota` will start and zero and increment by 1 for each constant. It will increment even if it isn't used in some. `iota` can be mentioned only for the first and the others will follow:

    ```go
    const (
    	c0 = iota 	// 0
    	c1			// 1
    	c2			// 2
    )
    
    // same as
    
    const (
    	c0 = iota 	// 0
    	c1 = iota	// 1
    	c2 = iota	// 2
    )
    
    // will apply to those below, but always start at 0:
    
    const (
    	c0 = 5 		// 5
    	c1 = iota	// 1
    	c2			// 2
    )
    
    // example series using bit shifting
    const (
    	_ = iota
        kb = 1 << (iota * 10)
        mb = 1 << (iota * 10)
        gb = 1 << (iota * 10)
    )
    ```

    - Generally useful when defining a series of constants, such as days of the week or months etc.

### Types

- Go is statically typed language - once assigned to a type, a variable cannot take another type

- A type can be probed in printout, using `%T`:

  ```go
  a := 2
  fmt.Printf("%T\n", a)
  ```

- Strings and chars are treated a little differently in Golang. In fact instead of ASCII 'chars' Golang uses 'runes' which are `int32` and encode UTF-8 Unicode characters. When probing the value and type of a rune, often returns a integer value:

  ```go
  c := 'a'
  fmt.Printf("%T\n", c)	// int32
  fmt.Println(c)			// 97
  c -= 32					
  fmt.Println(c)			// 65				
  fmt.Println(string(c))	// A
  ```

  - Strings are technically 'read only slices of bytes'. The bytes do not need to be in any specific format, however when using text, strings contain UTF-8 bytes : [link](https://blog.golang.org/strings) [go101](https://go101.org/article/string.html)

- Lots more can go into about types like other languages, are basic types and composite types [link](https://go101.org/article/type-system-overview.html)



### fmt and Printing

https://pkg.go.dev/fmt

- The fmt package  contains a range of input output printing and scanning functions.

- Generally there are two types of printing, normal and 'format' printing. Normal printing (`Print` or `Println`) prints strings as they are, format printing (`Printf`) allows elements to be passed in or the format to be modified

- As well as containing standard output (console) printing methods, fmt also also contains printing methods that output to file (e.g. `Fprintf`) or that return a string (e.g. `Sprint`)

  - e.g.

    ```go
    s := fmt.Sprint("cake is good")
    fmt.Println(s)	//=> cake is good
    ```

- For format printing, there are 'verbs' which modify the output. There are quite a few and are similar to those used in MATLAB, starting with a `%` e.g. the `%T` we saw before. Important general one is `%v` which passes the value. There are quite a few associated with different data types, such as `%c` which prints an `int` as the Unicode character it represents. Also useful are the verbs which allow for easy printing of numbers in octal, binary and hexadecimal:

  ```go
  e:= 26
  // Printing value, hex (with 0x), binary and hex (without 0x))
  fmt.Printf("%v\t%#x\t%b\t%x\n", e, e, e, e)
  ```

  - As with python when using `.format()` the things to insert/format into strings are passed after the string, separated by commas

**NOTE:** You can use the inbuilt `println()` method instead of `fmt.Pintln()` (to save time) when checking values and diagnosing code, however for production/ publishing `fmt.Println` is the correct method [link](https://golang.org/ref/spec#Bootstrapping)

### Creating own types

- Golang has a stronger emphasis on types than many other languages, with an inability to compare different types e.g. you cannot say `2 == 2.0` without converting/casting one of the values so that both are the same.

- You can create custom types (based on existing types) easily. These newly defined types follow the same rules, even if the types they are based on are the same.

  - e.g. a custom type based on an integer cannot be compared with a normal integer:

    ```go
    type counter int
    var z counter = 23
    fmt.Println(z)			// 23
    fmt.Printf("%T/n", z)	// main.counter
    a := 23
    a = z					// cannot use z (type counter) as type int in assiggnment
    ```

### Type conversion

- Golang uses the term 'conversion' instead of casting. While both are pretty similar, the way golang performs type conversion is a little different from the way Java performs typecasting etc. So generally best to stick to their terminology and call it conversion. [link](https://medium.com/@rocketlaunchr.cloud/type-conversions-casting-type-assertions-fb295430e387)
- Type conversion creates a copy of the original value in a new format.
- As usual can cause some loss of data when moving from larger or more precise format to a less precise one (e.g 64 to 32 or signed to unsigned, floats to ints etc.)
  - Though you know can lose information going both ways
  - Type conversion is required even if both types are fundamentally the same, especially true with the created types mentioned before, but also in cases where an `int` may already be stored as an `int64` on a 64-bit machine.
- There is also something called **Type Assertion**. In this process you assert that an object is actually of another type. This is generally used  in situations where two types based on the same interface are in use, or in the case where want to check if an object conforms to certain type [link](https://tour.golang.org/methods/15)
  - Can also be used if want to create objects of several types from a common interface and set of data etc.
  - Type assertion is not casting or conversion. The (stored) values are not modified in the process

### Numeric types

- Similar others, has signed & unsigned: `uint8` (0 to 255), `int8` (-128 to 127) etc. (same for floats)
- No distinction floats and doubles. `float64` is an `IEEE-754` 64 bit floating point
- Also have imaginary/complex numbers, handled by `complex64` and `complex128`
- `byte` is an alias of `uint8` and `rune` is an alias of `int32` (but distinction is useful e.g. characters are stored as runes)

### String type

https://blog.golang.org/strings

- The string type is a little different in golang. String values are defined as sequences of bytes, or **a slice of bytes**.

  - If convert a string to a byte array, can see the sequence of bytes which represent a string

- The length of a string in golang is the number of bytes in that string. As letters/characters are stored in UTF-8 format where a character is between 1 and 4 bytes, this means the length of a string will not necessarily correspond to the number of characters:

  ```go
  a := "english"
  b := "日本語"
  
  fmt.Println(len(a))	// 7
  fmt.Println(len(b))	// 9
  ```

- **Strings are immutable** in golang, they cannot be updated or changed.

- Like other languages elements of a string can be accessed. However in golang you access the bytes of a string, meaning a numerical byte value will be returned:

  ```go
  	a := "english"
  	b := "日本語"
  
  	fmt.Println(a[0])	// 101
  	fmt.Println(b[0])	// 230
  ```

  - To access the characters in a string by index, a string needs to be converted to an array of runes:

    ```go
    b := "日本語"
    c := []rune(b)
    
    fmt.Println(b[0])		// 230
    fmt.Println(c[0])		// 26085
    fmt.Printf("%c\n", c[0])// 日
    ```

  - For many English characters where ASCII rules still apply, converting the bytes from string positions may return the original characters, however this will not work for any char which is represented by 2 or more bytes in UTF-8 coding.

### Loops

- No separate while statement in golang, however the for loop can do rather similar things

- Several ways of writing loops: generally can initialise counting variable outside the loop (before) and increment within, or use the standard method (with a for clause):

  ```go
  // standard for loop (with a for clause)
  for i := 0; i < 5; i++ {
      fmt.Println("yes")
  }
  
  // initialise outside
  j := 0
  for j <= 5 {
  	fmt.Println("yes")
  	j ++
  }
  ```

- The second for loop above is an example of a **"for statement"** (which is more than a little similar to a while loop). A for statement will continue looping as long as a Boolean condition evaluates to true (with the evaluation occurring before each iteration of the loop - like a while loop~):

  ```go
  a := "a"
  for a != "aaaaaaa" {
      a = a + "a"
  }
  
  // or something even more while loopy
  c := true
  for c {
      a = a + "a"
      if a == "aaaaaaaaa" {
          c = false
      }
  }
  ```

- Can also use the `range` method (range clause) to iterate over different data structures.

- Golang allows for loops with no conditions. These can run indefinitely and are can be used in that capacity for web servers, or other things. Can also set the internal logic of a loop with no condition to `break` if a particular event occurs.

### Conditional statements

- All fairly standard

- Can initialise variables in a condition, allowing that variable to have a smaller / more limited scope. Can even use a variable name already in use off larger scope (but not advised ~~)

  ```go
  b := 13
  a := 200
  if b := 100; a > b {
      fmt.Println(b)				// 100
      fmt.Println("this is fine")
  }
  fmt.Println(b)					// 13
  ```

  - Seems like a label can be re-used in places of lower scope and once exiting that scope the original value returns:

    ```go
    b := 13
    {
        b := 23
        fmt.Println(b)		// 23
    }
    fmt.Println(b)			// 13
    
    // this also applies to loops, but again, not advised
    for i := 0; i < 2; i++ {
        for i := 0; i < 2; i++ {
            fmt.Println("This is valid code")
        }
    }
    ```

- Switch statements: No need to provide `break` statements. In go a switch will terminate after finding the first matching condition. This behaviour can be overridden by  providing a `fallthrough` statement on a condition.

  - Golang switch does have a `default` statement (for if none of the conditions are met)
  - Can either have switches with a condition or switches without conditions (where the switch expression defaults to `true`)
  - Can have multiple conditionals on the same `case` statement: just separate by commas
  - It is *idiomatic* to write `if-else` chains as condition-less switch statements 

```go
a := 23
switch {
case a <= 2021:
    fmt.Println("this prints")
    fallthrough
case a == 23:
    fmt.Println("also prints")
case a > 2:
    fmt.Println("no print here")
default:
    fmt.Println("this is the default")
}
```

## Arrays Maps and Slices

###  Arrays

https://golang.org/doc/effective_go#arrays

- Fairly standard, except that **the length of the array is part of the arrays type.** [link](https://golang.org/ref/spec#Array_types) This means if trying to copy the values of an array to a new variable (or existing one) **both** the stored type and the size must be the same!
- No mixed types, golang arrays are single type
- Usage is different from most languages, generally it is *idiomatic* to use **slices** instead of arrays in most use cases.
- As with most of golang, pass by value is used and if an array is assigned to a new variable, a **copy** is made (not a pointer)
  - Golang does allow for pass by reference with a pointer, by using the **address-of operator** `&` along with a function set up to use pointers (will cover later) [link](https://www.golang-book.com/books/intro/8)
- Initialising arrays

- Initialising and basic array probing/updating:

  ```go
  var x [4]int
  x = [4]int{3, 3, 3, 3}
  
  y := [4]int{3, 4, 4, 5}
  
  // ... will set the size to the number of values passed
  z := [...]int{3, 3, 3, 3, 3, 3, 3, 3}
  
  x[1] = 23
  fmt.Println(len(x), x[1])	// 4 23
  ```
  - The initialisation methods above are examples of **composite literals** [link](https://golang.org/ref/spec#Composite_literals). 

### Slices

- Slices are essentially references to arrays, however they have several methods associated with them which makes them (generally) easier to work with.

  - Slices can be extendable or **dynamically sized**
  - Slices have append functions associated

- A slice can either be initialised with reference to an array, or in one step (where a hidden reference array is created): *"A slice does not store data it just describes a section of an underlying array"*. 

  ```go
  var x [4]int
  x = [4]int{1, 2, 3, 4}
  // creating a slice from an array
  y := x[:]				// [1 2 3 4]
  y2 := x[1:3]			// [2 3]
  
  // initialising a slice directly
  z := []int{1, 2, 3, 4}	// [1 2 3 4]
  ```

  - slicing indexes are inclusive of the first value and exclusive of the second, so `[1:4]` would create a slice from element 1 up to and including element 3
  - Changing the elements of a slice also modifies the corresponding elements of the original array (and visa versa)

- A slice has **both** a **length** and a **capacity**. The capacity of a slice is the length of it's underlying array counting from the first element in a slice (capacity only extensible in the 'right' direction). Capacity can be probed with `cap()`:

  ```go
  x := [10]int{}
  y := x[0:5]
  z := x[5:10]
  
  fmt.Println(len(x), cap(x))		// 10 10
  fmt.Println(len(y), cap(y))		// 5 10
  fmt.Println(len(z), cap(z))		// 5 5
  ```

#### Appending to slices

https://golang.org/doc/effective_go#append

https://golang.org/ref/spec#Appending_and_copying_slices

- Using the `append()` method, which is variadic and can take any number of (compatible) parameters.

- When extending/appending to a slice with sufficient additional capacity, the underlying array will be updated:

  ```go
  x := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
  y := x[0:5]
  z := x[5:10]
  
  fmt.Println(len(y), cap(y))	// 5 10
  
  y = append(y, 1, 1, 1, 1, 1)
  fmt.Println(x, y, z)
  // [1 2 3 4 5 1 1 1 1 1] [1 2 3 4 5 1 1 1 1 1] [1 1 1 1 1]
  ```

- If the append exceeds the underlying array's capacity, a new underlying array will be assigned with sufficient capacity. This will not update the original array:

  ```go
  x := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
  y := x[0:5]
  z := x[5:10]
  
  fmt.Println(len(y), cap(y))	// 5 10
  
  y = append(y, 1, 1, 1, 1, 1, 1)
  fmt.Println(x, y, z)
  // [1 2 3 4 5 6 7 8 9 10] [1 2 3 4 5 1 1 1 1 1 1] [6 7 8 9 10]
  fmt.Println(len(y), cap(y)) // 11 20
  ```

- A slice can be appended to another slice.  Because of the way arguments are provided to append (variatic) a reference to a slice alone won't work. Instead `...` is used which unpacks a slice into elements, so to append slice `a` to slice `b`:

  ```go
  a := []int{1,2,3}
  b := []int{4,5,6}
  a = append(a, b...)	// [1 2 3 4 5 6]
  ```

- **Append can be used to delete**. When doing so the elements to the right of the deleted values will be shifted to the left in the underlying array, so just be wary of that~~

  ```go
  x := [6]int{1, 2, 3, 4, 5, 6}
  y := x[:]
  
  fmt.Println(y, len(y), cap(y))	// [1 2 3 4 5 6] 6 6
  // removing elements index 2 and 3
  y = append(x[0:2], x[4:]...)
  fmt.Println(y, len(y), cap(y))	// [1 2 5 6] 4 6
  fmt.Println(x)					// [1 2 5 6 5 6]
  ```

- Extending the underlying array (copying to a larger one when capacity is reached) is rather inefficient. A better idea is to create slices with an underlying array of sufficient size for what you want to do at the outset (reducing or eliminating any array copying).

  - Can either explicitly define the underlying array of the size wanted, then create a slice using it, or can use **make** to create a slice in this manner more directly (using a hidden array you don't need to initialise).

    - `make()` takes 3 arguments: type, length and capacity:

      ```go
      // int slice, size 50, capacity 100
      x := make([]int, 50, 100)
      ```

#### 2D slices

- Sometimes useful to have 2D slices. As slices are extensible, gotta be careful, probably better to use multidimensional arrays for certain things:

  ```go
  x := []int{1, 3, 5}
  y := []int{2, 4, 6}
  
  z := [][]int{x, y}
  
  fmt.Println(z) 	// [[1 3 5] [2 4 6]]
  z = [][]int{x, x, y, y}
  fmt.Println(z) // [[1 3 5] [1 3 5] [2 4 6] [2 4 6]]
  ```

#### Looping over arrays/slices (the range loop)

Can use standard loop and use `i` to get values at index, or can use `range` to perform a ranged loop:

```go
for i := 0; i < 10; i++ {
    fmt.Println(x[i])
}

// returns index and value (can discard by assigning to _ if wanted)
for i, v := range x {
    println(i, v)
}
```

- Actually instead of `i` and `v` (index and value) golang conventionally uses `i` and `s` with `s` being an assignment statement. so `for i, s := range x {...}`

The `range` loop can also be used to iterate over maps, channels and strings.

- When iterating over strings, the range loop iterates over UTF-8 code points instead of bytes. The index values returned will still be the index of the first byte of the code point:

  ```go
  for i, ch := range "日本語" {
      fmt.Printf("%v %c\n", i, ch)
  }
  // 0 日
  // 3 本
  // 6 語
  ```

### Maps

- Used for key-value pairs and allows for very fast lookup. Maps are unordered.
- Provide two types when initialising, the key type and the value type
- The length of a map is the number of value pairs (or map elements)

```go
m := map[string]int{
    "cake":  19,
    "cloud": 9,
    "life":  42,
}

fmt.Println(len(m), m["life"])	// 3 42
```

- If enter a key that does not exist in a map, the 'zero' value will be returned. Sometimes this is not desired, so the 'comma ok' method can be used to check for this:

  ```go
  v, ok := m["tea"]
  fmt.Println(v, ok)	// 0 false
  
  if _, ok := m["tea"]; !ok {
      fmt.Println("key does not exist") // prints
  }
  ```

- Adding to a map is simple:

  ```go
  m["tea"] = 23
  fmt.Println(m["tea"])	// 23
  ```

- Looping using `range` is also simple. The order printed may not be very predictable:

  ```go
  for k, v := range m {
      fmt.Print(k, " ", v, "\t")	// cake 19 cloud 9 life 42
  }
  ```

- Delete from a map using `delete()`:

  ```go
  delete(m, "tea")
  ```

  - The delete method does not throw an error if deleting a key which is not in the map. Delete also does not return anything. It is often a good idea to make sure a value exists before deleting using the comma ok:

    ```go
    if _, ok := m["tea"]; ok {
        delete(m, "tea")
    } else {
        fmt.Println("the tea is already gone!")
    }
    ```

## Struct(s)

- Allow for the storing values of multiple types. They store these values in a new type (which needs defining). There are similarities with objects or classes, however it's best to think of structs as their own thing.

- Example directly initialising a struct:

  ```go
  type person struct {
      first string
      last  string
      age   int
  }
  
  p := person{
      first: "Bob",
      last:  "Mccoy",
      age:   43,
  }
  fmt.Println(p) 			// {Bob Mccoy 43}
  fmt.Println(p.first)	// Bob
  
  // or can do (but less clear):
  p2 := person{"Tim", "Beans", 23}
  ```

- It is sometimes useful to create a constructor function to create structs of a given type. These are generally used when default initialisation values are not what's wanted. 

  ```go
  func newPerson(first string, last string, age int) person {
  	p := person{
  		first: first,
  		last:  last,
  		age:   age,
  	}
  	return p
  }
  ```

  - Can either return pointers to structs or the struct itself. If return the struct name the constructor `make____` instead of `new____`. Will cover pointers and advantages later~~

- You can use structs to create structs (sort of similar to inheritance):

  ```go
  type employee struct {
      person
      role string
  }
  
  em := employee{
      person: person{"Chris", "Janson", 53},
      role:   "internal communications supervisor",
  }
  
  em2 := employee{person{"Jill", "Simon", 23}, "twitter bot farmer"}
  fmt.Println(em)		//{{Chris Janson 53} internal communications supervisor}
  
  fmt.Println(em2)	// {{Jill Simon 23} twitter bot farmer}
  ```

  - When doing this the 'inner type' gets **promoted** to the 'outer type' - meaning that to access values still use dot notation e.g.:

    ```go
    fmt.Println(em.first)	// Chris
    fmt.Println(em2.role)	// twitter bot farmer
    ```

- **Anonymous structs** are helpful in cases where only need a certain data structure in one or two locations (and defining them globally or explicitly does not seem worthwhile).

  - They can be implemented by essentially writing the struct definition before the data to pass in:

    ```go
    p := struct {
        first string
        last  string
        age   int
    }{
        first: "Bob",
        last:  "Mccoy",
        age:   43,
    }
    ```

    
