# Golang 2

## Golang functions and methods

- Golang is **pass by value**. The values of variables are passed to functions in arguments, not the variables themselves. (i.e. the values of variables are passed in arguments, not references to the memory space they take up. Unless set up to a function will not update the variables passed to it. Two variables cannot refer to the same object in memory).

  - There are ways to pass memory addresses to functions to update the values stored there within functions though (pointers)

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

#### Anonymous Functions (Function Literals)

- Inline functions are sometimes useful (especially later with concurrency using `go` keyword)

- Anonymous functions are the same as normal functions, just without the identifier. Only difference is that the call to the function (in the parenthesis) comes immediately after the function definition:

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
  
  - There are two way of  mutating struct type properties/values in functions. Either can use `(*T).f`, or if withing a function which takes `*T` as input, can reference the same thing (implicitly) just using `T.f` (dereferencing the variable is implicit): [ref](https://golang.org/ref/spec#Selectors)
  
    ```go
    func changeName(p *person, name string) {
    	p.name = name		// equivalent
    	(*p).name = name	// equivalent
    }
    ```

### Method sets

https://golang.org/ref/spec#Method_sets

https://golang.org/doc/effective_go#methods

As mentioned method set is the set of methods associated with a type. In this set, there can be both methods that receive the type, or methods that receive pointers to the type.

- Methods which receive a type (`T`) can accept **both** a type object or a pointer to a type object
- Methods which receive a type pointer (`*T`) **only** accept pointers to a type object

The method set of the pointer type `*T` is therefore all the methods which receivers of type `T` or type `*T`.

Methods which receive a pointer `*T` can still be accessed by a instance of a type `T` using dot notation. In this case the compiler implicitly converts `T` to `&T` when making these calls (the method still only accepts `*T`).

Pointer receiving methods are useful for updating the values stored in a type (alongside some other operation) without the need to use returns to assign an updated value to an instance.

Because the method sets of `T` and `*T` are not the same (`*T` larger). It is fairly common to have the pointer of a type implement an interface, but not the type itself! e.g. `*T` may implement `io.Writer` but `T` does not

```go
type square struct {
	length float64
}

type shape interface {
	area() float64
}

// *square implements shape
func (c *square) area() float64 {
	return c.length * c.length
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	c := square{5}
	info(&c)		// works
    info(c)			// fails: square does not implement shape
}
```

Useful blog post:

https://gronskiy.com/posts/2020-04-golang-pointer-vs-value-methods/

## Commonly used standard lib tools

### JSON

https://pkg.go.dev/encoding/json

Implements encoding and decoding of JSON. The mapping between JSON and Go values is performed by Marshal and Unmarshal functions.

#### Marshal

Marshal returns the JSON encoding of Go types/objects.

It has default behaviour for encoding inbuilt types, including structs (performed by `TextMarshalJson()`). It is also possible to define custom JSON encoding behaviour by implementing the `Marshaler` interface, by implementing a `MarshalJSON` method for a type.

Basic example:

```go
type person struct {
	Name string
	Age  int   
}

func main() {
	p1 := person{
		Name: "Bob Mccoy",
		Age:  32,
	}
	p2 := person{
		Name: "Jill Jefferson",
		Age:  65,
	}
    
	persons := []person{p1, p2}
	bs, err := json.Marshal(persons)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
    // [{"Name":"Bob Mccoy","Age":32},{"Name":"Jill Jefferson","Age":65}]
}
```

- When using undefined Marsheler, only capitalised fields are processed (fields need to be exportable to be accessed by the json package):

  ```go
  type person struct {
  	Name string 	// will be processed
  	age  int    	// won't be processed
  }
  ```

- Generally best if both struct and fields are set to export / are capitalised.

- The keys used in JSON can be modified, along with whether fields are to be included by providing a **format string** under the json key in a structs field tag: 

  ```go
  type person struct {
      // renaming fields
  	Name     string  `json:"PersonName"`
  	Age      int     `json:"PersonAge"`
      // ommit if empty
  	Height   float64 `json:"PersonHeight, omitempty"`
      // don't include field
  	StarSign string  `json:"-"`
  }
  ```

  - From what I've seen providing a format string seems best practice

#### Unmarshal

https://pkg.go.dev/encoding/json#Unmarshal

Unmarshal parses JSON and stores the value(s) in an object: `func Unmarshal(data []byte, v interface{}) error`

Depending on the JSON data, it can be unmarshaled into a range of Types. e.g a JSON array -> a go array, a JSON object -> a go map etc.

Probably more common to Unmarshal into a struct. To do so a compatible struct needs to be created with matching field names (or JSON aliases provided by a format string):

- There is a useful website which can aid in writing structs to store JSON data https://mholt.github.io/json-to-go/
  - May still need to do some tweaking though
- (generally) provide Unmarshal with a pointer (`&T`) to a compatible struct type (or slice)

- Simple example unmarshal:

  ```go
  type person struct {
  	Name   string  `json:"PersonName"`
  	Age    int     `json:"PersonAge"`
  	Height float64 `json:"PersonHeight"`
  }
  
  func main() {
  
  	bs := []byte(`[{"PersonName":"Bob Mccoy","PersonAge":32,"PersonHeight":1.78},{"PersonName":"Jill Jefferson","PersonAge":65,"PersonHeight":1.58}]`)
  	
  	var people []person
  	
  	err := json.Unmarshal(bs, &people)
  
  	if err != nil {
  		fmt.Println(err)
  	}
  
  	fmt.Println(people[0].Name)
  }
  ```

#### Decoder/Encoder

For when receiving or passing on the data, generally either from/to a file, or over the internet. Like a stream~

Uses `io.Reader` and `io.Writer` Interfaces respectively

### IO, readers and writers

https://pkg.go.dev/io

https://pkg.go.dev/io/ioutil

Will come across them all the time. There are `Reader` and `Writer` interfaces in the `io` package and any Type implementing the associated Write() or Read() method respecitvely will be an `io.Writer` or `io.Reader`

Common example are files, where the file type [in os](https://pkg.go.dev/os#File) implements both Write and Read. (`os.stdout` which `fmt.Println` uses is also a file with associated `Read` and `Write` methods)

### Sort

https://pkg.go.dev/sort

- There are convenient inbuilt sort methods for slices of various types, as well as useful methods for searching and `IsSorted()` etc.

- There is a special interface - just called `Interface`. It is only really special because of it's name, but it uses the name because of it's importance, as a set of methods that are a good idea to implement for many defined types.

  - The `Interface` interface defines three methods for the slice of a type, which when defined allows the inbuilt `Sort()` methods to work on them. Because you define them, you can decide how to sort.

  - The three required methods are `Len`, `Less` and `Swap` [pkg ref](https://pkg.go.dev/sort#Interface):

    ```go
    
    type Interface interface {
    	// Len is the number of elements in the collection.
    	Len() int
    	// Less reports whether the element with index i
    	// must sort before the element with index j.
    	Less(i, j int) bool
    
    	// Swap swaps the elements with indexes i and j.
    	Swap(i, j int)
    }
    ```

  - Example [link](https://pkg.go.dev/sort#Interface)

    ```go
    type Person struct {
    	Name string
    	Age  int
    }
    
    func (p Person) String() string {
    	return fmt.Sprintf("%s: %d", p.Name, p.Age)
    }
    
    // ByAge implements sort.Interface for []Person based on
    // the Age field.
    type ByAge []Person
    
    func (a ByAge) Len() int           { return len(a) }
    func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
    func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
    
    func main() {
    	people := []Person{
    		{"Bob", 31},
    		{"John", 42},
    		{"Michael", 17},
    		{"Jenny", 26},
    	}
    
    	fmt.Println(people)
    	sort.Sort(ByAge(people))
    	fmt.Println(people)
    }
    ```

- A slice can also be sorted without implementing `Interface`. Instead call `sort.Slice()` along with the implementation of a `Less` function, which can either be defined or anonymous. 

  - `func Slice(x interface{}, less func(i, j int) bool)`

  - e.g. with an anonymous `less` function:

    ```go
    sort.Slice(people, func(i, j int) bool {
        return people[i].Age > people[j].Age
    })
    ```

### bcrypt and /x/ packages

`bcrypt` is not in the standard library (yet) but an official implementation  exists (but not finalised) in the `golang.org/x/crypto` package.

`/x/` packages are called sub-repositories. golang gives this definition:

> These packages are part of the Go Project but outside the main Go tree. They are developed under looser compatibility requirements than the Go core.

Generally go is built in a way so that all elements of the standard library will be forward compatible with all minor releases, meaning implementations using just these tools should not break between releases. The `/x/` packages are a collection of useful and experimental tools which do not necessarily make this guarantee. [ref](https://rodaine.com/2017/05/x-files-intro/)

However most are fairly safe to use. To import them, use `go get` in a terminal.

## Concurrency

https://golang.org/doc/effective_go#concurrency

A large motivation behind the creation of Go was to create a language with native support for multi-core CPUs.

Concurrency is a design pattern of writing code which is capable of being run in parallel. Whether it does is may be down to hardware. Parallelism is when the code runs at the same time (often facilitated by concurrent design).

[concurrency is not parallelism](https://www.youtube.com/watch?v=oV9rvDllKEg)

Creating a concurrent thread in go is very simple a `go` statement e.g. `x := go statement` or `go function()`.

Simply creating a concurrent thread is generally not enough. Code execution will continue on the main thread as the other thread initialises and starts to run. This means that the execution of the main thread may finish (exit the main function) before the concurrent thread even starts. (or that a concurrent statement may not have been evaluated in time for where it is needed in the main thread). When the main function reaches the end/exits, it does not wait for concurrent code to finish executing, the program just exits (by default).

e.g. very simple example, only loop1 will execute and print, despite being called after the concurrently run loop2:

```go
func loop1() {
	for i := 0; i < 10; i++ {
		fmt.Println("loop1", i)
	}
}

func loop2() {
	for i := 0; i < 10; i++ {
		fmt.Println("loop2", i)
	}
}

func main() {
    go loop2()		// too slow (main loop finishes before)
	loop1()			// prints to stdout
}
```

This behaviour is desirable (in certain circumstances), however often want to control the concurrent code to behave differently, this is done using 'synchronisation'.

Standard library Synchronisation primitives/ tools can be found in the `sync` package: https://pkg.go.dev/sync

#### Share by communicating

> Concurrent programming in many environments is made difficult by the subtleties required to implement correct access to shared variables. Go encourages a different approach in which shared values are passed around on channels and, in fact, never actively shared by separate threads of execution. Only one goroutine has access to the value at any given time. Data races cannot occur, by design. To encourage this way of thinking we have reduced it to a slogan:
>
> > Do not communicate by sharing memory; instead, share memory by communicating.
>
> https://golang.org/doc/effective_go#concurrency

A data race is when two threads access the same memory location at the same time. Issues can occur if both threads attempt to update the same location at the same time (possibly causing corruption), or if two methods that update a value take the initial value at the same time and write the updates one after each other (possibly causing one update to be lost).

n.b.  Think banking and updating balance values and memory locks

Can artificially create (more dramatic) data races easily by yielding CPU cores/threads when running concurrent code (yielding allows core to run other go routines). Using `runtime.Gosched()`. (something similar happens if use `time.Sleep()`)

A data race can be detected when running a file by running it with:
`go run -race file.go`

Data races are not always obvious, a race may occur, but not affect the results over 99% of the time. So it is important to be very conscious when writing concurrent code to avoid them.

### Wait groups

https://pkg.go.dev/sync#WaitGroup

One way of performing synchronisation. You assign the type `WaitGroup` to a named variable which acts as a counter, primarily for use with goroutines.

A `WaitGroup` has three methods:

1. `Add(delta int)` which increments the counter by delta (including -ve).
2. `Done()` which decrements the counter by one
3. `Wait()` which **blocks** code execution (below) until the WaitGroup counter is zero

`Add` should be called before a concurrent function or statement is called and `Done` should be called at the end of the concurrent function or statement (generally by using `defer` at the top of the block for readability).

The need to include `WaitGroup.Done()` inside a function, makes it's use directly with named functions less than ideal (as it requires a global `WaitGroup` and care to be taken to ensure the counter decrements to zero at the right time). 

Generally `WaitGroup` is used with anonymous concurrent functions (which may themselves call named functions).

**Note**: keep forgetting, but golang docs call **anonymous functions** **function literals**

example:

```go
func main() {
	var wg sync.WaitGroup
	
	wg.Add(1)
	go func() {
		defer wg.Done()
		loop2()					// prints third
	}()

	loop1()						// prints first
	fmt.Println("before wait")	// prints second
	wg.Wait()
	fmt.Println("after wait")	// prints last

}

func loop1() {
	for i := 0; i < 10; i++ {
		fmt.Println("loop1", i)
	}
}

func loop2() {
	for i := 0; i < 10; i++ {
		fmt.Println("loop2", i)
	}
}

// The print order is not important, but gives an indication. If loop1 was large, loop2 would be printing alongside part way through it's execution (on my machine, generally if loop1 > 300)
```

### Mutex

https://pkg.go.dev/sync#Mutex

https://tour.golang.org/concurrency/9

Mutual exclusion locks - they grant exclusive access to a shared resource, to  a single thread. If another thread wants the resource, it must wait for it to be released from the previous thread first.

Mutex's are one way of preventing data races. a `Mutex`has two methods associated: `Lock()` and `Unlock()`.  Code between lock and unlock will be executed in mutual exclusion. It is often idiomatic to `defer` the `Unlock()`.

A simple example:

```go
var count int
var wg sync.WaitGroup
var mu sync.Mutex

for i := 0; i < 100; i++ {
    wg.Add(1)
    go func() {
        mu.Lock()
        defer mu.Unlock()
        defer wg.Done()
        count++
    }()
}
wg.Wait()
fmt.Println(count)
```

- There is also `RWMutex` which allows for locks on reading and writing independently 

### Channels

https://golang.org/ref/spec#Channel_types

https://golang.org/doc/effective_go#channels

Channels provide a mechanism for concurrently executing functions to communicate with one another, by sending and receiving values of a specified type.

Channels are important, in golang, care was taken to avoid memory sharing and pass by reference. In concurrency this theme is continued with the mantra:

> *Do not communicate by sharing memory; instead, share memory by communicating.*
>
> https://go-proverbs.github.io/

Channels allow this to happen by passing values between routines/threads (instead of memory or variable references).

Channels can be unidirectional or bidirectional i.e. can be send only and receive only channels, or channels which perform both.

Channels should be created using `make()` and a capacity. In the case of a channel, the capacity is the size of the **buffer**. Unbuffered channels (capacity 0, default) only communicate if both the sending and receiving channel are ready.

Examples of making unidirectional and bidirectional int channels:

```go
a := make(chan<- int)		// Send only unbuffered
b := make(<-chan  int)		// Receive only unbuffered
c := make(chan int)			// Bidirectional unbuffered
d := make(chan int, 100)	// Buffer of 100
```

##### Sending and receiving

Channels pass values using the receive operator:  `<-` .

A **send** statement: `SendStmt = Channel "<-" Expression .`

```go
ch <- 3 		// send value 3 to channel ch
```

A **receive** expression:

```go
v1 := <-ch		// receive calue from ch
v2, ok := <-ch
```

A combined send and receive expression (passing on):

```go
ch2 <- <- ch1
```

#### Channels block

Without buffers, channels **block**. They block a sender until a receiver is ready and they block receiver until a sender is ready. A go routine will 'sleep' until two channels are ready to send and receive simultaneously i.e. the go routine will sleep until two are synchronised.

This blocking is one way of handling synchronisation. However is not desired for some tasks and there are ways around it if needed.

Simple example:

```go
c := make(chan int)

go func() {
    c <- 123		// send 123
}()

fmt.Println(<-c)	// receive 123
```

- Receive expressions block until a value is available
- Communication of a send statement is blocked until a receiver is available (evaluation is completed before communication)

#### Buffered channels

Buffered channels can store a defined number of values without blocking.

They have specific use cases, such as in divide and conquer algorithms, where may know the max number of concurrent operations and returns and want to perform these operations concurrently, but may not want to receive all the values untill they are all done.

Not enforcing synchronisation for each send/receive can also increase efficiency of concurrent operations if handled correctly.

There are other uses too, will return to this later, but for now studying basic synchronisation will stick to unbuffered channels~~

#### Using unidirectional channels

Reminder:

```go
c1 := make(<-chan int)	// Recceive only
c1 := make(chan<- int)	// Send only
```

A bidirectional channel can be converted to or used as a unidirectional channel, however the reverse is not true.

Unidirectional channels are useful when defining functions or methods. Can limit the channel passed into a function to act as a send only or receive only channel.

Simple example:

```go
func main() {
	c := make(chan int)

	go send(23, c)
	a := receive(c)
	fmt.Println(a)
}

func send(a int, c chan<- int) {
	c <- a
}

func receive(c <-chan int) int {
	return <-c
}
```

#### Using channels (general)

- You can **range** over a channel. Unlike most ranges where there are a fixed number of items in the object being ranged (when starting the loop), with ranges the program does not know how many items it will process. When ranging over a channel the program essentially waits for all the values to be presented. 
  - In order to tell the program when to end waiting you need to **close** the channel.
  - A **range** on a channel will block (until a close is encountered)

```go
c := make(chan int)

go func() {
    for i := 0; i < 1000; i++ {
        c <- i
    }
    close(c)
}()

for v := range c {
    fmt.Println(v)
}
```

- The **select statement** is associated with channels and sending/receiving. The select statement is very similar to a **switch** with multiple **cases**, however all cases refer to communication options. 

  ```go
  even := make(chan int)
  odd := make(chan int)
  quit := make(chan int)
  
  go func() {
      for i := 0; i < 100; i += 2 {
          even <- i
      }
      quit <- 0
  }()
  
  go func() {
      for i := 1; i < 100; i += 2 {
          odd <- i
      }
      quit <- 0
  }()
  
  for {
      select {
      case <-even:
          fmt.Println("even")
      case <-odd:
          fmt.Println("odd")
      case <-quit:
          return
      }
  }
  
  fmt.Println("Done")
  ```

  - Careful, a `break` in a select breaks out of the select, not any containing loop.
  - The distinction from a switch is likely that the select statement will block.

- When receiving from a stream, can assign the value to a single variable, or use 'comma ok' to also return a bool which idicates whether the communication succeeded. (`true` indicates the value received was delivered by a successful send operation. Zero values may be received even without a successful send operation, in this case `ok` will be `false`):

  ```go
  even := make(chan int)
  odd := make(chan int)
  
  go func() {
      for i := 0; i < 100; i += 2 {
          even <- i
      }
      close(even)
  }()
  
  go func() {
      for i := 1; i < 100; i += 2 {
          odd <- i
      }
      close(odd)
  }()
  
  evencount := 0
  oddcount := 0
  done := false
  for done == false {
  
      select {
      case _, ok := <-even:
          if ok != false {
              fmt.Println("even")
              evencount++
          } else {
              fmt.Println("Even finsihed first", evencount)
              done = true
          }
      case _, ok := <-odd:
          if ok != false {
              fmt.Println("odd")
              oddcount++
          } else {
              fmt.Println("Odd finsihed first", oddcount)
              done = true
          }
      }
  }
  ```

- Channels are often passed as inputs and outputs of functions. When returning a channel, the channel will not necessarily be 'complete', values may still be being added to the channel concurrently *after* it is returned. The channel may be returned and fed into another function, *before* any values are put into the channel. (returns do not block!) e.g.

  ```go
  func gen() <-chan int {
  	c := make(chan int)
  
  	go func() {
  		for i := 0; i < 10; i++ {
  			c <- i
  		}
  		close(c)
  	}()
  
  	return c
  }
  
  // c is returned (likely) before any values sent through the channel
  ```

#### Concurrent patterns - fan out, fan in

https://medium.com/@thejasbabu/concurrency-patterns-golang-5c5e1bcd0833

https://blog.golang.org/pipelines

https://talks.golang.org/2012/concurrency.slide#25

Fanning out and fanning in are common design patterns in concurrent programming (in golang). It is often used to execute multiple functions concurrently, then aggregate the results (think divide and conquer).

In other languages similar processes are sometimes referred to as multiplexing and demultiplexing.

#### Context

https://blog.golang.org/context

Often (especially in web) a request my spin off into several go routines (or processes). There is a need if a request is cancelled, to also cancel those routines, to save resources. This can also occur with processes and sub-processes, you want the associated sub-processes to end if the main process ends.

The 'context' package which helps send deadlines, cancellations and 'request scoped values' across API boundaries.

## Error Handling

