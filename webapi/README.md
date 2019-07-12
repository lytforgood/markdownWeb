编译成指定名称
go build -o goapi
nohup ./goapi >/root/log/httpApi_$(date +%Y-%m-%d).log 2>&1 &
将/root/log/目录下所有30天前带".log"的文件删除
find /root/log/ -mtime +30 -name "*.log" -exec rm -rf {} \;
vim auto_del_30_days_ago_log.sh
```
#!/bin/sh
find /root/log/ -mtime +30 -name "*.log" -exec rm -rf {} \;
```
每天凌晨0点10分执行auto-del-7-days-ago-log.sh文件进行数据清理任务
crontab -e
10 0 * * * /root/log/auto_del_30_days_ago_log.sh >/dev/null 2>&1