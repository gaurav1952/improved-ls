package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"path/filepath"
	"unicode"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
)

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)

	}
}

func isBinary(filePath string) bool {
	file, err := os.Open(filePath)
	if err != nil {
		return false 
	}
	defer file.Close()
	extension := filepath.Ext(filePath)
	if extension == ".mod" || extension == ".txt" {
		return false
	}

	reader := bufio.NewReader(file)
	buffer := make([]byte, 512) 
	_, err = reader.Read(buffer)
	if err != nil {
		return false
	}

	
	for _, b := range buffer {
		if b > 0x7F || (!unicode.IsPrint(rune(b)) && !unicode.IsSpace(rune(b))) {
			return true
		}
	}
	return false
}
func main() {
	//colors
	reset := "\033[0m"
	boldYellow := "\033[1;33m"
	underlineBoldGreen := "\033[1;4;32m"


	args := os.Args
	dirPath := "."
	if len(args) > 1 {
		dirPath = args[1]
	}

	info, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		log.Fatalf("Error: The path '%s' does not exist.\n", dirPath)
	} else if !info.IsDir() {
		log.Fatalf("Error: The path '%s' is not a directory.\n", dirPath)
	}

	// color design
	// helpStyle_underline := color.New(color.FgHiCyan).SprintFunc()
	// helpStyle_working_command := color.New(color.FgWhite).SprintFunc()
	// Open the directory
	files, err := os.ReadDir(dirPath)
	checkerr(err)

	headerStyle := color.New(color.FgWhite, color.Bold).SprintFunc()

	file_table := table.NewWriter()
	file_table.SetOutputMirror(os.Stdout)
	file_table.AppendHeader(table.Row{
		headerStyle("#"),
		headerStyle("Files/Folder"),
		headerStyle("Type"),
		// headerStyle("Length"),
		// headerStyle("Mode"),
		// headerStyle("LastWriteTime"),
	})

	for i, file := range files {
		fileName := file.Name()
		is_dir := file.IsDir()


		fileTypeIcon := "📄";
		if is_dir{
			fileTypeIcon = "📁";
		}else {
			extension := filepath.Ext(fileName)
			switch extension {
			case ".zip", ".tar", ".gz", ".rar", ".7z":
				fileTypeIcon = "📦" // Compressed file
			case ".py":
				fileTypeIcon = "🐍" // Python
			case ".go":
				fileTypeIcon = "🐹" // Go 
			case ".js":
				fileTypeIcon = "✨" // JavaScript 
			case ".ts":
				fileTypeIcon = "📘" // TypeScript 
			case ".java":
				fileTypeIcon = "☕" // Java
			case ".rb":
				fileTypeIcon = "💎" // Ruby 
			case ".html", ".htm":
				fileTypeIcon = "🌐" // HTML 
			case ".css":
				fileTypeIcon = "🎀" // CSS 
			case ".rs":
				fileTypeIcon = "🦀" // Rust 
			case ".c", ".h":
				fileTypeIcon = "🌊" // C 
			case ".cpp", ".hpp":
				fileTypeIcon = "➕➕" // C++ 
			case ".dart":
				fileTypeIcon = "🎯" // Dart 
			case ".sh":
				fileTypeIcon = "🖥️" // Shell script
			case ".txt":
				fileTypeIcon = "📜"
			case ".php":
				fileTypeIcon = "🐘"
			default:
				// Check if it's a binary file
				if isBinary(filepath.Join(dirPath, fileName)) {
					fileTypeIcon = "💾"
				}
				}
			
		}
		// size := fileInfo.Size()
		// permission := fileInfo.Mode()
		// last_modifired := fileInfo.ModTime().Format("2006-01-02 15:04:05")

		file_table.AppendRow(table.Row{
			boldYellow + strconv.Itoa(i) + reset,
			underlineBoldGreen + fileName + reset,
			fileTypeIcon,
			// underlineMagenta + strconv.Itoa(int(size)) + " bytes" + reset,
			// boldMagenta + permission.String() + reset,
			// last_modifired,
		})

	}
	file_table.SetStyle(table.StyleLight)
	file_table.Render()

}
