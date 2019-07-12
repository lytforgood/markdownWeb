package main

import (
    "log"
    "net/http"
)

func main() {

    //初始化配置
    initConfig()

    log.Printf("markdown api start version:%s", g_conf.Version)

    url := ":11224"

    http.HandleFunc("/markdown", markdownHandler)

    http.HandleFunc("/markdown/upload", uploadHandler)

    err := http.ListenAndServe(url, nil)

    if err != nil {
        log.Printf("markdown ListenAndServe:%s", err)
    }

}

