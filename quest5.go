package piscine

import "strings"

func FirstRune(s string) rune{
	r := []rune(s)
	return r[0]
}

func NRune(s string, n int) rune{
	r := []rune(s)
	return r[n-1]
}

func LastRune(s string) rune{
	r := []rune(s)
	return r[len(r)-1]
}

func Index(s, toFind string) int {
	return strings.Index(s,toFind)
}

func Compare (a,b string) int{
	return strings.Compare(a,b)
}

func ToUpper(s string)  string{
	return strings.ToUpper(s)
}

func ToLower(s string)  string{
	return strings.ToLower(s)
}

func Capitalize(s string)  string{
	return strings.Title(ToLower(s))
}

func IsAlpha(str string) bool {
	alpha:= "abcdefghijklmnopqrstuvwxyz0123456789"
	for _, char := range str{
		if !strings.Contains(alpha, ToLower(string(char))){
			return false
		}
	}
	return true
}

func IsNumeric(str string) bool {
	num:= "0123456789"
	for _, char := range str{
		if !strings.Contains(num, string(char)){
			return false
		}
	}
	return true
}

func IsLower(str string) bool {
	alpha:= "abcdefghijklmnopqrstuvwxyz"
	for _, char := range str{
		if !strings.Contains(alpha, string(char)){
			return false
		}
	}
	return true
}

func IsUpper(str string) bool {
	alpha:= ToUpper("abcdefghijklmnopqrstuvwxyz")
	for _, char := range str{
		if !strings.Contains(alpha, string(char)){
			return false
		}
	}
	return true
}

func IsPrintable(str string) bool {
	return (!strings.Contains(str, string("\n")) && !strings.Contains(str, string("\b")) && !strings.Contains(str, string("\v")) && !strings.Contains(str, string("\t")) && !strings.Contains(str, string("\f")) && !strings.Contains(str, string("\a")))
}

func Concat(str1, str2 string) string {
	return str1 + str2
}
func BasicJoin(strs []string) string {
	aux:=""
	for _, str:= range strs{
		aux+=str
	}
	return aux
}

func Join(strs []string ,sep string) string {
	aux:=""
	for i, str:= range strs{
		aux+=str
		if i<len(strs)-1{
			aux+=sep
		}
	}
	return aux
}