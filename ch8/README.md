# Go Memory Management Examples

This folder contains code examples from the chapter on memory management in Go, focusing on garbage collection, stack and heap allocation, and memory optimization techniques.

## Table of Contents

- [Memory Ballast](#memory-ballast)
- [Memory Arenas](#memory-arenas)
    - [Importing the Arena Package](#importing-the-arena-package)
    - [Creating a New Arena](#creating-a-new-arena)
    - [Allocating a New Reference for a Struct in the Arena](#allocating-a-new-reference-for-a-struct-in-the-arena)
    - [Creating a Slice with a Predetermined Capacity in the Arena](#creating-a-slice-with-a-predetermined-capacity-in-the-arena)
    - [Freeing the Arena](#freeing-the-arena)
    - [Cloning an Object from the Arena to the Heap](#cloning-an-object-from-the-arena-to-the-heap)
    - [Using the Address Sanitizer](#using-the-address-sanitizer)
- [GC Environment Variables](#gc-environment-variables)
    - [GODEBUG](#godebug)
    - [GOGC](#gogc)
    - [GOMEMLIMIT](#gomemlimit)

## Memory Ballast

Memory ballast is used to artificially inflate the heap size to optimize the behavior of the garbage collector (GC).

```go
ballast := make([]byte, 10<<30)
```

## Memory Arenas

Memory arenas are a useful tool for allocating objects from a contiguous region of memory and freeing them all at once with minimal memory management overhead.

### Importing the Arena Package

First, import the arena package:

```go
import "arena"
```

### Creating a New Arena

Create a new arena:

```go
mem := arena.NewArena()
```

### Allocating a New Reference for a Struct in the Arena

Allocate a new reference for a struct type in the arena:

```go
p := arena.New[Person](mem)
```

### Creating a Slice with a Predetermined Capacity in the Arena

Create a slice with a predetermined capacity in the arena:

```go
slice := arena.MakeSlice[string](mem, 100, 100)
```

### Freeing the Arena

Free the arena to deallocate all objects at once:

```go
mem.Free()
```

### Cloning an Object from the Arena to the Heap

Clone an object from the arena to the heap:

```go
p1 := arena.New[Person](mem) // arena-allocated
p2 := arena.Clone(p1) // heap-allocated
```

### Using the Address Sanitizer

Use the address sanitizer to detect issues with memory access after freeing the arena:

```go
type T struct {
    Num int
}
func main() {
    mem := arena.NewArena()
    o := arena.New[T](mem)
    mem.Free()
    o.Num = 123 // <- this is a problem
}
```

Run the program with the address sanitizer:

```sh
go run -asan main.go
```

## GC Environment Variables

### GODEBUG

The `GODEBUG` environment variable provides insights into the inner workings of the Go runtime, including garbage collection processes.

Enable GC tracing:

```sh
export GODEBUG=gctrace=1
```

### GOGC

The `GOGC` environment variable controls the aggressiveness of the garbage collection process.

Set `GOGC` to run GC more frequently:

```sh
export GOGC=50
```

### GOMEMLIMIT

The `GOMEMLIMIT` environment variable sets a soft memory limit for the Go runtime.

Set a memory limit:

```sh
export GOMEMLIMIT=500MiB
```
