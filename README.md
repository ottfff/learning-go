# Go cheat sheet

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

```go
m := sync.Map{} //any to any
m.Store("a", "1")
m.Store(1, "a")

if v, ok := m.Load("a"); ok { //contains
    fmt.Println("a =", v)
}

m.Delete(1)
```

## Interface

Type cast:
```go
var i interface{} = 10
var j interface{} = "10"
v := i.(int) //ok 
u := j.(int) //panic
```
