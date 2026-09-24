# Go Notes

## Commands, Modules, and Packages

```sh
go run .                         # Compile and run the main package in the current directory.
go build .                       # Build the current package (an executable for package main).
go test ./...                    # Run tests in every package in the module.
go mod init example.com/my-app   # Create a module and its go.mod file.
go get example.com/library@latest # Add or update a dependency.
go mod tidy                      # Add missing and remove unused module dependencies.
```

- A **module** is a versioned collection of packages, defined by `go.mod`.
- A **package** is a directory of Go source files that use the same `package` declaration.
- `package main` identifies an executable package. Its entry point is `func main()`.
- Identifiers beginning with an uppercase letter are exported from their package; lowercase identifiers are package-private.

## Types and Zero Values

Common types:

- `int`: signed integer whose size is implementation-specific (usually 32 or 64 bits).
- `float64`: 64-bit floating-point number.
- `string`: immutable sequence of bytes, commonly UTF-8 text.
- `bool`: `true` or `false`.
- `byte`: alias for `uint8`.
- `rune`: alias for `int32`, usually representing one Unicode code point.

Fixed-size integer types:

| Type     | Range                                 |
| -------- | ------------------------------------- |
| `int8`   | -128 to 127                           |
| `uint8`  | 0 to 255                              |
| `int32`  | -2<sup>31</sup> to 2<sup>31</sup> - 1 |
| `uint32` | 0 to 2<sup>32</sup> - 1               |
| `int64`  | -2<sup>63</sup> to 2<sup>63</sup> - 1 |

Every declared variable has a useful **zero value**:

| Type                                                    | Zero value              |
| ------------------------------------------------------- | ----------------------- |
| Numeric types                                           | `0`                     |
| `string`                                                | `""`                    |
| `bool`                                                  | `false`                 |
| Pointers, maps, slices, functions, channels, interfaces | `nil`                   |
| Structs                                                 | Each field's zero value |

## Functions and Errors

Functions can return multiple values. By convention, a fallible operation returns its result followed by an `error`.

```go
func PrintBalance(filename string) error {
	balance, err := ReadBalanceFromFile(filename)
	if err != nil {
		return err
	}
	fmt.Println(balance)
	return nil
}
```

Go does not force error handling, but ignoring a non-nil error usually produces incorrect behavior. Handle it, return it, or explicitly document why it is safe to ignore.

## Defer, Panic, and Recover

`defer` schedules a call to run when the surrounding function returns. It is commonly used for cleanup after a successful acquisition.

```go
func ReadNotes() error {
	file, err := os.Open("notes.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	// Read from file.
	return nil
}
```

Rules for deferred calls:

1. Arguments are evaluated when `defer` is executed.
2. Calls run in last-in, first-out order.
3. A deferred function can inspect or change named return values.

```go
func incrementResult() (result int) {
	defer func() { result++ }()
	return 1
}
```

Use `panic` for unrecoverable programmer errors or broken invariants, not ordinary expected failures. `recover` only works when called directly by a deferred function in the same goroutine; use it sparingly, usually at a program boundary.

## Values and Pointers

Go always passes arguments **by value**. Passing a pointer copies the pointer value, which still points at the original object.

```go
func Birthday(age *int) {
	*age++ // Equivalent to (*age)++.
}

age := 32
Birthday(&age)
fmt.Println(age) // 33
```

Use a pointer when a function must modify the caller's value, when `nil` has useful meaning, or when copying a large value is undesirable. Small values such as `int` are normally passed by value. Go has pointers, but it does not have C++-style reference parameters.

## Structs and Methods

Structs group related fields. Use a value receiver when the method does not need to modify the receiver and copying it is appropriate. Use a pointer receiver when it modifies the receiver, the struct is large, or the type should not be copied.

```go
type User struct {
	FirstName string
	LastName  string
	CreatedAt time.Time
}

func (u User) FullName() string {
	return u.FirstName + " " + u.LastName
}

func (u *User) Rename(firstName string) {
	u.FirstName = firstName
}
```

Go has no constructors. A function named `NewType` is a common convention when a type needs setup or validation.

```go
func NewUser(firstName, lastName string) *User {
	return &User{FirstName: firstName, LastName: lastName, CreatedAt: time.Now()}
}
```

### Embedding

A named struct field is ordinary composition. An embedded field promotes its fields and methods for convenient access; it is not inheritance.

```go
type Admin struct {
	User          // Embedded: admin.FullName() is available.
	Email string
}

type AuditRecord struct {
	User User // Named: access with record.User.FullName().
}
```

### JSON

The `encoding/json` package marshals exported struct fields. JSON tags change field names and other encoding behavior.

```go
type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

data, err := json.Marshal(note)
```

Use `MarshalJSON` only when the default representation is not sufficient, such as when keeping fields private. Implement it with an auxiliary exported struct to avoid recursive calls to `MarshalJSON`.

```go
type privateNote struct {
	title     string
	content   string
	createdAt time.Time
}

func (n privateNote) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Title     string    `json:"title"`
		Content   string    `json:"content"`
		CreatedAt time.Time `json:"created_at"`
	}{
		Title:     n.title,
		Content:   n.content,
		CreatedAt: n.createdAt,
	})
}
```

## Interfaces

An interface defines a set of methods. A type satisfies an interface implicitly when it has all of those methods; it does not need to declare that it implements the interface.

Name interfaces after the behavior they provide. Single-method interfaces commonly use the `-er` suffix, such as `io.Reader`, `io.Writer`, and `fmt.Stringer`.

```go
type Saver interface {
	Save(*Note) error
}

func CreateNote(store Saver, note *Note) error {
	return store.Save(note)
}
```

Guidelines:

1. Define interfaces near the code that consumes them.
2. Accept the narrowest interface needed; return concrete types unless callers need an abstraction.
3. Keep interfaces as small as the consumer needs. There is no strict method limit.
4. Name interfaces by behavior, such as `Reader` or `Writer`, not `IReader`.
5. Introduce an interface when a consumer needs a seam, not merely because a concrete type exists.
6. Use `any` instead of `interface{}` for the empty interface in Go 1.18+.

For a storage implementation with several operations, use focused capability interfaces. Compose them only where all operations are genuinely needed.

```go
type NoteLoader interface {
	Load(id string) (*Note, error)
}

type NoteDeleter interface {
	Delete(id string) error
}

type NoteLister interface {
	List() ([]*Note, error)
}

type NoteStorage interface {
	Saver
	NoteLoader
	NoteDeleter
	NoteLister
}
```

`NoteStorage` is useful for consumers that need the full API. A consumer that only saves notes should accept `Saver` instead.

### Type Switches and Assertions

Use a type switch when behavior depends on the concrete dynamic type stored in an interface.

```go
func printSomething(value any) {
	switch v := value.(type) {
	case int:
		fmt.Println("Integer:", v)
	case float64:
		fmt.Println("Float:", v)
	case string:
		fmt.Println("String:", v)
	default:
		fmt.Println("Other:", v)
	}
}
```

Use the comma-`ok` form of a type assertion when the type may not match. The single-result form panics on a mismatch.

```go
number, ok := value.(int)
if ok {
	fmt.Println(number + 1)
}
```

## Generics

Generics let a function or type work with values of several types while preserving type safety. A type parameter has a **constraint** that specifies the permitted types and operations.

```go
func Add[T int | float64 | string](a, b T) T {
	return a + b
}
```

Generics are for operations that are structurally the same across types. Interfaces are for behavior: values that provide particular methods. They solve different problems and can be used together.

## Arrays, Slices, and Maps

### Array

An array has a fixed length, and its length is part of its type. A slice is a small descriptor over an underlying array and is the usual collection type.

```go
var array [3]int          // Fixed length: [3]int{0, 0, 0}
numbers := []int{1, 2, 3} // Slice
numbers = append(numbers, 4)
```

`make` does **not** create arrays. Create an array with a declaration or an array literal. `make` creates initialized slices, maps, and channels; for a slice, its first argument after the type is its length and the optional second argument is its capacity.

```go
var scores [3]int           // [3]int{0, 0, 0}
names := [2]string{"Ada", "Lin"}

values := make([]int, 3)    // len 3, cap 3: []int{0, 0, 0}
buffer := make([]byte, 0, 64) // len 0, cap 64
```

A nil slice can be ranged over and appended to. `append` may allocate a new underlying array, so always use its returned slice.

#### Some useful functions:

len(arr) -> give current length of array
cap(arr) -> total capacity of an array
slicing can always reslice by selecting more items towards the right of the array but never the left.

1. sized array

```go
var arr1 [4]int = [4]int{1, 2, 3, 4}
slice1 := arr1[1:]
slice2 := slice1[:1]
slice2 := slice2[:3]
```

> // so here slice2 can reslice towards the right and grab more elements, but will never be able to grab `arr1[0]` because slice1 was created from `arr1[1:]`. So item at index zero is permanantly unreachable. Therefore, slice2 and slice1 cap() = 3, and arr1 cap() = 4

2. dynamic array

```go
var arr1 []int = []int{1, 2, 3, 4}
append(arr1, 5)
```

> append create a new array under the hood and returns that slice. To overwrite the previous array do arr1 = append(arr1, 5). Or get a new array with arr2 := append(arr1, 5).

### Maps

Maps associate unique, comparable keys with values. Valid key types include numbers, strings, pointers, arrays, structs whose fields are comparable, and interfaces containing comparable dynamic values. Slices, maps, and functions cannot be map keys.

```go
notes := make(map[string]Note)
notes["welcome"] = Note{Title: "Welcome"}

note, ok := notes["welcome"] // ok distinguishes a missing key from a zero value.
delete(notes, "welcome")
```

Use `make` to create an empty writable map. Its optional size hint reserves space for approximately that many entries; it does not limit the map's size. A map literal is useful when its initial entries are known.

```go
counts := make(map[string]int) // Empty, writable map.
seen := make(map[string]bool, 100) // Size hint, not a maximum.
ports := map[string]int{
	"http":  80,
	"https": 443,
}
```

A map's zero value is `nil`. Reading from, ranging over, deleting from, or calling `len` on a nil map is safe; a read returns the value type's zero value. Writing to a nil map panics, so initialize it with `make` or a literal first.

```go
var counts map[string]int
fmt.Println(counts["missing"]) // 0
delete(counts, "missing")      // Safe.
// counts["new"] = 1           // Panic: assignment to entry in nil map.
```

The comma-`ok` lookup distinguishes a missing key from a key whose value is the zero value. Map iteration order is deliberately unspecified, so do not rely on it. Maps are reference-like values: assigning a map or passing it to a function copies a header that refers to the same underlying data. Maps are not safe for concurrent reads and writes without synchronization.

```go
enabled := map[string]bool{"feature-a": false}
featureA, ok := enabled["feature-a"] // false, true
_, missing := enabled["feature-b"] // _, false
fmt.Println(featureA, ok, missing)

alias := enabled
alias["feature-a"] = true // Also changes enabled.
```
