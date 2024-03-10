
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
- Array's are copied by value not by reference.Same hold true when arrays are passed into function

### Slices [Sample Code](https://github.com/gauravvpa/golang-learning/blob/main/basics/go-tour/slices.go)
- Slices are dynamically sized data structure
- Slices can be create by following ways
	```
	//1. From Existing Array
	arr := [5]int{1, 4, 9, 16, 25}
	slice1 := arr[0:3]
	
	//2.Literal
	slice2 := []string{"C", "C++", "Java", "Go", "Python"}
	
	//3.From another slice
	slice3 := slice2[:2]
	
	// 4. make function make(Type, Length, Capacity)
	slice4 := make([]float32, 3)
	```
- When slices are created from arrays using the indices `arr[low: high]` slice includes the first element, but excludes the last one.
When any index is not present, then the default is zero for the low bound and the length of the slice for the high bound
  ```
  arr := [6]int{}
  slice1 := arr[1:4] // slice contains the elements from 1st index to 3rd index
  slice2 := arr[:2] // slice contains first two elements
  slice3 := arr[2:] // slice contains all elements except first two
  ```
- Slices are reference to the arrays. Changing the elements of a slice modifies the corresponding elements of its underlying array [Sample Code](https://github.com/gauravvpa/golang-learning/blob/main/basics/go-tour/slices.go#L54)
- `len(slice)` - The length of a slice is the number of elements it contains
- `cap(s)` - The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
- `append(s)` - Add new elements to slice

### Map [Sample Code](https://github.com/gauravvpa/golang-learning/blob/main/basics/go-tour/map.go)
- Map is a data structure where we store data as key value pairs
- Map can be create in following ways
```
// Nil map , you can not add key to nil maps
var nilMap map[string]string

// empty map (non-nil map)
map1 := make(map[string]string)

map2: = map[string]string{
	"Alex":"Alice"
}
```
- Like in other Programming Languages, we can use square notation to access or add a key in a map
```
sample:= make(map[string]string)

// add
sample["alex"] = "alice"

// update
sample["alex"] = "bob"

// delete a key
delete(sample, "random")

// check if a particular key exist
key, ok := sample["randomkey"]
// ok will be true if key exist
```
- [Map Iteration](https://github.com/gauravvpa/golang-learning/blob/main/basics/go-tour/map.go#L47)