package main

import (
	"fmt"
	"strings"
)

func ChangeOutput(data []string) map[string][]string {
	result := make(map[string][]string)

	for _, item := range data {
		parts := strings.Split(item, "-")

		if len(parts) != 4 {
			continue
		}

		header := parts[0]
		value := parts[3]

		if _, exists := result[header]; !exists {
			result[header] = make([]string, 0)
		}

		if parts[2] == "first" {
			result[header] = append(result[header], value)
		} else if parts[2] == "last" {

			if len(result[header]) > 0 {
				lastIndex := len(result[header]) - 1
				result[header][lastIndex] += " " + value
			}
		}
	}

	return result
}

func main() {
	data := []string{"account-0-first-John", "account-0-last-Doe", "account-1-first-Jane", "account-1-last-Doe", "address-0-first-Jaksel", "address-0-last-Jakarta", "address-1-first-Bandung", "address-1-last-Jabar", "phone-0-first-081234567890", "phone-1-first-081234567891"}
	res := ChangeOutput(data)

	fmt.Println(res)
}
