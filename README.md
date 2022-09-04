# concurrent map [![Build Status](https://travis-ci.com/smallnest/concurrent-map.svg?branch=master)](https://travis-ci.com/smallnest/concurrent-map)

> 一个线程安全的，支持泛型的，高性能的map库。基于`orcaman/concurrent-map`的实现和`alphadose/haxmap`的hash算法。

- Forked from [orcaman/concurrent-map](https://github.com/orcaman/concurrent-map) and supports any type of the key.
- The hash algorithm for key is copied from [alphadose/haxmap](https://github.com/alphadose/haxmap) and fixed for some types.


As explained [here](http://golang.org/doc/faq#atomic_maps) and [here](http://blog.golang.org/go-maps-in-action), the `map` type in Go doesn't support concurrent reads and writes. `safemap` provides a high-performance solution to this by sharding the map with minimal time spent waiting for locks, with full generic support.

Prior to Go 1.9, there was no concurrent map implementation in the stdlib. In Go 1.9, `sync.Map` was introduced. The new `sync.Map` has a few key differences from this map. The stdlib `sync.Map` is designed for append-only scenarios. So if you want to use the map for something more like in-memory db, you might benefit from using our version. You can read more about it in the golang repo, for example [here](https://github.com/golang/go/issues/21035) and [here](https://stackoverflow.com/questions/11063473/map-with-concurrent-access)

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

Running Benchmark:
```bash
go test -benchmem -bench github.com/smallnest/safemap
```

## license
MIT (see [LICENSE](https://github.com/orcaman/concurrent-map/blob/master/LICENSE) file)
