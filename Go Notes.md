# Go Notes:

========

## quick commands:

~$ `go run app.go` or `go run .` (which runs main module)

~$ `go run .` -> runs main package in this module

~$ `go build` -> build native binaries (.exe)

~$ `go mod init mydomain.com/project-name` -> initialises module (basically a project)

~$ go get module-origin (`go get github.com/Pallinder/go-randomdata`) -> gets specific library

~$ `go get` = `npm install` -> install all needed dependencies in `go.mod`

Go has packages that contain files. Like Java packages, they can contain multiple files (e.g., a service package with .imageProcessor, imageInference, etc.).

- package main: tell Go this is main entryPoint of project
- main module: needed for Go build (\~$ go build)

> The information above are quick tips about running, compiling and using Go on daily

## Types \& Null Values:

=============================================

- int: number without decimals
- float64: decimal number
- string: a string "Hello" or \`Hello\`
- bool: true or false

## Niche types:

unit: unsigned integer, which means a strictly non-negative number
int32: 32-bit signed integer (from -2^32-1 to +2^32-1), about -2 billion to 2 billion

rune: an alias for int32; represents a Unicode code point (i.e. a single character)

uint32: 32-bit unsigned integer (0 to about 4 billion)

int64: 64-bit signed int (-9\*10^18 to +9\*10^18) = 9 quintillion

int8: 8-bit signed integer (-128 to 127) -> 256 combinations

uint8: 8-bit unsigned integer (0 to 256) -> 256 combinations

Null values:

\---------------

int: 0

float64: 0.0

string: ""

bool: false

### Functions:

====================

Example signature ->

func ReadBalanceFromFile(filename string) (string, error) {return "", nil}

- Functions can return an error (must be caught in the func and the caller func with if err != nil {})

- For function inside packages only functions starting with capital letter can start ex: `GetUser(principal string) (user User, err error) {}`

#### Defer, Panic, Recover:

Defer:

A defer statement pushes a function call onto a list. The list of saved calls is executed after the surrounding function returns. Defer is commonly used to simplify functions that perform various clean-up actions.

- Example:

  Closing file buffers:

  ```go
  defer src.Close()
  ```

There are three simple rules:

1. A deferred function’s arguments are evaluated when the defer statement is evaluated.

2. Deferred function calls are executed in Last-In, First-Out order after the surrounding function returns.

3. Deferred functions may read and assign to the returning function’s named return values.

```go
func c() (i int) {
    defer func() { i++ }()
    return 1
}
```

## Pointers & References:

pointers are good for:

- avoiding unnecessary value copies
- directly mutate values

pointer example:
age := 32

compyters stores in memory value 32 at a specific memory address [32][0xc000018050]

agePointer := &age -> assigned the memory address

- Go functions default behavior is pass by reference. Therefore:

```go
age := 32
func ComputeAge(age int) {}
```

> will create a local scope copy of age.

- If our app needs optimisation:

for simple types such as `float`, `string`, `int` etc., we can then pass the address reference directly and edit directly the original value in memory without creating a copy. So if we are doing a lot of matrix operations etc. then editing the original value

```go
age := 32
func ComputeAge(&age int) {
    value := *age -> gets value from the passed pointer
}
```

## Structs:

```go
type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func (u user) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)
}
```

### Receiver:

func (u user) -> u here is the called the `Receiever`. `Receiver` are like any parameters, so if a method should edit a struct we have to pass the struct as a pointer `(u *user)`. For regular reads passing a copy `(u user)` is usually fine, unless the object is very big in memory. Think of a struct that stores a large matrix from the ML side of stuff needed for heavy computation.

### new StructType pattern:

Go doesn't really have constructors, but we can define our own construction reusable methods with a `factory function`, following the naming `New -> func New(...) user {}`

here we can either return a copy or a \* pointer, to avoid copying a value multiple times.

### Struct Embedding :

You can embed structs as part of other structs.

```go
type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
    email    string
    password string
    user     User
}
```

`user User` is explicit embedding which gives a name to the field, meaning to access it we would need to do `admin.user.function_name`,

With **anonymous** embedding, we only need to name specify the type, which will then give us direct access to its public fields & methods on Admin instance directly

### Struct to Json :

The idomatic approach is to use `encoding/json` module with `json.Marshal(myStruct)` to convert struct to json object, but the exportable fields must be made public (private fields - lower case - are never Marshalled)

```go
Option 1: Export the fields and add JSON tags:
type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

Then update New to use Title:, Content:, CreatedAt: and the getter methods to reference the exported fields.

Option 2: Keep fields private but implement MarshalJSON:
func (n *Note) MarshalJSON() ([]byte, error) {
	type NoteAlias Note
	return json.Marshal(&NoteAlias{
		Title:     n.title,
		Content:   n.content,
		CreatedAt: n.createdAt,
	})
}
```

> Struct Tags can also be added to struct to change format of json fields

```go
type Note struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
```

# Interfaces

An interface defines a set of methods. A type satisfies an interface implicitly when it has all of those methods; it does not need to declare that it implements the interface.

Name interfaces after the behavior they provide. Single-method interfaces commonly use the `-er` suffix, such as `io.Reader`, `io.Writer`, and `fmt.Stringer`. For a `Save` method, `Saver` is a natural name.

```go
type Saver interface {
	Save(*Note) error
}
```

Key conventions:

1. Keep interfaces as small as the consumer needs. One or two methods is common, but there is no strict method limit.
2. Define an interface near the code that consumes it, not automatically next to the concrete type.
3. Accept the narrowest interface needed; return concrete types unless callers need an abstraction.
4. Name interfaces by behavior: `Reader`, `Writer`, `Stringer`, not `IReader`.
5. Do not introduce an interface only to abstract a single implementation. Add one when a consumer needs a seam, such as substituting a dependency in a test or using another implementation.
6. Use `any` instead of `interface{}` for the empty interface in Go 1.18+.

```go
// This function only needs to save notes, so it accepts only a Saver.
func CreateNote(store Saver, note *Note) error {
	return store.Save(note)
}
```

For a storage implementation with several operations, define focused capability interfaces:

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
```

Compose them only where code genuinely needs the full set of operations:

```go
type NoteStorage interface {
	Saver
	NoteLoader
	NoteDeleter
	NoteLister
}
```

`NoteStorage` is not inherently bad. The problem is requiring every caller to depend on it when that caller only needs one capability.
