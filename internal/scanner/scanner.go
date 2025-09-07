package scanner

import (
	"fmt"
	"goblin_scout/internal/output"
	"net/http"
	"os"
	"strings"
	"sync"
)

func StartScan(target string, statusFilter []int, outputFormat, wordlistPath string) {
	fmt.Println("[+] Starting directory brute-force on:", target)

	// Read wordlist
	wordlist, err := os.ReadFile(wordlistPath)
	if err != nil {
		fmt.Println("[-] Error reading wordlist:", err)
		return
	}
	words := strings.Split(string(wordlist), "\n")
	fmt.Printf("[*] Loaded %d entries from wordlist\n", len(words))

	var wg sync.WaitGroup
	results := make([]map[string]interface{}, 0)
	mutex := sync.Mutex{}

	statusMap := make(map[int]bool)
	for _, code := range statusFilter {
		statusMap[code] = true
	}

	for _, word := range words {
		word = strings.TrimSpace(word)
		if word == "" {
			continue
		}

		wg.Add(1)
		go func(path string) {
			defer wg.Done()
			url := fmt.Sprintf("%s/%s", strings.TrimSuffix(target, "/"), path)
			resp, err := http.Get(url)
			if err != nil {
				return
			}
			defer resp.Body.Close()

			if statusMap[resp.StatusCode] {
				fmt.Printf("[+] Found: %s (%d)\n", url, resp.StatusCode)

				mutex.Lock()
				results = append(results, map[string]interface{}{
					"url":    url,
					"status": resp.StatusCode,
				})
				mutex.Unlock()
			}
		}(word)
	}
	wg.Wait()

	// Output
	if outputFormat == "json" {
		output.WriteJSON("output.json", results)
		fmt.Println("[+] Results saved to output.json")
	} else {
		output.WriteText("output.txt", results)
		fmt.Println("[+] Results saved to output.txt")
	}
}
