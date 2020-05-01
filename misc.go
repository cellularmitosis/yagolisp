package main

// A slice operator which is forgiving of bounds violations.
func safeSlice(bytes []byte, start int, end int) []byte {
	if start < 0 {
		start = 0
	}
	if start >= len(bytes) {
		start = len(bytes) - 1
	}
	if end < start {
		end = start
	}
	if end >= len(bytes) {
		end = len(bytes) - 1
	}
	return bytes[start:end]
}
