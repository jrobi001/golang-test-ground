# GoLang

## Background/ Install stuff

- Workspace three main folders: bin, pkg and src.
  - Within src will be projects etc.

- Apparently method has changed, now instead of using env variables like gopath, use go modules [link](https://stackoverflow.com/questions/10838469/how-to-compile-go-program-consisting-of-multiple-files/61793820#61793820)

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
  - Not sure if any further conventions like package first letters being capitalised (or methods for that matter)
    - I prefer lowercase start, so will roll with that

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

