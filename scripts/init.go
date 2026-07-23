package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run scripts/init.go <category> <number> \"<Problem Name>\"")
		fmt.Println("Example: go run scripts/init.go arrays_hashing 0001 \"Two Sum\"")
		return
	}

	category := strings.ToLower(os.Args[1])
	rawNumber := os.Args[2]
	problemName := os.Args[3]

	// Format 4-digit problem number
	num := fmt.Sprintf("%04s", rawNumber)

	// Create slug for folder (e.g., "Two Sum" -> "two-sum")
	reg := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	slug := strings.Trim(strings.ToLower(reg.ReplaceAllString(problemName, "-")), "-")

	// Create safe package name (e.g., "Two Sum" -> "twosum")
	pkgName := strings.ReplaceAll(strings.ToLower(problemName), " ", "")
	pkgName = reg.ReplaceAllString(pkgName, "")

	// Create safe snake_case filename (e.g., "two_sum")
	fileBase := strings.ReplaceAll(slug, "-", "_")

	// Directory path: <category>/<number>-<slug>
	dirName := filepath.Join(category, fmt.Sprintf("%s-%s", num, slug))

	if err := os.MkdirAll(dirName, 0755); err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}

	goFile := filepath.Join(dirName, fmt.Sprintf("%s.go", fileBase))
	testFile := filepath.Join(dirName, fmt.Sprintf("%s_test.go", fileBase))

	// Template for Go Solution File
	goTemplate := fmt.Sprintf(`package %s

/*
Problem: %s
Link: https://leetcode.com/problems/%s/
Difficulty: Easy

Time Complexity: O(N)
Space Complexity: O(1)
*/

// Solution implementation
`, pkgName, problemName, slug)

	// Template for Go Table-Driven Test File
	testTemplate := fmt.Sprintf(`package %s

import (
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected interface{}
	}{
		{
			name:     "Example 1",
			input:    nil,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call solution and assert expected result
		})
	}
}

func BenchmarkSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Benchmark execution
	}
}
`, pkgName)

	if err := os.WriteFile(goFile, []byte(goTemplate), 0644); err != nil {
		fmt.Printf("Error writing solution file: %v\n", err)
		return
	}

	if err := os.WriteFile(testFile, []byte(testTemplate), 0644); err != nil {
		fmt.Printf("Error writing test file: %v\n", err)
		return
	}

	fmt.Printf("✅ Successfully scaffolded problem [%s - %s] in directory:\n   %s\n", num, problemName, dirName)
}
