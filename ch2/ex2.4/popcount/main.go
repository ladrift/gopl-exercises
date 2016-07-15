/*
Exercise 2.4:
Write a version of PopCount that counts bits by shifting its argument through
64 bit positions, testing the rightmost bit each time. Compare the performance
to the table-lookup version.
*/
package popcount

func PopCount(x uint64) int {
	count := 0
	for x > 0 {
		if x&1 == 1 {
			count++
		}
		x >>= 1
	}
	return count
}
