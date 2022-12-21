module github.com/smallnest/safemap/benchmark

go 1.19

require (
	github.com/alphadose/haxmap v1.2.0
	github.com/cornelk/hashmap v1.0.8
	github.com/puzpuzpuz/xsync/v2 v2.4.0
	github.com/smallnest/safemap v0.0.0-20220904104306-ea4af14ec8ea
)

require (
	github.com/dolthub/maphash v0.0.0-20221220182448-74e1e1ea1577 // indirect
	github.com/orcaman/concurrent-map v1.0.0 // indirect
	github.com/orcaman/concurrent-map/v2 v2.0.1 // indirect
	golang.org/x/exp v0.0.0-20221031165847-c99f073a8326 // indirect
)

replace github.com/smallnest/safemap => ./..
