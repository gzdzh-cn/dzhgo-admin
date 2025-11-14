# internal


## 资源打包命令

```bash
gf pack internal/resource internal/packed/packed.go -p internal/resource

# 指定internal里面的表
gf gen dao -t=base_eps_admin,base_eps_app,base_sys_addons,base_sys_addons_types,base_sys_conf,base_sys_department,base_sys_init,base_sys_log,base_sys_menu,base_sys_param,base_sys_role,base_sys_role_department,base_sys_role_menu,base_sys_setting,base_sys_user,base_sys_user_role,base_sys_action_log,base_sys_notice,base_sys_notice_user_read,base_sys_feedback

gf gen service

```
