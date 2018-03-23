package common

import "testing"

func Test_DownLoadImage(t *testing.T) {
	img := "http://img.zcool.cn/community/018d4e554967920000019ae9df1533.jpg"

	DownLoadImage(img)
}
