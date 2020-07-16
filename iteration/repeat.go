/*
Module to repeat a given character a defined number of times.
Provides examples of the following:
 for
 += "the Add AND assignment operator"
 benchmark
 string library
*/
package iteration

func Repeat(character string, times int) (repeated string) {
	for i := 0; i < times; i++ {
		repeated += character
	}
	return repeated
}
