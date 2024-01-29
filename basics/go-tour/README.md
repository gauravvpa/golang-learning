
### Packages
- Every Go program consists of packages
- There is a special package called `main` in which program starts running
- A package can import other packages  with `import` statement
- Whatever you want to export from a package, it should starts with capital letter

### Variables [Sample Code](https://github.com/gauravvpa/golang-learning/blob/main/basics/go-tour/variable.go)
- Go is statically typed lang
- Variables are declared like `var name type`
- Variables declaration is implicit when initialized `var name = value`
- the `:=` short assignment statement can be used in place of a var declaration with implicit type.
- constants are declared using `const` like `const c = 1`. 
- Shorthand (`:=`) can not be used for declaring constants because compiler can not decide that it is actually a varaiable (the value can be changed later) or constant.

### Functions [Sample Code](https://github.com/gauravvpa/golang-learning/blob/main/basics/go-tour/function.go)
- While defining a function you need to define the data type of a parameters and return values too.
- Function can return multiple values like in Javascript


### Loops
- Only for loop exists compared to other Programming Languages
- Code examples-
   ```
   // i variable scope is only limited to for loop
   for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Print(i, "\n")
	}
   ```
   ```
   // while loop 
	value := 10

	for value < 15 {
		value += 2
	}
   ```
   ```
   // infinite loop
	for {

	}
   ```

### What's new with if else
- We can add a short statement for execution before the if condition   
- Variables declared by the statement are only in scope if and else blocks.
   ```
   if c := 25; c < 25 {
		fmt.Print("If block", c)
	} else {
		fmt.Print("Else block", c)
	}
	// fmt.Print(c) undefined: c
   ```

### What's new with switch
- `break` statement is not required because it only runs the case which matches, not all the cases that follow.
- `switch` without condition can be written. It is evaluated as `switch true`, in this case, the first condition which evaluates to `true` will execute its corresponding block
   ```
   switch {
	case condition1:
    	// code to be executed if condition1 is true
	case condition2:
    	// code to be executed if condition2 is true
	default:
    	// code to be executed if none of the conditions match
	}
   ```

### defer 
- TODO


### Structs
- Collection of fields.
- Go do not have classes. So whenever we want to represent the data which have multiple properties (for ex- A Student , which can have Name, Age, Class, Role Number etc), we can use struct
	```
	type Student struct {
		Name string
		Age int
	}
	student1 := Student{"Alice Anderson", 14}
	fmt.Println(student1.Name) // Alice
	```

### Arrays 
- We can create fixed size arrays in Go
- Arrays Constructors
	```
	var arr1 [5]int // default is 0
	var arr2 [2]int = [2]int{1, 2}
	arr3 := [6]int{2, 4, 8, 16, 32, 64}
	arr4 := [...]string{"SDE"} // size is decide by the elements present, size is 1 here

	// multiDimensional Array
	arr5 := [2][2]int{{1, 2}, {3, 4}} 
	```
- Array's length can be obtained from `len` function