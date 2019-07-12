package main

import (
    "encoding/json"
    "fmt"
    "github.com/satori/go.uuid"
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "time"
)

const (
    HTTP_SUCCESS int = 1 //成功
    HTTP_ERROR   int = 0 //请求错误/其他错误
)

func HttpError(w http.ResponseWriter, err int, module string) {
    str := fmt.Sprintf("{\"code\":%d, \"message\":\"%s\"}\n", err, module)
    io.WriteString(w, str)
}

func markdownHandler(response http.ResponseWriter, request *http.Request) {

    switch request.Method {
    case "GET":
        if err := request.ParseForm(); err != nil {
            HttpError(response, HTTP_ERROR, "error")
            return
        }

        HttpError(response, HTTP_ERROR, "error")

    case "POST":
        //fmt.Print("action")
        if err := request.ParseForm(); err != nil {
            HttpError(response, HTTP_ERROR, "post error")
            return
        }

        //自定义 请求模式 method决定后续处理函数
        method, ok := request.PostForm["method"]
        if !ok {
            //参数错误
            HttpError(response, HTTP_ERROR, "method not find")
            return
        }

        switch method[0] {
        case "save_as_service":
            //保存到服务器 覆盖保存
            saveMarkdown(response, request)
        default:
            HttpError(response, HTTP_ERROR, "method error")
        }
    default:
        HttpError(response, HTTP_ERROR, "request error")

    }

}

func uploadHandler(response http.ResponseWriter, request *http.Request) {

    //上传为POST 方式
    if err := request.ParseForm(); err != nil {
        HttpError(response, HTTP_ERROR, "post error")
        return
    }

    imgFile, imgHead, imgErr := request.FormFile("editormd-image-file")
    if imgErr != nil {
        HttpError(response, HTTP_ERROR, "upload image error")
        return
    }

    //生成唯一文件名称
    u, _ := uuid.NewV1()
    picFileName := u.String() + imgHead.Filename

    defer imgFile.Close()

    file, err := os.Create(g_conf.PicPath + picFileName)
    if err != nil {
        HttpError(response, HTTP_ERROR, "upload image error")
        return
    }
    defer file.Close()
    _, err = io.Copy(file, imgFile)
    if err != nil {
        HttpError(response, HTTP_ERROR, "upload image error")
        return
    }

    var result Result
    result.Success = HTTP_SUCCESS
    last_modified_time := time.Now().Format("2006-01-02 15:04:05")
    result.Message = "上传时间:" + last_modified_time
    result.Url = g_conf.PicUrl + picFileName

    b, err1 := json.Marshal(result)
    if err1 != nil {
        HttpError(response, HTTP_ERROR, "json error")
        return
    }
    _, err2 := io.WriteString(response, string(b))
    if err2 != nil {
        HttpError(response, HTTP_ERROR, "io.Write error")
        return
    }
    log.Printf("upload success picture name:%s", picFileName)

}

func saveMarkdown(response http.ResponseWriter, request *http.Request) {

    var result Result
    titles := request.PostForm["title"]
    content := request.PostForm["content"][0]
    fmt.Print(content)

    last_modified_time := time.Now().Format("2006-01-02 15:04:05")

    title := titles[0] + time.Now().Format("2006_01_02") + ".md"

    path := g_conf.MarkdownPath + title

    err := ioutil.WriteFile(path, []byte(content), 0666)
    checkErr(err)

    log.Printf("title %s \n", title)

    result.Success = HTTP_SUCCESS
    result.Message = "保存时间:" + last_modified_time
    result.Url = title

    b, err1 := json.Marshal(result)
    if err1 != nil {
        HttpError(response, HTTP_ERROR, "json error")
        return
    }
    _, err = io.WriteString(response, string(b))
    if err != nil {
        HttpError(response, HTTP_ERROR, "io.Write error")
        return
    }
    log.Printf("save markdown success, title:%s", title)

}


func checkErr(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

type Result struct {
    Success int    `json:"success"`
    Message string `json:"message"`
    Url     string `json:"url"`
}

