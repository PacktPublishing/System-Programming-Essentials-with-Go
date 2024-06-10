package main

import (
	"bytes"
	"fmt"
	"sync"
)

type BufferPool struct {
	pool sync.Pool
}

func NewBufferPool() *BufferPool {
	return &BufferPool{
		pool: sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
	}
}

func (bp *BufferPool) Get() *bytes.Buffer {
	return bp.pool.Get().(*bytes.Buffer)
}

func (bp *BufferPool) Put(buf *bytes.Buffer) {
	buf.Reset()
	bp.pool.Put(buf)
}

func ProcessData(data []byte, bp *BufferPool) {
	buf := bp.Get()
	defer bp.Put(buf)

	buf.Write(data)
	fmt.Println(buf.String())
}

func main() {
	bp := NewBufferPool()
	data := []byte("Hello, World!")
	ProcessData(data, bp)
}
