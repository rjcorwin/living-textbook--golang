# Structs and Interfaces in Go

## Overview

Go is a statically typed language that provides powerful ways to structure and organize data using structs and interfaces. Understanding these concepts is essential for writing idiomatic and maintainable Go code.

---

## Structs

A struct is a composite data type that groups together variables (fields) under a single type. Structs are used to model real-world entities and data structures.

### Example: Defining and Using a Struct
```go
package main
import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    rj := Person{Name: "RJ", Age: 30}
    fmt.Println(rj)
    fmt.Printf("Name: %s\n", rj.Name)
    fmt.Printf("Age: %d\n", rj.Age)
}
```

This example prints the struct as a whole, and then prints each field name and value explicitly using Go's dot notation for struct fields.

### Key Points
- Struct fields can be accessed using dot notation (e.g., `rj.Name`).
- Structs can have methods associated with them.

### Example: Using Methods as Setters in Go (Compared to JS/TS Properties)

In Go, you can use methods to control how struct fields are set, similar to how you might use setters in JavaScript/TypeScript properties. Here's an example:

```go
package main
import "fmt"

type Person struct {
    name string // unexported field (private)
    Age  int
}

// Setter method for name
func (p *Person) SetName(newName string) {
    if (newName != "") {
        p.name = newName
    }
}

// Getter method for name
func (p Person) Name() string {
    return p.name
}

func main() {
    rj := Person{Age: 30}
    rj.SetName("RJ")
    fmt.Printf("Name: %s\n", rj.Name())
    fmt.Printf("Age: %d\n", rj.Age)
}
```

**Go Naming Convention Note:**
- Setter methods in Go are typically named with a `Set` prefix (e.g., `SetName`).
- Getter methods are usually just the field name, capitalized (e.g., `Name()`), not `GetName()`.
- This is different from languages like Java or C#, where both getters and setters are prefixed (e.g., `getName`/`setName`).
- This convention keeps Go code concise and idiomatic.

**JS/TS Comparison:**
```typescript
class Person {
  private _name: string = "";
  age: number;
  set name(newName: string) {
    if (newName) this._name = newName;
  }
  get name() {
    return this._name;
  }
  constructor(age: number) {
    this.age = age;
  }
}
const rj = new Person(30);
rj.name = "RJ";
console.log(rj.name); // "RJ"
console.log(rj.age);  // 30
```

**Key Points:**
- In Go, you use methods to control access to fields (no built-in property syntax).
- In JS/TS, you use `get`/`set` for properties.
- Go fields can be made private (unexported) by starting with a lowercase letter.
- This pattern is useful for validation or encapsulation in Go.

---

## Interfaces

An interface is a type that specifies a set of method signatures. Any type that implements those methods satisfies the interface.

### Example: Defining and Using an Interface
```go

// Note:
// The parameter g in sayHello has the type Greeter (the interface), not Person (the struct).
// Because Person implements the Greet() method, it satisfies the Greeter interface.
// This means you can pass a Person to sayHello, even though the function expects a Greeter.
// In Go, we say that Person 'implements' the Greeter interface by having the required method(s).
// This is a key feature of Go's interfaces: any type that has the necessary methods automatically satisfies the interface (implicit implementation).
// In Go, we usually say 'type' for the parameter, so: 'the g parameter's type is Greeter.'

type Greeter interface {
    Greet() string
}

type Person struct {
    Name string
}

func (p Person) Greet() string {
    return "Hello, my name is " + p.Name
}

func sayHello(g Greeter) {
    fmt.Println(g.Greet())
}

func main() {
    rj := Person{Name: "RJ"}
    sayHello(rj)
}
```

### Key Points
- Interfaces are satisfied implicitly in Go (no `implements` keyword).
- Interfaces enable polymorphism and decoupling.

---

## Aha! Moment
> "Aha! In Go, the variables inside a struct are called fields, not properties. In languages like JavaScript, C#, or Python, 'properties' often refer to variables that can have custom getter and setter logic, allowing you to control how values are accessed or modified. In Go, struct fields are always accessed directly—there are no built-in getters or setters for fields. If you want to control access or add logic, you define explicit methods, but the fields themselves remain simple and direct. This makes Go's approach straightforward and explicit."

> "I realized that Go's interfaces are satisfied implicitly, which makes the language very flexible and encourages composition over inheritance! In Go, you don't need to declare that a type implements an interface—if it has the required methods, it just works. This is different from JavaScript/TypeScript, where you must explicitly declare that a class implements an interface. This implicit satisfaction means you can write more decoupled and flexible code, and types can satisfy multiple interfaces without extra boilerplate. It also encourages designing small, focused interfaces (sometimes called 'interface segregation'), making your code easier to test and extend. This is a powerful feature that supports Go's philosophy of simplicity and composability. __It's like, if it rhymes, it works.__"

---

[Back to Curriculum](curriculum.md)
