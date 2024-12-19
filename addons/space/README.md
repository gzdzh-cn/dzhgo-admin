# Space 文件空间

提供上传图片的管理功能,注意在空间删除文件时仅删除了文件在数据库的索引.未删除实际文件.

```bash
gf gen dao -p=addons/space -t=addons_space_info,addons_space_type

gf gen service -s=addons/dzhCrm/logic -d=addons/dzhCrm/service

```
