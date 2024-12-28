package main

import (
	"context"
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

const (
	MaxLines  = 100
	UpperFile = "upper.txt"
	LowerFile = "lower.txt"
)

type FileWriter struct {
	file  *os.File
	lines int
}

type ChannelProcessor struct {
	ch      chan rune
	wg      *sync.WaitGroup
	cancel  context.CancelFunc
	content string
}

type IProcessor interface {
	WriteToChannel(ctx context.Context)
	ReadFromChannel(ctx context.Context)
}

func NewFileWriter(filename string) *FileWriter {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating file %s: %v\n", filename, err)
		os.Exit(1)
	}
	return &FileWriter{file: file}
}

func (fw *FileWriter) Write(char rune) {
	if _, err := fmt.Fprintln(fw.file, string(char)); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing to file %s: %v\n", fw.file.Name(), err)
	}
	fw.lines++
}

func NewChannelProcessor(content string) *ChannelProcessor {
	return &ChannelProcessor{
		ch:      make(chan rune),
		wg:      &sync.WaitGroup{},
		content: content,
	}
}

func (cp *ChannelProcessor) WriteToChannel(ctx context.Context) {
	defer cp.wg.Done()
	defer close(cp.ch)

	for _, char := range cp.content {
		select {
		case <-ctx.Done():
			return
		case cp.ch <- char:
		}
	}
}

func (cp *ChannelProcessor) ReadFromChannel(ctx context.Context) {
	defer cp.wg.Done()
	defer cp.cancel()

	upperWriter := NewFileWriter(UpperFile)
	defer upperWriter.file.Close()

	lowerWriter := NewFileWriter(LowerFile)
	defer lowerWriter.file.Close()

	for char := range cp.ch {
		if upperWriter.lines >= MaxLines || lowerWriter.lines >= MaxLines {
			break
		}

		if unicode.IsUpper(char) {
			upperWriter.Write(char)
		} else if unicode.IsLower(char) {
			lowerWriter.Write(char)
		}
	}
	if upperWriter.lines >= MaxLines {
		fmt.Println("upper.txt reached 100 lines first!")
	} else {
		fmt.Println("lower.txt reached 100 lines first!")
	}
}

func (cp *ChannelProcessor) Start() {
	ctx, cancel := context.WithCancel(context.Background())
	cp.cancel = cancel
	cp.wg.Add(2)

	go cp.WriteToChannel(ctx)
	go cp.ReadFromChannel(ctx)

	cp.wg.Wait()
}


func main() {
	processor := NewChannelProcessor(sampleString)
	processor.Start()
	fmt.Println("Processing complete.")
}
