# Go cheat sheet

## Table of Contents

- [Types](#types)
- [Zero value / Default value](#zero-value--default-value)
- [Slice](#slice)
- [Rune](#rune)
- [Map](#map)
  - [sync.Map](#syncmap)
- [Interface](#interface)
- [Concurrency](#concurrency)
  - [WaitGroup](#waitgroup)
  - [Mutex](#mutex)
  - [Pool](#pool)
  - [Once / OnceFunc / memoize function](#once--oncefunc--memoize-function)
    - [Once](#once)
    - [OnceFunc](#oncefunc)
    - [Memoize function](#memoize-function)
  - [Atomic](#atomic)
  - [Semaphore](#semaphore)

## Types

* **int8**
* **int16**
* **int32**
* **int64**
* **uint8**
* **uint16**
* **uint32**
* **uint64**
* **byte**: synonym uint8
* **rune**: synonym int32
* **int** synonym int32 or int64 based on platform
* **uint**: synonym uint32 or uint64 based on platform
* **float32**
* **float64**
* **complex64**
    ```go 
    var f complex64 = 1+2i
    ```
  Real part **float32**, imaginary part **float32**.
* **complex128**
    ```go
    var g complex128 = 4+3i
    ```
* **bool**
    ```
    &   // bitwise AND
    |   // bitwise ИЛИ
    ^   // bitwise exclusive OR (XOR)
    &^  // AND NOT
    <<  
    >>  
    ```
* **string**  
  Immutable.  
  Literal located in `.rodata`.

`type any = interface{}`

## Zero value / Default value

| Type             | Zero value                | Example                |
|------------------|---------------------------|------------------------|
| int, int32 …     | `0`                       | `var i int`            |
| float32, float64 | `0.0`                     | `var f float64`        |
| bool             | `false`                   | `var b bool`           |
| string           | `""`                      | `var s string`         |
| *T               | `nil`                     | `var p *int`           |
| []T              | `nil`                     | `var s []int`          |
| map[K]V          | `nil`                     | `var m map[string]int` |
| chan T           | `nil`                     | `var ch chan int`      |
| interface{}      | `nil`                     | `var i interface{}`    |
| func             | `nil`                     | `var f func()`         |
| struct           | all fields are zero value | `var s MyStruct`       |
| array [N]T       | all items are zero value  | `var a [3]int`         |

## Slice

```go
type slice struct {
    ptr    *T  // pointer to first element
    len    int // current size
    cap    int // available capacity
}
```

```go
//init
var s1 []int // nil slice, len=0, cap=0
s2 := make([]int, 4, 7) // zero values, len=4, cap=7
s3 := []int{1, 2, 3, 4, 5} // len=5, cap=5

// append. Capacity grows twice (or shorten). After 1024 grow less.
s = append(s, 2, 3, 4, 5)
len(s)
cap(s)

//copy
copy(dst, src)
s = s[:]

//remove i-th element
s = append(s[:i], s[i+1:]...)
```

## Rune

`type rune = int32`

UTF8 use 1-4 bytes for symbols.

Heading 0 used for 1-byte symbol (128 ASCII symbols).

Heading 10 used for middle byte of some symbol.

Heading 110 used for 2-byte symbols (e.g. Cyrillic symbols).

Heading 1110 used for 3-byte symbol.

Heading 11110 used for 4-byte symbol.

```go
s := "😀"
len(emoji) // 4 - number of bytes
len([]rune(emoji)) // 1 - number of symbols
utf8.RuneCountInString(emoji) // 1

[]byte("hello")     // string → []byte
string([]byte{...}) // []byte → string

for i, r := range s {
	// iteration over runes
}
```

## Map

```go
// key - string, value - int.
var m1 map[string]int // nil
m2 := make(map[string]int)
m3 := make(map[string]int, 100)
m4 := map[string]int{
    "one": 1,
    "two": 2,
}

v, ok := m["key"] //get
delete(m, "key") //remove

for k, v := range m {
    fmt.Println(k, v)
}

var m map[string]int
i, ok := map["key"] // 0 false
m["key"] = 1 // panic: assignment to entry in nil map
```

### sync.Map
Read-optimized concurrent map.

For write-optimized solution better to use map + RWMutex

* read-only part 
* dirty part
* promotion logic.
  Frequent reads from dirty part promote to read part.
```go
m := sync.Map{} //any to any
m.Store("a", "1")
m.Store(1, "a")

if v, ok := m.Load("a"); ok { //contains
    fmt.Println("a =", v)
}

m.Delete(1)
var swapped bool = m.CompareAndSwap(1, "oldValue", "newValue")
var deleted bool = m.CompareAndDelete(1, "oldValue")
value, keyWasPresent := m.LoadAndDelete(1)
prevValueOrNewValue, keyWasPresent := m.LoadOrStore(1, "newValue")
prevValueOrNil, keyWasPresent := m.Swap(3, "newValue")
```

## Interface

Type cast:
```go
var i interface{} = 10
var j interface{} = "10"
v := i.(int) //ok 
u := j.(int) //panic
```

Interfaces in go have duck typing
```go
type Animal interface {
    Say() (n string, err error)
}
type Duck struct{}
func (duck Duck) Say() (string, error) {
    return "Krya", nil
}

//verify in compile stage that Duck implements Animal
var _ Animal = (*Duck)(nil)
```
```go
type any = interface {}

type error interface {
    Error() string
}
```

Nil trap
```go
var i interface {} // nil
i == nil // true

// bad function. Don't return nil of certain type because returnsError() != nil
func returnsError() error {
    var err *MyError = nil
    return err // interface error != nil
}
if returnsError() == nil { 
	// basically such stile is good, but if you don't manage returnsError() function then better to check by reflect
	// reflect.ValueOf(err).IsNil()
}
```

## Concurrency

### WaitGroup
```go
var wg sync.WaitGroup
wg.Add(1)
wg.Done()
wg.Wait()
```
Always pass WaitGroup by pointer.

### Mutex
```go
var mu sync.Mutex
mu.Lock()
mu.Unlock()
my.TryLock()
```
```go
var rwMutex sync.RWMutex
rwMutex.Lock()
rwMutex.Unlock()
rwMutex.TryLock()
rwMutex.RLock()
rwMutex.RUnlock()
rwMutex.TryRLock()
rwMutex.RLocker()
```

### Pool
```go
type A struct {
    name string
}
var pool = sync.Pool{
    New: func() interface{} {
        return A{name: "a0"}
    },
}
o0 := pool.Get().(A) // "a0" as default value
pool.Put(A{"a1"})
pool.Put(A{"a2"})
o1 := pool.Get().(A) // random value, but most likely "a1"
o2 := pool.Get().(A) // next random value "a2"
```

### Once / OnceFunc / memoize function

#### Once
```go
var once sync.Once // sharing value
func f() {
    once.Do(func() {
        fmt.Println("Run only once")
    })
}

f() // "Run only once"
f() // nothing
f() // nothing
```

#### OnceFunc
```go
f := sync.OnceFunc(func() {
    fmt.Println("Run only once")
})

f() // "Run only once"
f() // nothing
f() // nothing
```

#### Memoize function
To create function with cached resul use next pattern
```go
var once sync.Once
var result int

func getValue() int {
    once.Do(func() {
        fmt.Println("initializing")
        result = 42
    })
    return result
}

fmt.Println(getValue()) // initialize 42
fmt.Println(getValue()) // returns cached 42
```

### Atomic

| Function                                                       | Description                                |
|----------------------------------------------------------------|--------------------------------------------|
| `atomic.LoadInt32(addr *int32) int32`                          | read `int32`                               |
| `atomic.StoreInt32(addr *int32, val int32)`                    | write `int32`                              |
| `atomic.AddInt32(addr *int32, delta int32) int32`              | add `delta` and return new value           |
| `atomic.CompareAndSwapInt32(addr *int32, old, new int32) bool` | compare and swap if equals `old` |
| analogs `int64`, `uint32`, `uint64`, `uintptr`, `Pointer` e.g. |                                            |

```go
var counter int64 = 0
atomic.AddInt64(&counter, 1)
```

### Semaphore
There is no default implementations.
Example of custom implementation:
```go
type Semaphore chan struct{}

func NewSemaphore(max int) Semaphore {
    return make(Semaphore, max)
}

func (s Semaphore) Acquire() {
    s <- struct{}{} // acquiring slot
}

func (s Semaphore) Release() {
    <-s // release the slot
}
```
