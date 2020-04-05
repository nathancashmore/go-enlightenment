/*
Repeat( character ) :
Provides examples of the following:
 for
 +=
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
