package files

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

func FileCreation(nameOfFile string, wg *sync.WaitGroup) {
	time.Sleep(5 * time.Second)
	file, err := os.Create(nameOfFile)
	if err != nil {
		fmt.Println("Error while creating the file : ", err)
	}
	defer file.Close()
	content := "hello world by abhianv"
	fmt.Println("File created successfully")
	numberOfBytesWrittenIntoFile, err := io.WriteString(file, content+"\n")
	fmt.Println("Number of bytes written into file is ", numberOfBytesWrittenIntoFile)
	if err != nil {
		fmt.Println("Error while writing the file")
		return
	}

	fmt.Println("Successfully written in file")
	defer wg.Done()

}
func FileReading(nameOfFile string) {
	file, err := os.Open(nameOfFile)
	if err != nil {
		fmt.Println("Error while opening file : ", err)
		return
	}
	defer file.Close()
	buffer := make([]byte, 1024)

	//Read the full content into the buffer
	for {
		numberOfBytesRead, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error while reading the file", err)
			return
		}

		//Process the read content
		fmt.Println(string(buffer[:numberOfBytesRead]))
	}

}

func FileHandling() {
	var wg sync.WaitGroup
	wg.Add(1)
	go FileCreation("example.txt", &wg)
	wg.Wait()
	fmt.Println("Reading from the file -------")
	FileReading("example.txt")
}
