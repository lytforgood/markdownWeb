package main

import (
    "encoding/json"
    "flag"
    "io/ioutil"
    "log"
)

var (
    g_conf_file = flag.String("conf", "config.json", "config file")
    g_conf      goWebConfig
)


type goWebConfig struct {
    Version      string `josn:"version"`      //版本
    MarkdownPath string `json:"markdownpath"` //存放文件目录
    PicPath      string `json:"picpath"`      //存放图片目录
    PicUrl       string `json:"picurl"`       //存放图片目录
}

func initConfig() {
    buf, err := ioutil.ReadFile(*g_conf_file)
    if err != nil {
        log.Fatalf("Read config file error, %v\n", err)
    }

    if err = json.Unmarshal(buf, &g_conf); err != nil {
        log.Fatalf("Read config file error, %v\n", err)
    }
}

