/*
 Author: Kernel.Huang
 Mail: kernelman79@gmail.com
 File: ArrayService
 Date: 3/19/21 2:39 PM
*/
package utils

import "strings"

type Arrays struct{}

var Array = new(Arrays)

// 字符串数组转字符串
func (a *Arrays) strArrayToString(arr []string, need string) string {
	return strings.Join(arr, need)
}
