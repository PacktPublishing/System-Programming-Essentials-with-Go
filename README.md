# System Programming Essentials with Go	

<a href="https://www.packtpub.com/product/system-programming-essentials-with-go/9781837634132"><img src="https://content.packt.com/_/image/original/B21662/cover_image_large.jpg" alt="no-image" height="256px" align="right"></a>

This is the code repository for [System Programming Essentials with Go](https://www.packtpub.com/product/system-programming-essentials-with-go/9781837634132), published by Packt.

**System calls, networking, efficiency, and security practices with practical projects in Golang**

## What is this book about?
From file operations to process management and network programming, this hands-on guide equips software engineers with the skills to build efficient, reliable applications and optimize their performance.	

This book covers the following exciting features:
* Understand the fundamentals of system programming using Go
* Grasp the concepts of goroutines, channels, data races, and managing concurrency in Go
* Manage file operations and inter-process communication (IPC)
* Handle USB drives and Bluetooth devices and monitor peripheral events for hardware automation
* Familiarize yourself with the basics of network programming and its application in Go
* Implement logging, tracing, and other telemetry practices
* Construct distributed cache and approach distributed systems using Go

If you feel this book is for you, get your [copy](https://www.amazon.com/System-Programming-Essentials-networking-efficiency/dp/1837634130/ref=tmm_pap_swatch_0?_encoding=UTF8&dib_tag=se&dib=eyJ2IjoiMSJ9.V74ree9n-By3iEAv6O5AZ80GMgp-RQ06f2ateXTAAu-samELuP-q_zhuOyaBqsRxUiyqF60yvStRz62CHPeo2F6qEiY2uqxKvEe8ib6CkArIwnWzGYNMgC_S2sdL11uAZVOb56FzNwZO_RdXKjlQSko8ev7kZgSPHqN_VZfNbBM_5QsHLG3vvDsYpU9kgAmzldh2HNPEzCkfO76LsRQ2Ydx0E4tZtkRLxDTaLGm8txc.pUa0SG0zCGPkykCKJUYvWiiz1JTbdEcS6L0Z0QT27i4&qid=1717763508&sr=1-1) today!


## Instructions and Navigations
All of the code is organized into folders. For example, ch2.

The code will look like the following:
```
func main() {
  cache := NewCache(5) // Setting capacity to 5 for LRU
  cache.startEvictionTicker(1 * time.Minute)
}
```

**Following is what you need for this book:**
This book is for software engineers looking to expand their understanding of system programming concepts. Professionals with a coding foundation seeking profound knowledge of system-level operations will also greatly benefit. Additionally, individuals interested in advancing their system programming skills, whether experienced developers or those transitioning to the field, will find this book indispensable.

With the following software and hardware list you can run all code files present in the book (Chapter 1-15).
### Software and Hardware List
| Chapter | Software required | OS required |
| -------- | ------------------------------------ | ----------------------------------- |
| 1-15 | Golang (1.16+) | Windows, macOS, or Linux |



### Related products
* Go Programming - From Beginner to Professional [[Packt]](https://www.packtpub.com/product/go-programming-from-beginner-to-professional-second-edition/9781803243054) [[Amazon]](https://www.amazon.com/Go-Programming-Beginner-Professional-everything/dp/1803243058/ref=sr_1_1?crid=GB1XN1O9W9B9&dib=eyJ2IjoiMSJ9.9qi4XiKwA90sP2288upVW_T2gM08M8CA79EHqeiQtiwwVw_rJ1IjaiSkOAr3httgpBruqGJgXutvAjNqdRMcy2xycSiwsAp_A0s3h_F706Ki4YQ_x25os96pxyb120GCT3hrAbbpBWwTzA0ICOOMHOrTFYY9zFZ5jrQDfmKag2gZ882ir1oJjTG04rDbH8Bq17xwYmJTyHcayDjQ4UMhoHUeJs0dgjngqO8KNnJ7rjw.n3teqRtS-Jo-sOXNURrArPERLqQeG8eRaAGAxzsIiBE&dib_tag=se&keywords=Go+Programming+-+From+Beginner+to+Professional&qid=1717763707&s=books&sprefix=go+programming+-+from+beginner+to+professional%2Cstripbooks-intl-ship%2C292&sr=1-1)

* gRPC Go for Professionals [[Packt]](https://www.packtpub.com/product/grpc-go-for-professionals/9781837638840) [[Amazon]](https://www.amazon.com/gRPC-Professionals-Implement-production-grade-microservices/dp/1837638845/ref=tmm_pap_swatch_0?_encoding=UTF8&dib_tag=se&dib=eyJ2IjoiMSJ9.b-buosIEMYbEsWl3m8HzWaFN9uzHPsAauCF5bC0CAh_GjHj071QN20LucGBJIEps.ywsVDnPhSUhXXA2CTXsXz71EP-qTv7ZIhqOC40RnqPY&qid=1717763771&sr=1-1)

## Get to Know the Author
**Alex Rios**
is an established Brazilian software engineer with a 15-year track record of success in large-scale solution development. Alex specializes in Go and creates high-throughput systems that address diverse needs across fintech, telecom, and gaming industries. As a Staff Engineer at Stone Co., Alex applies his expertise using unconventional system designs, ensuring top-notch delivery. Also, Alex uses his expertise to evaluate books and publications as a technical reviewer. Alex is an enthusiastic community member, actively participating in its growth and development as Curitiba's Go meetup organizer. His dedication is evident in his regular presence as a speaker at major national tech events like GopherCon Brazil.

