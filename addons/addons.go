package addons

import (
	"dzhgo/addons/crm"
	"dzhgo/addons/customer_pro"
	"dzhgo/addons/dict"
	"dzhgo/addons/member"
	"dzhgo/addons/space"
	"dzhgo/addons/task"
	"dzhgo/addons/watermark_camera"
)

func NewInit() {

	dict.NewInit()
	space.NewInit()
	task.NewInit()
	member.NewInit()
	crm.NewInit()
	//file_upload.NewInit()
	customer_pro.NewInit()
	watermark_camera.NewInit()
}
