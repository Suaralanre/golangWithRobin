package main

import (
	"fmt"
	"os"
	"sync"
	"unicode"
)

var sampleString = `Golden sunsets paint the sky,
Peaceful evening, twilight high.
Stars appear, like diamonds bright,
Night's soft whisper, a gentle delight.
Silver moonbeams dance upon the sea,
A path of light, for dreams to be.
The world is hushed, in quiet sleep,
As the night's mysteries, our souls keep.
In this calm moment, I find my peace,
A sense of freedom, my worries release.
So let the night, with its darkness fall,
For in its shadows, I hear my heart's call.`

func main() {
	wg := &sync.WaitGroup{}

	ch := make(chan rune)

	wg.Add(1)

	go func() {
		defer wg.Done()
		for _, char := range sampleString {
			ch <- char
		}

		close(ch)
		
	}()
	
	go readChan(ch, wg)

	wg.Wait()
}

// createTextFile creates and returns a file handler.
func createTextFile(filename string) *os.File {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "File could not be created:", err)
		return nil
	}
	return file
}

// readChan reads from the channel and writes to the appropriate files.
func readChan(ch chan rune, wg *sync.WaitGroup) {
	defer wg.Done()

	upperFile := createTextFile("upper.txt")
	defer upperFile.Close()
	lowerFile := createTextFile("lower.txt")
	defer lowerFile.Close()

	upperLineCount := 0
	lowerLineCount := 0

	for char := range ch {
		if upperLineCount >= 100 || lowerLineCount >= 100 {
			break
		}

		if unicode.IsUpper(char) {
			// Write to upper.txt and count lines.
			fmt.Fprintln(upperFile, string(char))
			upperLineCount++

		} else if unicode.IsLower(char) {
			// Write to lower.txt and count lines.
			fmt.Fprintln(lowerFile, string(char))
			lowerLineCount++
		}
	}

	if upperLineCount >= 100 {
		fmt.Println("upper.txt reached 100 lines first!")
	} else if lowerLineCount >= 100 {
		fmt.Println("lower.txt reached 100 lines first!")
	}
}
