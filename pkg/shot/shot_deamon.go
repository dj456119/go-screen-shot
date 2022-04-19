/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-04-19 22:01:16
 * @LastEditors: cm.d
 * @LastEditTime: 2022-04-19 22:18:38
 */
package shot

import (
	"time"
)

var IsStart bool = false

func DeamonShot() string {
	IsStart = true
	for IsStart {
		l := Shot()
		if l == "" {
			IsStart = false
			return l
		}
		time.Sleep(1 * time.Second)
	}
	return ""
}

func Goshot() string {
	go DeamonShot()
	return "ok"
}
