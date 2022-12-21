# safemap [![Build Status](https://travis-ci.com/smallnest/safemap.svg?branch=master)](https://travis-ci.com/smallnest/safemap)

> 一个线程安全的，支持泛型的，高性能的map库。基于`orcaman/concurrent-map`的实现和`github.com/dolthub/maphash`的hash算法。

## usage

Import the package:

```go
import (
	"github.com/smallnest/safemap"
)

```

```bash
go get github.com/smallnest/safemap@latest
```

The package is now imported under the "safemap" namespace.

## example

```go

	// Create a new map (K type is string, and V type is string too).
	m := safemap.New[string,string]()

	// Sets item within map, sets "bar" under key "foo"
	m.Set("foo", "bar")

	// Retrieve item from map.
	bar, ok := m.Get("foo")

	// Removes item under key "foo"
	m.Remove("foo")

	//  Create a new map (K type is int, and V type is string too).
	m2 := safemap.New[int,string]()

	// Sets item within map, sets "bar" under key 1
	m2.Set(1, "bar")

	// Retrieve item from map.
	bar, ok = m2.Get(1)

	// Removes item under key 1
	m2.Remove(1)

```

For more examples have a look at safemap_test.go.

Running tests:

```bash
go test github.com/smallnest/safemap
```


## license
MIT (see [LICENSE](https://github.com/orcaman/concurrent-map/blob/master/LICENSE) file)
