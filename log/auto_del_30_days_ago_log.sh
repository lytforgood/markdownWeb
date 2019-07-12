#!/bin/sh
find /root/markdown/markdownWeb/log/ -mtime +30 -name "*.log" -exec rm -rf {} \;

