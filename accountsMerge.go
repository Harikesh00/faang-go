package main

import (
	"sort"
)

// 37
func accountsMerge(accounts [][]string) [][]string {
	graph := make(map[string][]string)
	emailToName := make(map[string]string)

	// Construct the graph and map each email to its corresponding name
	for _, account := range accounts {
		name := account[0]
		for _, email := range account[1:] {
			graph[email] = append(graph[email], account[1])
			graph[account[1]] = append(graph[account[1]], email)
			emailToName[email] = name
		}
	}

	visited := make(map[string]bool)
	var dfs func(email string, emails []string)
	var mergedAccounts [][]string

	// Perform DFS to group connected emails (belonging to the same account)
	dfs = func(email string, emails []string) {
		visited[email] = true
		emails = append(emails, email)
		for _, neighbor := range graph[email] {
			if !visited[neighbor] {
				dfs(neighbor, emails)
			}
		}
	}

	// Traverse the graph and merge accounts
	for email := range emailToName {
		if !visited[email] {
			var emails []string
			dfs(email, emails)
			sort.Strings(emails)
			name := emailToName[email]
			mergedAccount := append([]string{name}, emails...)
			mergedAccounts = append(mergedAccounts, mergedAccount)
		}
	}

	return mergedAccounts
}

// func main() {
//     accounts := [][]string{
//         {"John", "johnsmith@mail.com", "john00@mail.com"},
//         {"John", "johnnybravo@mail.com"},
//         {"John", "johnsmith@mail.com", "john_newyork@mail.com"},
//         {"Mary", "mary@mail.com"},
//     }
//     fmt.Println(accountsMerge(accounts))
// }
