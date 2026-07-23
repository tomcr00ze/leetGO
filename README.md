# 🐹 leetGO - Mastering Go through LeetCode

[![Go CI](https://github.com/tomcr00ze/leetGO/actions/workflows/test.yml/badge.svg)](https://github.com/tomcr00ze/leetGO/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/tomcr00ze/leetGO)](https://goreportcard.com/report/github.com/tomcr00ze/leetGO)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

A structured, production-grade Go repository dedicated to solving LeetCode problems while mastering idiomatic Go design patterns, **Table-Driven Testing**, **Benchmarking**, and **Automated Scaffolding**.

---

## 🎯 Repository Principles

1. **Idiomatic Go**: Code structured following standard Go conventions (`go fmt`, `go vet`).
2. **Table-Driven Unit Tests**: Every solution is paired with comprehensive table-driven unit tests testing standard & edge cases.
3. **Performance Benchmarking**: Integrated Go benchmarks (`BenchmarkXxx(b *testing.B)`) to measure execution speed (ns/op) and memory allocations.
4. **Pattern-Based Organization**: Solutions are organized into algorithm and data structure categories (`arrays_hashing`, `two_pointers`, `binary_search`, `trees`, etc.) rather than a flat file dump.
5. **Zero-Friction Scaffolding**: Built-in CLI tool to instantaneously scaffold problem directories and test templates.

---

## 📁 Repository Structure

```text
leetGO/
├── .github/
│   └── workflows/
│       └── test.yml                  # GitHub Actions CI pipeline
├── arrays_hashing/                   # Topic category
│   ├── 0001-two-sum/
│   │   ├── two_sum.go                # Solution & complexity analysis
│   │   └── two_sum_test.go           # Table-driven tests & benchmarks
│   ├── 0066-plus-one/
│   └── 0217-contains-duplicate/
├── two_pointers/
├── strings/
├── stack/
├── linked_list/
├── binary_search/
├── math/
├── bit_manipulation/
├── dynamic_programming/
├── sliding_window/
├── trees/
├── scripts/
│   └── init.go                       # CLI automation script
├── go.mod                            # Go module definition
└── README.md                         # Problem tracking index
```

---

## 🚀 Automation CLI Tool (`init.go`)

Generate a new problem directory with solution and table-driven test templates effortlessly:

```bash
go run scripts/init.go <category> <problem_number> "<Problem Title>"
```

### Example:
```bash
go run scripts/init.go arrays_hashing 0242 "Valid Anagram"
```

---

## 🧪 Testing & Benchmarking

### Run All Unit Tests
```bash
go test -v ./...
```

### Run Benchmarks
```bash
go test -bench=. ./...
```

### Run Code Linter & Formatting
```bash
go fmt ./...
go vet ./...
```

---

## 📊 LeetCode Progress (25 Foundational Solutions)

| # | Problem Title | Pattern / Category | Difficulty | Time | Space | Solution & Tests |
|---|---|---|---|---|---|---|
| 0001 | [Two Sum](https://leetcode.com/problems/two-sum/) | `arrays_hashing` | Easy | $O(N)$ | $O(N)$ | [Code](arrays_hashing/0001-two-sum/two_sum.go) \| [Tests](arrays_hashing/0001-two-sum/two_sum_test.go) |
| 0009 | [Palindrome Number](https://leetcode.com/problems/palindrome-number/) | `math` | Easy | $O(\log_{10} N)$ | $O(1)$ | [Code](math/0009-palindrome-number/palindrome_number.go) \| [Tests](math/0009-palindrome-number/palindrome_number_test.go) |
| 0013 | [Roman to Integer](https://leetcode.com/problems/roman-to-integer/) | `strings` | Easy | $O(N)$ | $O(1)$ | [Code](strings/0013-roman-to-integer/roman_to_integer.go) \| [Tests](strings/0013-roman-to-integer/roman_to_integer_test.go) |
| 0014 | [Longest Common Prefix](https://leetcode.com/problems/longest-common-prefix/) | `strings` | Easy | $O(N \cdot M)$ | $O(1)$ | [Code](strings/0014-longest-common-prefix/longest_common_prefix.go) \| [Tests](strings/0014-longest-common-prefix/longest_common_prefix_test.go) |
| 0020 | [Valid Parentheses](https://leetcode.com/problems/valid-parentheses/) | `stack` | Easy | $O(N)$ | $O(N)$ | [Code](stack/0020-valid-parentheses/valid_parentheses.go) \| [Tests](stack/0020-valid-parentheses/valid_parentheses_test.go) |
| 0021 | [Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/) | `linked_list` | Easy | $O(N + M)$ | $O(1)$ | [Code](linked_list/0021-merge-two-sorted-lists/merge_two_sorted_lists.go) \| [Tests](linked_list/0021-merge-two-sorted-lists/merge_two_sorted_lists_test.go) |
| 0026 | [Remove Duplicates from Sorted Array](https://leetcode.com/problems/remove-duplicates-from-sorted-array/) | `two_pointers` | Easy | $O(N)$ | $O(1)$ | [Code](two_pointers/0026-remove-duplicates-from-sorted-array/remove_duplicates.go) \| [Tests](two_pointers/0026-remove-duplicates-from-sorted-array/remove_duplicates_test.go) |
| 0027 | [Remove Element](https://leetcode.com/problems/remove-element/) | `two_pointers` | Easy | $O(N)$ | $O(1)$ | [Code](two_pointers/0027-remove-element/remove_element.go) \| [Tests](two_pointers/0027-remove-element/remove_element_test.go) |
| 0028 | [Find First Occurrence in a String](https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/) | `strings` | Easy | $O(N \cdot M)$ | $O(1)$ | [Code](strings/0028-find-the-index-of-the-first-occurrence-in-a-string/index_first_occurrence.go) \| [Tests](strings/0028-find-the-index-of-the-first-occurrence-in-a-string/index_first_occurrence_test.go) |
| 0035 | [Search Insert Position](https://leetcode.com/problems/search-insert-position/) | `binary_search` | Easy | $O(\log N)$ | $O(1)$ | [Code](binary_search/0035-search-insert-position/search_insert.go) \| [Tests](binary_search/0035-search-insert-position/search_insert_test.go) |
| 0058 | [Length of Last Word](https://leetcode.com/problems/length-of-last-word/) | `strings` | Easy | $O(N)$ | $O(1)$ | [Code](strings/0058-length-of-last-word/length_of_last_word.go) \| [Tests](strings/0058-length-of-last-word/length_of_last_word_test.go) |
| 0066 | [Plus One](https://leetcode.com/problems/plus-one/) | `arrays_hashing` | Easy | $O(N)$ | $O(1)$ | [Code](arrays_hashing/0066-plus-one/plus_one.go) \| [Tests](arrays_hashing/0066-plus-one/plus_one_test.go) |
| 0067 | [Add Binary](https://leetcode.com/problems/add-binary/) | `bit_manipulation` | Easy | $O(\max(N, M))$ | $O(\max(N, M))$ | [Code](bit_manipulation/0067-add-binary/add_binary.go) \| [Tests](bit_manipulation/0067-add-binary/add_binary_test.go) |
| 0069 | [Sqrt(x)](https://leetcode.com/problems/sqrtx/) | `binary_search` | Easy | $O(\log x)$ | $O(1)$ | [Code](binary_search/0069-sqrtx/sqrtx.go) \| [Tests](binary_search/0069-sqrtx/sqrtx_test.go) |
| 0070 | [Climbing Stairs](https://leetcode.com/problems/climbing-stairs/) | `dynamic_programming` | Easy | $O(N)$ | $O(1)$ | [Code](dynamic_programming/0070-climbing-stairs/climbing_stairs.go) \| [Tests](dynamic_programming/0070-climbing-stairs/climbing_stairs_test.go) |
| 0083 | [Remove Duplicates from Sorted List](https://leetcode.com/problems/remove-duplicates-from-sorted-list/) | `linked_list` | Easy | $O(N)$ | $O(1)$ | [Code](linked_list/0083-remove-duplicates-from-sorted-list/delete_duplicates.go) \| [Tests](linked_list/0083-remove-duplicates-from-sorted-list/delete_duplicates_test.go) |
| 0088 | [Merge Sorted Array](https://leetcode.com/problems/merge-sorted-array/) | `two_pointers` | Easy | $O(M + N)$ | $O(1)$ | [Code](two_pointers/0088-merge-sorted-array/merge_sorted_array.go) \| [Tests](two_pointers/0088-merge-sorted-array/merge_sorted_array_test.go) |
| 0094 | [Binary Tree Inorder Traversal](https://leetcode.com/problems/binary-tree-inorder-traversal/) | `trees` | Easy | $O(N)$ | $O(N)$ | [Code](trees/0094-binary-tree-inorder-traversal/inorder_traversal.go) \| [Tests](trees/0094-binary-tree-inorder-traversal/inorder_traversal_test.go) |
| 0100 | [Same Tree](https://leetcode.com/problems/same-tree/) | `trees` | Easy | $O(N)$ | $O(H)$ | [Code](trees/0100-same-tree/same_tree.go) \| [Tests](trees/0100-same-tree/same_tree_test.go) |
| 0101 | [Symmetric Tree](https://leetcode.com/problems/symmetric-tree/) | `trees` | Easy | $O(N)$ | $O(H)$ | [Code](trees/0101-symmetric-tree/symmetric_tree.go) \| [Tests](trees/0101-symmetric-tree/symmetric_tree_test.go) |
| 0104 | [Maximum Depth of Binary Tree](https://leetcode.com/problems/maximum-depth-of-binary-tree/) | `trees` | Easy | $O(N)$ | $O(H)$ | [Code](trees/0104-maximum-depth-of-binary-tree/max_depth.go) \| [Tests](trees/0104-maximum-depth-of-binary-tree/max_depth_test.go) |
| 0121 | [Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/) | `sliding_window` | Easy | $O(N)$ | $O(1)$ | [Code](sliding_window/0121-best-time-to-buy-and-sell-stock/best_time_stock.go) \| [Tests](sliding_window/0121-best-time-to-buy-and-sell-stock/best_time_stock_test.go) |
| 0125 | [Valid Palindrome](https://leetcode.com/problems/valid-palindrome/) | `two_pointers` | Easy | $O(N)$ | $O(1)$ | [Code](two_pointers/0125-valid-palindrome/valid_palindrome.go) \| [Tests](two_pointers/0125-valid-palindrome/valid_palindrome_test.go) |
| 0136 | [Single Number](https://leetcode.com/problems/single-number/) | `bit_manipulation` | Easy | $O(N)$ | $O(1)$ | [Code](bit_manipulation/0136-single-number/single_number.go) \| [Tests](bit_manipulation/0136-single-number/single_number_test.go) |
| 0217 | [Contains Duplicate](https://leetcode.com/problems/contains-duplicate/) | `arrays_hashing` | Easy | $O(N)$ | $O(N)$ | [Code](arrays_hashing/0217-contains-duplicate/contains_duplicate.go) \| [Tests](arrays_hashing/0217-contains-duplicate/contains_duplicate_test.go) |

---

## 📜 License
This repository is licensed under the [MIT License](LICENSE).
