package main

import (
	"bufio"
	"fmt"
	"github.com/stoewer/go-docx"
	"os"
)

func main() {
	// Get the filename from the user
	fmt.Print("Enter the name of the .docx file: ")
	reader := bufio.NewReader(os.Stdin)
	filename, _ := reader.ReadString('\n')
	filename = filename[:len(filename)-1]

	// Open the docx file
	file, err := docx.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	// Create a new file for the markdown output
	mdFile, err := os.Create(filename + ".md")
	if err != nil {
		fmt.Println("Error creating file:", err)
		os.Exit(1)
	}
	defer mdFile.Close()

	// Loop through all paragraphs in the docx file
	for _, p := range file.Paragraphs() {
		// Write the text of each paragraph to the markdown file
		_, err = mdFile.WriteString(p.Text() + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			os.Exit(1)
		}
	}

	// Loop through all images in the docx file
	for i, img := range file.Images() {
		// Save each image to a separate file
		err = img.Save("image" + string(i) + ".png")
		if err != nil {
			fmt.Println("Error saving image:", err)
			os.Exit(1)
		}
		// Write the markdown code for each image to the output file
		_, err = mdFile.WriteString("![Image" + string(i) + "](" + "image" + string(i) + ".png)\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			os.Exit(1)
		}
	}

	fmt.Println("Conversion complete.")
}
