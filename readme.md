# Introduction

Binder is a serialize & deserialize library to easily customize serialize / deserialize from *Type A* to *Type B*. Binder will be default binder for oscrud framework.

# Binder

Every type have default binding method included primitive types, slice, array and struct. By default, primitive types will convert to string and `reflect.Set()` to field. For slice, struct, array will be check `assignabltTo` only `reflect.Set()` to field. You can still customize all of the binder method, registered method will have higher priority than default method.

# Example

```go
package main

import (
    "log"
    "github.com/oscrud/oscrud-binder"
)

func main() {
    binder := binder.NewBinder()
    
    var data int32
    
    if err := binder.Bind(&data, "30"); err != nil {
        log.Println(err)
    }
    log.Println(data) // 30

    if err := binder.Bind(&data, int64(10)); err != nil {
        log.Println(err)
    }
    log.Println(data) // 10

    if err := binder.Bind(&data, "20.5"); err != nil {
        log.Println(err) // trying to convert 20.5 to int32
    }
}
```

# Registering Binder Method

```go
package main

import (
    "log"
    "fmt"
    "github.com/oscrud/oscrud-binder"
)

type AnyStruct struct {
    Data string
}

func main() {
    binder := binder.NewBinder()
    binder.Register(string(""), AnyStruct{}, func(raw interface{}) (interface{}, error) {
        return AnyStruct{fmt.Sprintf(raw)}
    })

    binder.Register(AnyStruct{}, string(""), func(raw interface{}) (interface{}, error) {
        strct := raw.(AnyStruct)
        return strct.Data
    })

    strct := new(AnyStruct)
    if err := binder.Bind(&strct, "will set to data"); err != nil {
        log.Println(err)
    }
    log.Println(strct.Data) // will set to data

    var str string
    if err := binder.Bind(&str, strct); err != nil {
        log.Println(err)
    }
    log.Println(str) // will set to data

    if err := binder.Bind(&strct, 10); err != nil {
        log.Println(err) // Trying to convert 10 to struct AnyStruct
    }
}
```

# Commands

```
$ go test --coverprofile=coverage.txt
$ go tool cover -html=coverage.txt -o coverage.html
```

# Test Coverage

```
=== RUN   TestFloat
--- PASS: TestFloat (0.00s)
=== RUN   TestInt
--- PASS: TestInt (0.00s)
=== RUN   TestUInt
--- PASS: TestUInt (0.00s)
=== RUN   TestString
--- PASS: TestString (0.00s)
=== RUN   TestBoolean
--- PASS: TestBoolean (0.00s)
=== RUN   TestSlice
--- PASS: TestSlice (0.00s)
=== RUN   TestArray
--- PASS: TestArray (0.00s)
=== RUN   TestMap
--- PASS: TestMap (0.00s)
=== RUN   TestStruct
--- PASS: TestStruct (0.00s)
=== RUN   TestTypeName
--- PASS: TestTypeName (0.00s)
=== RUN   TestRegister
--- PASS: TestRegister (0.00s)
=== RUN   TestError
--- PASS: TestError (0.00s)
PASS
coverage: 100.0% of statements
ok      github.com/oscrud/binder        0.173s
```