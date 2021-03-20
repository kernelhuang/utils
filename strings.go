/*
 Author: Kernel.Huang
 Mail: kernelman79@gmail.com
 File: StringService
 Date: 3/19/21 2:23 PM
*/
package utils

import "strings"

type Strings struct{}

var String = new(Strings)

/**
 * 分割字符串
 */
func (s *Strings) Substr(str string, pos, length int) string {
	runes := []rune(str)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

/**
 * 分割文件路径以获取文件夹路径
 */
func (s *Strings) SplitDir(path string) string {
	return s.Substr(path, 0, strings.LastIndex(path, "/"))
}

/**
 * 分割文件路径以获取文件名
 */
func (s *Strings) SplitFileName(path string) string {
	fileArray := strings.Split(path, "/")
	length := len(fileArray)
	return fileArray[length-1]
}

/**
 * 分割文件名以获取文件后缀名
 */
func (s *Strings) SplitSuffixName(filename string) string {
	fileArray := strings.Split(filename, ".")
	length := len(fileArray)
	return fileArray[length-1]
}

/**
 * 分割字符串以获取字符串后缀
 */
func (s *Strings) SplitGetSuffix(str string) string {
	fileArray := strings.Split(str, "/")
	length := len(fileArray)
	return fileArray[length-1]
}

/**
 * 分割字符串以获取字符串前缀
 */
func (s *Strings) SplitGetPrefix(str string) []string {
	fileArray := strings.Split(str, "/")
	length := len(fileArray)
	return append(fileArray[:length-1], fileArray[(length-1)+1:]...)
}

/**
 * 反转字符串
 */
func (s *Strings) ReverseRune(runes []rune) []rune {
	for front, back := 0, len(runes)-1; front < back; front, back = front+1, back-1 {
		runes[front], runes[back] = runes[back], runes[front]
	}
	return runes
}
