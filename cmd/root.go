package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "marsho",
	Short: "A CLI File Manager written in Go",
	Long:  `A fast and reliable CLI File Manager written in Go using the Cobra library`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Marsho FM")
		folder, err := os.Getwd()
		if err != nil {
			fmt.Printf("Error getting current directory: %v\n", err)
			return
		}
		showFiles(folder)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func showFiles(dir string) {
	contents, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	if len(contents) == 0 {
		fmt.Println("No files found.")
		return
	}

	// Convert directory entries to a slice of strings
	var fileNames []string

	// Add the option to go up one directory
	fileNames = append(fileNames, "..")

	for _, entry := range contents {
		if entry.IsDir() {
			fileNames = append(fileNames, entry.Name()+"/")
		} else {
			fileNames = append(fileNames, entry.Name())
		}
	}

	// Define the template for the prompt
	template := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   " \u21D2 {{ . | cyan | bold | bgGreen }}",
		Inactive: "  {{ . | faint }}",
		Selected: "✔ {{ . | red | bold | bgCyan }}",
	}

	// Create the prompt
	prompt := promptui.Select{
		Label:     "Select a file or directory",
		Items:     fileNames,
		Templates: template,
		Size:      100,
	}

	// Run the prompt and capture the selected index and value
	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed: %v\n", err)
		return
	}

	// Handle the selected option
	if result == ".." {
		// Navigate up one directory
		parentDir := filepath.Dir(dir)
		showFiles(parentDir)
	} else {
		// Navigate into the selected directory or handle the selected file
		selectedPath := filepath.Join(dir, result)
		fileInfo, err := os.Stat(selectedPath)
		if err != nil {
			fmt.Printf("Error accessing the selected path: %v\n", err)
			return
		}

		if fileInfo.IsDir() {
			// If it's a directory, show its contents
			showFiles(selectedPath)
		} else {
			// Display file details
			fmt.Printf("You selected file: %s\n", fileInfo.Name())
			fmt.Printf("File size: %d bytes\n", fileInfo.Size())
			fmt.Printf("Modification Time: %s\n", fileInfo.ModTime().Format(time.DateTime))

			// Open the file based on the OS
			err := openFile(selectedPath)
			if err != nil {
				fmt.Printf("Error opening file: %v\n", err)
			}
		}
	}
}

func openFile(filePath string) error {
	// Get the OS type
	switch runtime.GOOS {
	case "linux":
		// Linux command to open a file
		return exec.Command("xdg-open", filePath).Start()
	case "windows":
		// Windows command to open a file
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", filePath).Start()
	case "darwin":
		// macOS command to open a file
		return exec.Command("open", filePath).Start()
	default:
		return fmt.Errorf("unsupported platform")
	}
}
