package addons

import (
	"dzhgo/addons/customer_pro"
	"dzhgo/addons/dict"
	"dzhgo/addons/member"
	"dzhgo/addons/task"
)

func NewInit() {

	dict.NewInit()
	//space.NewInit()
	task.NewInit()
	member.NewInit()
	//crm.NewInit()
	//fileUpload.NewInit()
	customer_pro.NewInit()
}
