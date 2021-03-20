/*
 Author: Kernel.Huang
 Mail: kernelman79@gmail.com
 File: TimeDuration
 Date: 3/20/21 12:54 PM
*/
package utils

import "time"

type Times struct{}

var Time = new(Times)

func (t *Times) IntToSecond(second int) time.Duration {
	timeSecond := time.Duration(second)
	return timeSecond * time.Second
}
