/*
Exercise 2.5:
The expression x&(x-1) clears the rightmost non-zero bit of x. Write a version
of PopCount that counts bits by using this fact, and assess its performance.
*/
package popcount

func PopCount(x uint64) int {
	count := 0
	for x > 0 {
		x &= x - 1
		count++
	}
	return count
}
