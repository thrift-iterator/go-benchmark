# go-benchmark

gogoprotobuf

```
5000000	       366 ns/op	     144 B/op	      12 allocs/op
```

thrift 

```
1000000	      1549 ns/op	     528 B/op	       9 allocs/op
```

thrifter static codegen

```
5000000	       389 ns/op	     192 B/op	       6 allocs/op
```

thrifter dynamic codegen

```
2000000	       585 ns/op	     192 B/op	       6 allocs/op
```