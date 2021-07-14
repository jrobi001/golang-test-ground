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

    

