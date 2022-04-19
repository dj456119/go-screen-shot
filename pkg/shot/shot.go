/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-04-19 14:36:45
 * @LastEditors: cm.d
 * @LastEditTime: 2022-04-19 23:33:28
 */
package shot

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"time"

	"github.com/kbinani/screenshot"
)

// save *image.RGBA to filePath with PNG format.
func save(img *image.RGBA, filePath string) {
	file, err := os.Create(filePath + ".tmp")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	png.Encode(file, img)
	os.Rename(filePath+".tmp", filePath)
}

func Shot() string {
	//获取所有活动屏幕
	n := screenshot.NumActiveDisplays()
	//全屏截取所有活动屏幕
	if n > 0 {
		for i := 0; i < n; i++ {
			img, err := screenshot.CaptureDisplay(i)
			if err != nil {
				panic(err)
			}
			save(img, fmt.Sprintf("screens/%d-screen.png", i))
		}
		return fmt.Sprintf("screen-%s.png", time.Now().Format("2006-01-02 15-04-05"))
	}
	return ""
}
