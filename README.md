[TOC]

# 如何搭建在线markdown
1. 搭建web服务器 部署markdownWeb eg:caddy
2. 部署webapi 部署log/auto_del_30_days_ago_log.sh 自动删除日志
3. 设置本地代理 方便ajax访问webapi 参考markdown在线编辑器搭建.md
		```
		示例 caddy 配置

		:11223 {
		root /root/markdown/markdownWeb/markdownWeb
		gzip
		browse
		proxy  /markdown 127.0.0.1:11224
		}
		:80 {
		gzip
		browse
		proxy  / 127.0.0.1:11223
		}
		```