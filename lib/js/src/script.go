package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	log.Println("INFO: Start")
	clearDir()

	dir, err := os.ReadDir(".")
	if err != nil {
		log.Fatalln(err)
	}

	for _, entry := range dir {
		if entry.IsDir() {
			continue
		}

		fileName := entry.Name()

		if !strings.HasSuffix(fileName, ".js") {
			continue
		}

		content, err := os.ReadFile(fileName)
		if err != nil {
			log.Fatalln(err)
		}

		compile(fileName, content)
	}
}

func compile(fileName string, content []byte) {
	fName, fileName := functionName(fileName)
	log.Printf("INFO: Compiling %s", fName)

	fContent := "package lib_js\n"
	fContent += fmt.Sprintf("\nconst %s = `\n", fName)
	fContent += string(content)
	fContent += "`\n"

	filePath := filepath.Join("../", fmt.Sprintf("%s.go", fileName))
	if err := os.WriteFile(filePath, []byte(fContent), 0644); err != nil {
		log.Fatalln(err)
	}
}

func functionName(name string) (string, string) {
	name = strings.TrimSuffix(name, ".js")

	parts := strings.Split(name, "_")
	for i := range parts {
		str := parts[i]
		parts[i] = strings.ToUpper(str[0:1]) + str[1:]
	}

	return strings.Join(parts, ""), name
}

func clearDir() {
	dir, err := os.ReadDir("../")
	if err != nil {
		log.Fatalln(err)
	}

	for _, entry := range dir {
		if entry.IsDir() {
			continue
		}

		if !strings.HasSuffix(entry.Name(), ".go") {
			continue
		}

		if err := os.Remove(filepath.Join("../", entry.Name())); err != nil {
			log.Fatalln(err)
		}
	}
}
