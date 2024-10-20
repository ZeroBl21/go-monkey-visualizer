# Mini Tour of Monkey Language

This project is a **visualizer** for the Monkey language that includes a **REPL**, **lexer**, and **parser**.

## Requirements

### Dependencies

1. **Go**: You need to have Go installed (version 1.23 or higher).
   - [Install Go](https://go.dev/doc/install)

2. **Flex**: Required to generate the lexer for the Monkey language.
   - **Linux**: `flex` and `libfl` must be installed.
   - **Windows (cross-compilation)**: You will need MinGW and the appropriate tools to compile Flex from Linux.

### For Linux

1. **Install necessary dependencies**:
   ```bash
   sudo pacman -S flex mingw-w64-cross
   sudo pacman -S mingw-w64-gcc mingw-w64-binutils mingw-w64-make
   ```

2. **Build for Linux**:
   If you're only building for Linux, simply run:
   ```bash
   go build -o web ./cmd/web
   ```

3. **Run the server**:
   ```bash
   ./web
   ```

### For Windows (Cross-compiling from Linux)

1. **Install necessary dependencies**:
   If you haven't installed the MinGW tools for cross-compilation on Arch Linux, use:
   ```bash
   sudo pacman -S mingw-w64-gcc mingw-w64-binutils mingw-w64-make
   ```

2. **Build for Windows**:
   First, configure the environment for cross-compilation:
   ```bash
   export CC=x86_64-w64-mingw32-gcc
   export CXX=x86_64-w64-mingw32-g++
   export CGO_ENABLED=1
   export GOOS=windows
   export GOARCH=amd64
   ```

   Then, build the project for Windows:
   ```bash
   go build -o web.exe ./cmd/web
   ```

3. **Run the server on Windows**:
   Once the `.exe` file is built, you can run it on a Windows machine:
   ```bash
   ./web.exe
   ```

### For Compiling the Lexer

1. **Generate the lexer** using `flex`:
   - Make sure Flex is installed:
     ```bash
     sudo pacman -S flex
     ```

2. **Compile the lexer for Linux**:
   ```bash
   flex lexer.l
   gcc -o lexer lexer.c -lfl
   ```

   For **Windows**, you should use the MinGW tools as described above.

---

## Mini Tour of Monkey Language

### 1. Variables and Constants

```js
// x is now 10
let x = 10;
// y is now "Hello"
let y = "Hello";
```

### 2. Data Types

```js
// Integer
let num = 42;
// Boolean
let isValid = true; 
// String
let message = "Hello!";
```

### 3. Functions

```js
fn add(a, b) {
  return a + b;
}

// result is now 8
let result = add(5, 3);
```

### 4. Control Flow

#### If Statement

```js
let age = 20;

if (age >= 18) {
  puts("Adult");
} else {
  puts("Minor");
}
```

#### Returning Values from If Expressions

```js
let checkValue = fn(x) {
  if (x > 100) {
    true;  // The result of the if expression is `true` if x > 100
  } else {
    false; // The result of the if expression is `false` if x <= 100
  }
};

// Returns true
let result = checkValue(150);
// Returns false
let anotherResult = checkValue(50);
```

### 5. Arrays

```js
let numbers = [1, 2, 3, 4, 5];
// prints 1
print(numbers[0]);
```

### 6. Hash Maps

```js
let name = "Monkey";
let age = 1;
let inspirations = ["Scheme", "Lisp", "JavaScript", "Clojure"];
let book = {
  "title": "Writing A Compiler In Go",
  "author": "Thorsten Ball",
  "prequel": "Writing An Interpreter In Go"
};

let printBookName = fn(book) {
  let title = book["title"];
  let author = book["author"];
  puts(author + " - " + title);
};

printBookName(book);
```

### 7. Comments

```js
// This is a single-line comment

/*
  This is a
  multi-line comment
*/
```

### 8. Error Handling

```js
fn divide(a, b) {
  if (b == 0) {
    return "Error: Division by zero!";
  }
  return a / b;
}
```

### 9. Higher-Order Functions

#### Example 1: makeGreeter

```js
let makeGreeter = fn(greeting) { 
  fn(name) {
    greeting + " " + name + "!";
  }
};

let hello = makeGreeter("Hello");

hello("Zero");
```

#### Example 2: map function

```js
let map = fn(arr, f) {
  let iter = fn(arr, accumulated) {
    if (len(arr) == 0) {
      accumulated;
    } else {
      iter(rest(arr), push(accumulated, f(first(arr))));
    }
  };

  iter(arr, []);
};

let a = [1, 2, 3, 4, 5];
let double = fn(x) { x * 2 };

// Returns [2, 4, 6, 8, 10]
map(a, double);
```

### 10. Recursive Functions

```js
let fibonacci = fn(x) {
  if (x == 0) {
    0;
  } else {
    if (x == 1) {
      return 1;
    } else {
      fibonacci(x - 1) + fibonacci(x - 2);
    }
  }
};

let map = fn(arr, f) {
  let iter = fn(arr, accumulated) {
    if (len(arr) == 0) {
      accumulated;
    } else {
      iter(rest(arr), push(accumulated, f(first(arr))));
    }
  };

  iter(arr, []);
};

let numbers = [1, 1 + 1, 4 - 1, 2 * 2, 2 + 3, 12 / 2];
map(numbers, fibonacci);
```
