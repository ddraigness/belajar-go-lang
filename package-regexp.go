package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex = regexp.MustCompile("a([a-z])q([a-z])")

	fmt.Println(regex.MatchString("alqo"))
	fmt.Println(regex.MatchString("alqi"))
	fmt.Println(regex.MatchString("alq0"))

	fmt.Println(regex.FindAllString("alqo alqi alqy alqe alq0 alqc 4lqo", 10))
}

/**
	package regexp :
1. package regexp adalah utilitas di go-lang untuk melakukan pencarian regular expression
2. regular expression di go-lang menggunakan library C yang dibuat Google bernama RE2
3. https://github.com/google/re2/wiki/Syntax
4. https://golang.org/pkg/regexp/
*/

/**
	beberapa function di package regexp
function : 							| kegunaan :
1. regexp.MustCompile(string)			Membuat regexp
2. Regexp.MatchString(string) bool		Mengecek apakah Regexp match dengan string
3. Regexp.FindAllString(string, max)	Mencari string yang match dengan maximum jumlah hasil
*/
