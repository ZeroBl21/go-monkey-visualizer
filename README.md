# Mini Tour of Monkey Language

## 1\. Variables and Constants

```js
let x = 10; // x is now 10
let y = "Hello"; // y is now "Hello"
```
## 2\. Data Types

```js
let num = 42; // Integer
let isValid = true; // Boolean
let message = "Hello!"; // String
``

## 3\. Functions

```js
fn add(a, b) {
  return a + b;
}

let result = add(5, 3); // result is now 8
```

## 4\. Control Flow

### If Statement

```js
let age = 20;

if (age >= 18) {
  print("Adult");
} else {
  print("Minor");
}
```

### Returning Values from If Statements

```js

let result = 
  if (num % 2 == 0) {
      return "Even";
  } else {
      return "Odd";
  }

result; // result is now "Odd"
    
```

### For Loop

```js
for (i = 0; i < 5; i=i + 1) { print(i); } 
```

## 5\. Arrays

```js
let numbers = [1, 2, 3, 4, 5];
print(numbers[0]); // prints 1
```

## 6\. Hash Maps

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

## 7\. Comments

```js
// This is a single-line comment

/*
This is a
multi-line comment
*/
```

## 8\. Error Handling

```js
fn divide(a, b) {
  if (b == 0) {
	  return "Error: Division by zero!";
  }
  return a / b;
}
```
