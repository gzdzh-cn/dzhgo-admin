# member

## 扫码加好友拉群
<img src="./dzh/weixin.jpg" alt="Description of image" width="200" height="280">

## dzhgo后台界面
<img src="./dzh/dzhgo.png" alt="Description of image" width="800" height="360">

## 主程序dzhgo
* 后台项目地址：https://github.com/gzdzh-cn/dzhgo
* 前端项目地址：https://github.com/gzdzh-cn/dzhgo-admin-vue

 

## addons/main.go示例：
```shell
package addons

import (
	"dzhgo/addons/member"
)

func NewInit() {
	
	member.NewInit()

}
```

## 插件实现的功能
 

## 注意事项

 

## 资源打包命令

```bash

gf pack addons/member/resource addons/member/packed/packed.go -p addons/member/resource

gf gen dao -p=addons/member -t=addons_member_manage,addons_member_attr

gf gen service -s=addons/member/logic -d=addons/member/service

```


### 更新日志

v1.0.1
- 修改数据库字段类型
 
v1.0.0
- 完成对接前端登录接口


