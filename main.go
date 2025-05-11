package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	var userInput int
	var Url, startTime, endTime string

	// Create a 'downloads' directory if it doesn't exist
	err := os.MkdirAll("downloads", os.ModePerm)
	if err != nil {
		fmt.Println("Error creating downloads directory:", err)
		return
	}

	inputFilePath := "Input1.webm"
	outputFilePath := "downloads/output.webm"
	outputAudio := "downloads/output.opus"

	// Remove old files from previous runs if they exist
	err = RemoveFileIfExists(outputFilePath)
	if err != nil {
		fmt.Println("Error removing output file:", err)
		return
	}

	err = RemoveFileIfExists(outputAudio)
	if err != nil {
		fmt.Println("Error removing input file:", err)
		return
	}

	fmt.Println("Enter a number \n 1. Video Download Full \n 2. Download Audio \n 3. Trimmed Video \n 4. Audio Trim.")
	fmt.Scanln(&userInput)

	switch userInput {
	case 1:
		// Video Download Full
		fmt.Println("Enter the URL to download:")
		fmt.Scanln(&Url)
		cmd := exec.Command("yt-dlp", "-o", "downloads/%(title)s.%(ext)s", Url)
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Video downloaded and saved to downloads folder.")
	case 2:
		// Download Audio
		fmt.Println("Enter the URL to download:")
		fmt.Scanln(&Url)
		cmd := exec.Command("yt-dlp", "-x", "-o", "downloads/%(title)s.%(ext)s", Url)
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Audio saved to downloads folder.")
	case 3:
		// Trimmed Video
		fmt.Println("Enter the URL to download:")
		fmt.Scanln(&Url)
		fmt.Println("Enter Start:")
		fmt.Scanln(&startTime)

		fmt.Println("Enter End:")
		fmt.Scanln(&endTime)
		cmd := exec.Command("yt-dlp", "-o", "Input1", Url)
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Generate a timestamped filename for trimmed video
		trimmedFilePath := fmt.Sprintf("downloads/trimmed_%s.webm", time.Now().Format("20060102_150405"))
		cmd = exec.Command("ffmpeg", "-i", "Input1.webm", "-ss", startTime, "-to", endTime, "-c:v", "copy", "-c:a", "copy", trimmedFilePath)
		err = cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Trimmed video saved as:", trimmedFilePath)
	case 4:
		// Audio Trim
		fmt.Println("Enter the URL to download:")
		fmt.Scanln(&Url)
		fmt.Println("Enter Start:")
		fmt.Scanln(&startTime)
		fmt.Println("Enter End:")
		fmt.Scanln(&endTime)
		cmd := exec.Command("yt-dlp", "-x", "-o", "downloads/FullAudio.%(ext)s", Url)
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		cmd = exec.Command("ffmpeg", "-ss", startTime, "-to", endTime, "-i", "downloads/FullAudio.opus", "downloads/output.opus")
		err = cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Audio trimmed successfully.")
	default:
		fmt.Println("Invalid input. Please enter a number between 1 and 4.")
	}

	// Clean up input files after processing
	err = RemoveFileIfExists(inputFilePath)
	if err != nil {
		fmt.Println("Error removing input file:", err)
		return
	}
	err = RemoveFileIfExists("downloads/FullAudio.opus")
	if err != nil {
		fmt.Println("Error removing input file:", err)
		return
	}
}

func RemoveFileIfExists(filePath string) error {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return err
	}
	err = os.Remove(filePath)
	if err != nil {
		return err
	}
	fmt.Println("File", filePath, "removed successfully!")
	return nil
}
