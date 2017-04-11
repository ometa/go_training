package integer_translator

/*
create an integer translation package
package will convert an int to a string
package uses an arbitrary conversion function
the returned result should contain the input
the returned result should contain the conversion type
provide fizzbuzz and roman numeral translators
*/

type IntegerTranslator interface {
	Translate(i int) (output string, input int, conversionType string, err error)
}
