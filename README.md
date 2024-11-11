# Run Language Interpreter

**Run** is an interpreter for a custom-made programming language, developed in Go. It supports various programming constructs such as mathematical expressions, variable binding, functions, conditionals, return statements, higher-order functions, closures, and arrays.

## Supported Datatypes
- **Integers**
- **Booleans**
- **Strings**
- **Arrays**
- **Hashes**

## Getting Started

To run the interpreter, follow these steps:

```bash
git clone https://github.com/supremex04/Run.git
cd Run
go run main.go
```

## Built-in Functions

The **Run** language includes the following built-in functions:

- `len()`: Returns the length of a string or array.
- `first()`: Returns the first element of an array.
- `last()`: Returns the last element of an array.
- `rest()`: Displays all elements of an array except the first.
- `push()`: Appends an element to an array.
- `display()`: Prints a value to the screen.


**Sample code to run:**

```
let t = true;
let f = false;
let info = {"name": "Ram", "age": 999};
let x = len(info["name"]);
if (x > 2) {return t} else {return f};
let hw = ["Hello", " World in Run!"];
display( hw[0] + hw[1] );
```


## Screenshot

![Snapshot of the interpreter](./screenshot.png)


## Acknowledgment

This project was done by following Thorsten Ball's book on writing interpreters.



