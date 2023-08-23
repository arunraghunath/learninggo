package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type MyCustomReader struct {
	reader  bufio.Reader
	scanner bufio.Scanner
	length  int
}

func New(rd io.Reader, len int) *MyCustomReader {
	reader := *bufio.NewReader(rd)
	scanner := bufio.NewScanner(&reader)

	myreader := &MyCustomReader{
		reader:  reader,
		scanner: *scanner,
		length:  len,
	}
	myreader.scanner.Split(myreader.SplitFunc)
	return myreader

}

func (my *MyCustomReader) Read() bool {
	return my.scanner.Scan()
}

func (my *MyCustomReader) Bytes() []byte {
	return my.scanner.Bytes()
}

func (my *MyCustomReader) SplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if len(data) > my.length {
		// We have a full newline-terminated line.
		return my.length, data[0:my.length], nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data[0:], nil
	}
	// Request more data.
	return 0, nil, nil
}

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error opening file")
		return
	}
	myread := New(file, 4)
	for myread.Read() {
		myread.Bytes()
	}
}
