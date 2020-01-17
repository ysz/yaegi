package main

import "net"

var lookupHost = net.DefaultResolver.LookupHost

func main() {
	println(lookupHost != nil)
}

// Output:
// true
