package output

import (
	"encoding/json"
	"fmt"
	"os"
)

// WriteJSON saves results to a JSON file
func WriteJSON(filename string, data interface{}) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("[-] Could not create file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(data)
	if err != nil {
		fmt.Println("[-] Could not write JSON:", err)
	}
}

// WriteText saves results to a text file AND also prints to terminal
func WriteText(filename string, data []map[string]interface{}) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("[-] Could not create file:", err)
		return
	}
	defer file.Close()

	for _, item := range data {
		line := fmt.Sprintf("%s (%d)\n", item["url"], item["status"])
		fmt.Print(line) // ✅ also show in terminal
		_, _ = file.WriteString(line)
	}
}
