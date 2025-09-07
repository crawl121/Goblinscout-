package main

import (
	"flag"
	"fmt"
	"goblin_scout/internal/scanner"
	"goblin_scout/internal/web"
	"strconv"
	"strings"
)

func parseStatusCodes(input string) []int {
	parts := strings.Split(input, ",")
	codes := make([]int, 0)
	for _, p := range parts {
		if p == "" {
			continue
		}
		code, err := strconv.Atoi(strings.TrimSpace(p))
		if err == nil {
			codes = append(codes, code)
		}
	}
	return codes
}

func main() {
	url := flag.String("url", "", "Target URL for scanning")
	statusCodes := flag.String("status", "200,403", "Comma-separated list of HTTP status codes to include")
	outputFormat := flag.String("output", "json", "Output format: json or text")
	wordlist := flag.String("wordlist", "wordlist.txt", "Path to wordlist file")
	flag.Parse()

	if *url != "" {
		fmt.Println("[*] Running in CLI mode")
		status := parseStatusCodes(*statusCodes)
		scanner.StartScan(*url, status, *outputFormat, *wordlist)
	} else {
		fmt.Println("[*] Running in GUI mode at http://localhost:8080")
		web.StartServer()
	}
}
