// Program in GO language to demonstrates how to use base log package.
package main

import (
	"log"
)

func init() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")
}
func main() {
	// Println writes to the standard logger.
	log.Println("main started")

	// Fatalln is Println() followed by a call to os.Exit(1)
	log.Fatalln("fatal message")

	// Panicln is Println() followed by a call to panic()
	log.Panicln("panic message")
}
