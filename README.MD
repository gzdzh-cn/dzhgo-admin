# dzhgo

## 扫码加好友拉群
<img src="./dzh/weixin.jpg" alt="Description of image" width="200" height="280">


## dzhgo后台界面
<img src="./dzh/dzhgo.png" alt="Description of image" width="800" height="360">

## 主程序dzhgo
### github
* 后台项目地址：https://github.com/gzdzh-cn/dzhgo-admin
* 前端项目地址：https://github.com/gzdzh-cn/dzhgo-admin-vue


### gitee
* 后台项目地址：https://gitee.com/gzdzh_cn/dzhgo-admin
* 前端项目地址：https://gitee.com/gzdzh_cn/dzhgo-admin-vue.git





## 使用

1. 安装依赖

```bash
go mod tidy
```

2. 运行

```bash
gf run main.go
```

3. 打包静态文件
```bash
gf pack public,template,resource packed/data.go
```

## 集成命令

Makefile 中集成了一些常用的命令，可以直接使用 `make` 命令执行。

```bash
clean                          清理项目,用于删除开发容器及存储卷,需在本地开发环境执行
help                           查看帮助
init                           初始化项目,用于在开发容器生成后配置一些常用镜像,如: golang, nodejs, docker
mysql-backup                   备份mysql
mysql-down                     停止mysql
mysql-up                       启动mysql
redis-down                     停止redis
redis-up                       启动redis
setmirror                      设置国内镜像源,用于在开发容器生成后配置国内镜像源
```




## 计划更新
* -如果菜单在某个权限组里面增加增加一个方法权限，需要手动去角色再打开编辑保存一次才可以，比如菜单目录：通用，里面有权限page和list，添加多一个add后，需要到角色编辑保存一次（升级）
* [插件市场分为本地插件和云插件]（升级）
* [用户表和角色表的数据权限逻辑，部门主管看团队，成员看自己]（升级）
- [修复安装插件后，超管要更新一次权限缓存]
- [dict改为不是插件方式]
- [修复未安装的插件，菜单列表为不可见]


## 更新日志
### v1.2.3 - 20250703
- 升级 dzhcore 到新版v1.2.7

### v1.2.2 - 20250703
- 升级 dzhcore 到新版v1.2.6

### v1.2.1 - 20250701
- 升级 dzhcore 到新版，重构包加载逻辑

### v1.2.0 - 20250622
- 重构依赖包加载顺序

### v1.1.14 - 20250505
- 修改获取 person 的接口

### v1.1.13 - 20250401
- 修改为超管也可以禁止登录
- 修改为超管登录 token 也存储到 redis，会自动过期

### v1.1.12 - 2024-10-15
- 修复会员列表编辑问题，修复会员禁止登录

### v1.1.12 - 2024-10-15
- dict接口改为comm权限

### v1.1.11 -日期：2024-09-10
- 优化运行日志 

### v1.1.10 -日期：2024-09-10
- 增加运行日志记录每个接口情况

### v1.1.9 -日期：2024-09-08
- 更新中间件日志写入
- 更改config.yaml的中间件日志写入忽略路径参数
- 更改依赖顺序
- 更新初始化数据
- 修改config里面的Config为Cfg
- 更改获取setting数据由原来NewConfig为新建一个baseConfig
- dzhcore升级到1.1.7

### v1.1.8 --
- dzhcore升级到1.1.6


### v1.1.7
- 修改首页


### v1.1.6
* 系统会员账号，有同名提示不能创建
* 获取菜单列表，由判断超管id改为判断为角色id
* 正则过滤open不鉴权接口，一律不记录到日志
* 登录后获取ctx的admin信息，由原来的gjson库改为json.Unmarshal解码

### v1.1.5
* 更新中间件写入日志
* 更新菜单权限
* 日志搜索字段增加ip和action
* 重新更新id为雪花值（升级）
* 修改会员列表超管不显示删除按钮 （前后端）
* 修复会员中心不能修改密码（前后端）
* 修复会员列表编辑没有role-list权限时，不应该显示role选择下拉框，修复报错问题（前后端）
* 优化删除用户列表逻辑（升级）
* 修复日志写入规则

### v1.1.4
- 更新依赖库

### v1.1.3
- 修复问题

### v1.1.2
- 插件修改

### v1.1.0
- 核心包改为线上包，dzhcore版本是v1.1.3
- dict和space，task，member都默认为插件包
- 插件安装之后，插件市场才有选择
- logic支持model和dao调用，controller路由注册改为model接口式

### 20240707
- 修改插件初始化函数名称

### v0.1.1
- 更新获取登录账号的方法
- 修复dzhcore的list逻辑
- 优化一些字段和数据

### v0.1
- 大改，更换成gf的分层设计目录（升级）

### v0.0.95
- 修复超管在用户列表时显示角色
- 更改eps接口不需要登录权限
- cms会员迁移到base主模块
- cms模块的上传接口取消
- 获取登录后的会员信息由core迁移到base
* 修改cms插件

### v0.0.94
- 修复角色权限问题
- 初始化数据更新
- 增加版本控制和获取版本接口
- 增加插件市场
- 修复登录失败没有报错
- 修复添加用户失败
- 修复新增角色，角色表写入失败
- 增加service的where条件支持元素长度为1的传入
- 兼容go1.21的包引入
- 更新默认菜单顺序，任务管理和字典管理并入系统管理

### v0.0.93
- 前后端：增加菜单：字段类型
- 后端：为了后期用户自增数据生成的id不要和初始数据的id冲突，自增id修改为雪花id
- 后端：更改为分层模式，主逻辑在internal，插件移动到addons
- 前后：更新了新版UI

### v0.0.92
* 修改后台前端订单流水和前端会员文件路径，放到cms文件夹下
* 修改代理图片前缀为 /dzhimg

### 20230730 -v0.0.9
- 修改核心命名为dzhcore和其他文件命名
-
### 20230730 -v0.0.2
- 增加会员登录几种模式

### 20230528 -v0.0.1
- 完成基本目录架构




