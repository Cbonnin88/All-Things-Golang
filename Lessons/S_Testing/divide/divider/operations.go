// Package divider ask to begin the dividend equation
package divider

// Divide divides an unlimited number of values of type int
func Divide(n ...int) int {
	result := 1
	for _, v := range n {
		result = v / result
	}
	return result
}
