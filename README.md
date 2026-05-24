# LeetCode 75 — Go

Parallel companion to `leetcode-75-py`. One folder per problem, each its own Go package.

## Run

```pwsh
# Single problem
go test ./01-merge-strings-alternately/

# Everything
go test ./...
```

## Layout

- `NN-kebab-name/solution.go` — solution, LeetCode-style signature, lowercase entry function
- `NN-kebab-name/solution_test.go` — table-driven tests

## Progress

- [x] 01 Merge Strings Alternately (1768)
- [ ] 02 Greatest Common Divisor of Strings (1071)
- [ ] 03 Kids With the Greatest Number of Candies (1431)
- [ ] 04 Can Place Flowers (605)
- [ ] 05 Reverse Vowels of a String (345)
- [ ] ... fill in the rest from https://leetcode.com/studyplan/leetcode-75/

## Go idioms worth front-loading

- `strings.Builder` for string concat in loops (avoid `+=`)
- Slices: `s[i:]`, `append`, `len`/`cap`, share underlying array — copy when needed
- `container/heap` requires implementing `heap.Interface` (5 methods) — verbose but mechanical
- `sort.Slice(s, func(i,j int) bool { ... })` for custom sort
- Strings are immutable byte sequences; use `[]rune` for unicode-correct iteration
- Map zero-value reads return zero of value type; `v, ok := m[k]` to distinguish missing
