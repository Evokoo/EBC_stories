package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type File struct {
	fileName string
	content  string
	open     bool
}

func main() {
	if !rootCheck() {
		log.Fatal("Run from root")
	}

	quest := getQuest()
	fileList := []File{
		{fileName: "example_I.txt", content: "", open: true},
		{fileName: "example_II.txt", content: "", open: false},
		{fileName: "example_III.txt", content: "", open: false},
		{fileName: "input_I.txt", content: "", open: true},
		{fileName: "input_II.txt", content: "", open: false},
		{fileName: "input_III.txt", content: "", open: false},
		{fileName: fmt.Sprintf("%s_suite_test.go", quest), content: loadTemplate("cmd/generate/testSuite.tmpl", quest), open: false},
		{fileName: fmt.Sprintf("%s_test.go", quest), content: loadTemplate("cmd/generate/test.tmpl", quest), open: true},
		{fileName: fmt.Sprintf("%s_logic.go", quest), content: loadTemplate("cmd/generate/logic.tmpl", quest), open: true},
		{fileName: fmt.Sprintf("%s.go", quest), content: loadTemplate("cmd/generate/solver.tmpl", quest), open: true},
	}

	createFolder(quest)
	createFiles(fileList)
}
func getQuest() string {
	quest := flag.String("quest", "", "Which quest are you creating?")
	flag.Parse()

	if *quest == "" {
		log.Fatal("error: -quest flag is required")
	}
	return *quest
}
func createFolder(quest string) {
	//Make the folder
	err := os.Mkdir(quest, 0755)
	if err != nil {
		log.Fatalf("Failed to create quest folder: %v", err)
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	} else {
		fmt.Printf("%s folder created at %s\n", quest, dir)
	}

	//Move into folder
	err = os.Chdir(quest)
	if err != nil {
		log.Fatalf("Failed to change directory: %v", err)
	}

	//Create inputs folder
	err = os.Mkdir("inputs", 0755)
	if err != nil {
		log.Fatalf("Failed to create inputs folder: %v", err)
	}

	dir, err = os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	} else {
		fmt.Printf("%s folder created at %s\n", "inputs", dir)
	}

}
func createFiles(files []File) {
	for _, context := range files {
		filename := context.fileName

		if strings.HasSuffix(filename, ".txt") {
			filename = filepath.Join("inputs", filename)
		}

		//Create file
		file, err := os.Create(filename)
		if err != nil {
			log.Fatalf("Failed to create file: %v", err)
		}
		//Write contents
		_, err = file.WriteString(context.content)
		if err != nil {
			log.Fatalf("Failed to write to file: %v", err)
		}
		fmt.Println(filename, "created")

		if context.open {
			openFile(filename)
		}

		file.Close()
	}
}
func loadTemplate(template, quest string) string {
	data, err := os.ReadFile(template)
	if err != nil {
		log.Fatalf("Failed to read template: %v", err)
	}
	return strings.ReplaceAll(string(data), "$", quest)
}
func rootCheck() bool {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}

	if files, err := os.ReadDir(dir); err != nil {
		log.Fatalf("Failed to read filed in %s: %v", dir, err)
	} else {
		for _, file := range files {
			if file.Name() == "go.mod" {
				return true
			}
		}
	}
	return false
}
func openFile(file string) {
	cmd := exec.Command("code", file)
	err := cmd.Start()
	if err != nil {
		log.Printf("Failed to open %s in VS Code: %v", file, err)
	}

}
