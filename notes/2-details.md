# Golang 2

## Golang functions and methods

- Golang (by default) is **pass by value**. The values of variables are passed to functions in arguments, not the variables themselves. (i.e. the operation of a function should not update the inputs, unless it is set to, returning it).

  - There are ways to override this behaviour (but more advanced)

- Declaring functions is fairly simple, and familiar, following the general format: `func (r receiver) identifier(parameter(s)) (return(s)) {}`, or can have no return type

  ```go
  func addInts(a int, b int) int {
  	return a + b
  }
  ```

- Golang allows for functions to return multiple values:

  ```go
  func addInts(a int, b int) (int, bool) {
  	return a + b, true
  }
  ```

- Can define functions with a variable number of parameters (variadic):

  ```go
  func sumInts(a ...int) int {
  	sum := 0
  	for _, v := range a {
  		sum = sum + v
  	}
  	return sum
  }
  ```

  - Variadic functions take **zero or more** parameters. If no arguments are passed to a function call, the zero value for the parameter type will be returned.

- (reminder) `...` can be used to pass slices of a type (which are their own type) to variadic functions which accept the type of the elements in the slice. Can be thought of un-wrapping a slice to it's elements: [ref](https://golang.org/ref/spec#Passing_arguments_to_..._parameters)

  ```go
  x := []int{1, 2, 3, 4}
  a := sumInts(x...)
  fmt.Println(a)		// 10
  ```

#### Anonymous Functions

- Inline functions are sometimes useful (especially later with concurrency using `go` keyword)

- Anonymous functions are the same as normal functions, just without the identifier. Only difference is that the call to the function (in the parenthesis) comes immediately after the function:

  ```go
  x := func(x int) int {
      return x + 4
  }(4)
  
  fmt.Println(x)		// 8
  ```

- An interesting feature is that an anonymous function (with no return) can be assigned to a variable. when that variable is called i.e. `var()` then the anonymous function runs. This is sometimes called a **function expression**

  - This works because functions are also types

  - This seems a roundabout way of naming a function, but it also allows for anonymous functions to be passed to functions (as arguments), or to be returned from functions:

    ```go
    func incrementor() func() int {
    	i := 0
    	return func() int {
    		i++
    		return i
    	}
    }
    
    func main() {
    	inc := incrementor()
    	fmt.Println(inc())	// 1
    	fmt.Println(inc())	// 2
    	fmt.Printf("%T\t%T\n", inc, inc())	
        // func() int      int
    }
    ```

    - `inc` is a function, `inc()` evaluates to an int. 
    - The state of `i` is saved in the instance of incrementor - a new instance would have it's own `i` value. This is an example of **closure** where a function references a variable outside it's 'body'.  This referenced variable is available only to that instance of the function. [link](https://tour.golang.org/moretypes/25)

##### Closure

Enclosing variables to limit their scope, like in the incrementor above. May see it in code examples and is useful to keep mind of and may be worth reading more on later.

```go
// another possible use

func durationMicroseconds() func() string {
	start := time.Now()
	return func() string {
		t := time.Since(start)
		elapsed := fmt.Sprintln("Microseconds:", t.Microseconds())
		return elapsed
	}
}
```

##### Callbacks

Callback is the term for passing a function as an argument to another function. They have a variety of uses e.g. in ML may want to pass a distance function to another function, so that any distance function can be used:

```go
func euclidian(a, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

func sumDistance2D(dist func(x, y float64) float64, points [][]float64) (sum float64, ok bool) {

	if len(points[0]) != len(points[1]) {
		sum -= 1
		return sum, ok
	}
	ok = true

	for i, _ := range points[0] {
		sum += dist(points[0][i], points[1][i])
	}

	return sum, ok
}

func main() {
	x := []float64{2, 4, 5, 1, 2, 32}
	y := []float64{6, 23, 1, 52, 5, 12}

	z := [][]float64{x, y}

	sumZ, _ := sumDistance2D(euclidian, z)
	fmt.Println(sumZ)	// 126.33960417797505
}
```

Note: Probably not the right way of handling errors~~

##### Recursion

As usual there are the standard memory trade-offs for recursion, however (as usual) it can simplify some tasks and writing something out recursively first can aid in finding a non-recursive solution e.g. Dynamic programming

Don't think there's anything different about golang recursion:

```go
func recursiveFactorial(n int) int{
    if n == 0 {
        return 1
    }
	return n * recursiveFactorial(n-1)
}
```

### Defer

https://golang.org/doc/effective_go#defer

https://golang.org/ref/spec#Defer_statements

A useful statement which *defers* a function call to run immediately before the function executing `defer` returns.

- It will also execute if the corresponding goroutine 'panics'

So in a function with multiple return paths, often need to make sure the same operation is performed before the return, no matter the path taken. Such as closing a file reader 

- Arguments to deferred functions are evaluated when the defer executes, not when the call executes (i.e. arguments will have the values where the defer occurs in code execution, rather than when the deferred method runs):

  ```go
  func main() {
  	for i := 0; i < 5; i++ {
  		if i == 0 {
  			defer fmt.Println("")
  		}
  		defer fmt.Printf("%d ", i)
  	}
  	fmt.Println("wow")
  }
  // output:
  // wow
  // 4 3 2 1 0
  ```

  - Can see the order `defers` are called is preserved (most recent first, so probably a stack) (run in reverse order they were deferred)

- Deferred functions execute **after** any result/return parameters are set, but **before** the function returns to it's caller

### Methods

https://golang.org/ref/spec#Method_declarations

- Essentially functions directly associated with a type (the method's **receiver**)
- After defining a method for a type, instances of that type can access that method using dot notation.
- By default still pass by value, so cannot (by default) update a type instance's values from a method (without a return)
  - There are ways to get round this with pointers and such

```go
type car struct {
	model string
	speed int
}

func (c car) goBrrr() {
	fmt.Println("Car go brrrrrr")
	c.speed = 100
	fmt.Println("Car speed is", c.speed)	// 100
}

func main() {
	c := car{
		model: "civic",
		speed: 0,
	}
	c.goBrrr()
	fmt.Println("Car speed is", c.speed)	// 0
}
```

To update the car type instance with a method, set the receiver to a pointer:

```go
func (c *car) goBrrr() {
	fmt.Println("Car go brrrrrr")
	c.speed = 100
}
```

### Interfaces

https://golang.org/ref/spec#Interface_types

https://golang.org/doc/effective_go#interfaces

Interfaces specify a [**method set**](https://golang.org/ref/spec#Method_sets). They specify the behaviour of an object (or custom type). 

Like Java, interfaces contain abstract methods where the inputs, outputs and names are defined, but the implementation is not.

- Interfaces are also **types** 

- If a type implements all the methods of an interface, then the instances of that type are **also** of the interface type. i.e. the variables will be of both types.

- An interface might be implemented by several types

- A type might implement several interfaces

- You can define functions to take in interface types as parameters. Because that interface might be implemented by several types, this allows functions to be written that can take several types as arguments. This is the primary way **polymorphism** is achieved in golang.

  - Often will may need to check the type in these methods (to use their different properties), which can be done with a switch statement

- Interface names often end in 'er' - like 'writer' or 'stringer'

- Interfaces can be incredibly useful to make types compatible with inbuilt interface methods in the standard library

- Knowledge of standard library interfaces is also extremely useful, allowing one to write more universal methods by setting interfaces as parameters instead of the types which implement them [example](https://www.alexedwards.net/blog/interfaces-explained)

- The types which implement interfaces are sometimes referred to as **concrete types**

- Basic interface:

  ```go
  type vehicle interface {
  	goBrrr()
  }
  
  type car struct {
  	model string
  	speed int
  }
  
  func (c car) goBrrr() {
  	fmt.Println("Car go brrrrrr")
  	c.speed = 100
  }
  
  type train struct {
  	model string
  	speed int
  }
  
  func (t train) goBrrr() {
  	fmt.Println("train go brrrrrr")
  	t.speed = 100
  }
  ```

  - Note how there's no keyword like `implements`. Instead concrete types will implicitly take on the interface type if it implements **all** of it's methods. (if it does not an error may be thrown, because won't be of that interface type)

- Example of a basic File interface where the abstract methods include the types of the inputs and outputs [ref](https://golang.org/ref/spec#Interface_types):

  ```go
  interface {
  	Read([]byte) (int, error)
  	Write([]byte) (int, error)
  	Close() error
  }
  ```

- The 'empty interface' is an interface with no methods. All types implement this empty interface and it often pops up in the docs when a function can take any type.

## Pointers

- Can use the `&` operator to probe the memory address of variables:

  ```go
  a := 42
  fmt.Println(&a) //0xc0000b8048
  ```

- The type of a memory address is a pointer type e.g. a 'pointer to an int'. Pointer types are distinct types:

  ```go
  fmt.Printf("%T %T\n", a, &a)	//int *int
  ```

  - Pointers are donated by `*` before the type: `*TYPE`

- Can get the value from an address using the `*` operator:

  ```go
  a := 42
  b := &a
  fmt.Println(*b)	// 42
  ```

- Because they are tied to the same memory address, updating the value either directly or with reference to the address will update them both:

  ```go
  a := 42
  b := &a
  a++
  fmt.Println(a, *b)	// 43 43
  *b++	
  fmt.Println(a, *b)	// 44 44
  ```

- Using these pointers and addresses, can generate functions which take addresses as arguments and directly modify the values stored in those addresses. This is a way of performing pass by reference - like operations if so desired (**but is still pass by value)**:

  ```go
  func incrementInt(a *int) {
      *a++
  } 
  
  func main() {
  	g := 42
  	incrementInt(&g)
  	fmt.Println(g)		//43
  }
  ```

  - Is still pass by value as the memory addresses are themselves stored in their own addresses, there is no sharing of addresses [link](https://dave.cheney.net/2017/04/29/there-is-no-pass-by-reference-in-go)
  - Getting the value from an address is called "**dereferencing**" or "**indirecting**"
  - Updating a value from an address is sometimes called "**mutating**"

### Method sets

https://golang.org/ref/spec#Method_sets

As mentioned method set is the set of methods associated with a type. In this set, there can be both methods that receive the type, or methods that receive pointers to the type.

- Methods which receive a type (`T`) can accept **both** a type object or a pointer to a type object
- Methods which receive a type pointer (`*T`) **only** accept pointers to a type object

The method set of the pointer type `*T` is therefore all the methods which receivers of type `T` or type `*T`.

Pointer receiving methods are useful for updating the values stored in a (usually struct) type (alongside some other operation).

