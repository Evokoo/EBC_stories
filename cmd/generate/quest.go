package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Config struct {
	Event string `json:"event"`
	CurrentDir string
	Quest string
}

type File struct {
	fileName string
	content  string
	open     bool
}


func main() {
	quest 		:= getQuest()
	rootDir 	:= findModuleRoot()
	config 		:= loadConfig(rootDir, quest)

	fileList := []File{
		{fileName: "example_I.txt", content: "", open: true},
		{fileName: "example_II.txt", content: "", open: false},
		{fileName: "example_III.txt", content: "", open: false},
		{fileName: "input_I.txt", content: "", open: true},
		{fileName: "input_II.txt", content: "", open: false},
		{fileName: "input_III.txt", content: "", open: false},
		{fileName: fmt.Sprintf("%s_suite_test.go", config.Quest), content: loadTemplate(filepath.Join(rootDir, "cmd/generate/testSuite.tmpl"), config), open: false},
		{fileName: fmt.Sprintf("%s_test.go", config.Quest), content: loadTemplate(filepath.Join(rootDir, "cmd/generate/test.tmpl"), config), open: true},
		{fileName: fmt.Sprintf("%s_logic.go", config.Quest), content: loadTemplate(filepath.Join(rootDir, "cmd/generate/logic.tmpl"), config), open: true},
		{fileName: fmt.Sprintf("%s.go", config.Quest), content: loadTemplate(filepath.Join(rootDir, "cmd/generate/solver.tmpl"), config), open: true},
	}

	createFolder(config.Quest)
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
func loadConfig(rootDir,quest string) Config {
	configPath := filepath.Join(rootDir, "cmd/generate/config.json")

	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Failed to read config.json: %v", err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		log.Fatalf("Failed to parse config.json: %v", err)
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}

	config.CurrentDir = filepath.Base(dir)
	config.Quest = quest
	
	return config
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
func loadTemplate(template string, config Config) string {
	data, err := os.ReadFile(template)
	if err != nil {
		log.Fatalf("Failed to read template: %v", err)
	}

	content := string(data)
	content = strings.ReplaceAll(content, "$EVENT$", config.Event)
	content = strings.ReplaceAll(content, "$STORY$", config.CurrentDir)
	content = strings.ReplaceAll(content, "$QUEST$", config.Quest)

	return content
}
func openFile(file string) {
	cmd := exec.Command("code", file)
	err := cmd.Start()
	if err != nil {
		log.Printf("Failed to open %s in VS Code: %v", file, err)
	}

}
func findModuleRoot() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	log.Fatal("Could not find go.mod in any parent directory")
	return ""
}

