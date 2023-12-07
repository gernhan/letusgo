// Reading and writing files are basic tasks needed for
// many Go programs. First we'll look at some examples of
// reading files.

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Perhaps the most basic file reading task is
	// slurping a file's entire contents into memory.
	readEntireContentsIntoMemory()

	// You'll often want more control over how and what
	// parts of a file are read. For these tasks, start
	// by `Open`ing a file to obtain an `os.File` value.
	f := justOpenTheFile()

	// Read some bytes from the beginning of the file.
	// Allow up to 5 to be read but also note how many
	// actually were read.
	readFileToBuffer(5, f)

	// You can also `Seek` to a known location in the file
	// and `Read` from there.
	seekContentFromFile(f, 6, io.SeekStart, 2)

	// The `io` package provides some functions that may
	// be helpful for file reading. For example, reads
	// like the ones above can be more robustly
	// implemented with `ReadAtLeast`.
	readFileToBufferUsingReadAtLeast(f)

	// There is no built-in rewind, but `Seek(0, 0)`
	// accomplishes this.
	rewindFileReaderPointerToPosition(f, 0, io.SeekStart)

	// The `bufio` package implements a buffered
	// reader that may be useful both for its efficiency
	// with many small reads and because of the additional
	// reading methods it provides.
	effectiveSmallReadWithBufio(f)

	// Close the file when you're done (usually this would
	// be scheduled immediately after `Open`ing with
	// `defer`).
	defer func() {
		err := f.Close()
		if err != nil {
			fmt.Printf("\n Encouter error when closing file %v: error %v", f.Name(), err.Error())
		}
	}()
}

func readFileToBufferUsingReadAtLeast(f *os.File) {
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))
}

func effectiveSmallReadWithBufio(file *os.File) {
	r4 := bufio.NewReader(file)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))
}

func rewindFileReaderPointerToPosition(f *os.File, position int64, mode int) {
	_, err := f.Seek(position, mode)
	check(err)
}

func seekContentFromFile(f *os.File, position int64, startingIndexOfFile int, contentSize int) {
	startingPosition, err := f.Seek(position, startingIndexOfFile)
	check(err)
	buffer := make([]byte, contentSize)
	length, err := f.Read(buffer)
	check(err)
	fmt.Printf("%d bytes from position %d: ", length, startingPosition)
	fmt.Printf("%v\n", string(buffer[:length]))
}

func readFileToBuffer(bufferSize int, f *os.File) {
	for {
		b1 := make([]byte, bufferSize)
		n1, err := f.Read(b1)
		if err == io.EOF {
			break
		}
		fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
	}
}

func justOpenTheFile() *os.File {
	f, err := os.Open("text.txt")
	check(err)
	fmt.Println(*f)
	return f
}

func readEntireContentsIntoMemory() {
	dat, err := ioutil.ReadFile("text.txt")
	check(err)
	fmt.Println(string(dat))
}
