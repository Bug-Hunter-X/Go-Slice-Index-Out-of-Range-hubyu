# Go Slice Index Out of Range Bug

This repository demonstrates a common error in Go programs: accessing a slice element with an index that's out of bounds. This results in a runtime panic, which can be difficult to debug if not handled carefully.

## Bug Description

The bug occurs when attempting to access an element in a slice using an index that is greater than or equal to the length of the slice. Go slices have a zero-based index, which means the first element is at index 0 and the last element is at index len(slice)-1.  Accessing any index beyond this range leads to a runtime panic with the message `runtime error: index out of range`.

## Bug Reproduction

1. Clone the repository.
2. Run the `bug.go` file.
3. Observe the runtime panic.

## Bug Solution

The solution is to carefully check array and slice boundaries before accessing elements.  Always ensure your index is within the valid range (0 to len(slice)-1) to avoid this error.  The provided `bugSolution.go` demonstrates safe ways of accessing slice elements by checking the length and adding error handling.