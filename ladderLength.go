package main

// 46
func ladderLength(beginWord string, endWord string, wordList []string) int {
	// Convert wordList to a set for fast lookup
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}

	if !wordSet[endWord] {
		return 0 // endWord not in wordList
	}

	// Perform BFS
	queue := []string{beginWord}
	visited := make(map[string]bool)
	visited[beginWord] = true
	level := 1

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			currentWord := queue[0]
			queue = queue[1:]

			// Check if we reached the endWord
			if currentWord == endWord {
				return level
			}

			// Explore neighbors of currentWord
			for j := 0; j < len(currentWord); j++ {
				for k := 'a'; k <= 'z'; k++ {
					nextWord := currentWord[:j] + string(k) + currentWord[j+1:]
					if wordSet[nextWord] && !visited[nextWord] {
						visited[nextWord] = true
						queue = append(queue, nextWord)
					}
				}
			}
		}
		level++
	}

	return 0 // No transformation sequence found
}

// func main() {
//     beginWord := "hit"
//     endWord := "cog"
//     wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
//     fmt.Println(ladderLength(beginWord, endWord, wordList)) // Output: 5
// }
