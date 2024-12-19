# task

任务管理模块,提供基于`corefun`c 的任务管理功能

## 资源打包命令

```bash
gf pack addons/task/resource addons/task/packed/packed.go -p addons/task/resource

gf gen dao -p=addons/task -t=addons_task_info,addons_task_log

gf gen service -s=addons/task/logic -d=addons/task/service
```
