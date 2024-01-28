
### Packages
- Every Go program consists of packages
- There is a special package called `main` in which program starts running
- A package can import other packages  with `import` statement
- Whatever you want to export from a package, it should starts with capital letter

### Variables
- Go is statically typed lang
- Variables are declared like `var name type`
- Variables declaration is implicit when initialized `var name = value`
- the `:=` short assignment statement can be used in place of a var declaration with implicit type.
- constants are declared using `const` like `const c = 1`. 
- Shorthand (`:=`) can not be used for declaring constants because compiler can not decide that it is actually a varaiable (the value can be changed later) or constant.

### Functions
- While defining a function you need to define the data type of a parameters and return values too.
- Function can return multiple values like in Javascript

