package scanner

import (
	"fmt"
	"net"
	"strings"
)

func BruteForceDNS(domain string, subdomains []string) {
	fmt.Println("[*] Starting DNS brute-force on:", domain)
	for _, sub := range subdomains {
		subdomain := fmt.Sprintf("%s.%s", sub, domain)
		ips, err := net.LookupHost(subdomain)
		if err == nil {
			fmt.Printf("[+] Found subdomain: %s -> %s\n", subdomain, strings.Join(ips, ", "))
		}
	}
}
