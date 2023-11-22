## Race Detection
let’s use the race detector:

go test -race

The result in the console will be something like:

==================

WARNING: DATA RACE

Read at 0x00c00000e288 by goroutine 9:

example1.PackItems.func1()

      /tmp/main.go:35 +0xa8 

example1.PackItems.func2()

      /tmp/main.go:45 +0x47 



Previous write at 0x00c00000e288 by goroutine 8:

example1.PackItems.func1()

      /tmp/main.go:39 +0xba 

example1.PackItems.func2()

      /tmp/main.go:45 +0x47 



// Other lines omitted for brevity

The output could be intimidating at first glance, but the most revealing information for now is the message WARNING: DATA RACE . 
