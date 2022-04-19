/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-04-19 22:37:23
 * @LastEditors: cm.d
 * @LastEditTime: 2022-04-20 00:06:40
 */
package v

import (
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
)

func ChangeContent(cavas fyne.Canvas) {
	for {
		imageC := canvas.NewImageFromFile("screens/0-screen.png")
		cavas.SetContent(imageC)
		time.Sleep(1 * time.Second)
	}
}
