package main
/*
	Objective: 
		make a experiment to help me keep using defer as usual in sussesive file operation circustums.

	Explanation: 
		To use "defer" to deal with os.File.Close in sussesive file operation circustums,
		I make use of some characteristic in os.File.Open and defer to make a util function 
	
			- os.File.Open (fileName string) (*File, error)
				Because open file function return the pointer type of File,
				we can use the pointer to close the exact File with the address in Heap memory.
			
			- defer will be execute in the order of "Last in first out".

			- defer + func()
				When defer function is first invoke (not execute yet),
				we can input some custom parameters to defer function call.

				Thus, I input the file pointer (*File) into the function call
				to help me close the exact file I want.

					example: defer func(){ ...}(*File)
*/

import (
	"fmt"
	"log"
	"os"
	"io"
)

func fileClose (file *os.File, fileName string) {
	fmt.Printf("Pointer for filename %s is %v\n", fileName, file)
	if err := file.Close(); err != nil {
		log.Fatalf("File close %s failed\n", fileName)
	}
}

func main() {
	const bufferSize = 2048
	buffer := make([]byte, bufferSize)
	n1 := "test.txt"
	n2 := "test.json"
	n3 := "test.html"
	
	// Process file 1
	// Close with pointer &{0x14000062240}
	f, err := os.Open(n1)
	
	if err != nil {
		log.Fatal(err)
	}

	defer fileClose(f, n1)

	for {
		contents, err := f.Read(buffer)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		fmt.Println(string(buffer[:contents]))
	}

	// Process file 2
	// Close with pointer &{0x140000621e0}
	f, err = os.Open(n2)

	if err != nil {
		log.Fatal(err)
	}

	defer fileClose(f, n2)

	for {
		contents, err := f.Read(buffer)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		fmt.Println(string(buffer[:contents]))
	}

	// Process file 3
	// Close with pointer &{0x14000062180}
	f, err = os.Open(n3)

	if err != nil {
		log.Fatal(err)
	}

	defer fileClose(f, n3)

	for {
		contents, err := f.Read(buffer)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		fmt.Println(string(buffer[:contents]))
	}

}