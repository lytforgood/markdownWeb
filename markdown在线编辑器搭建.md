# markdown-notepad 在线编辑器
安装Caddy
```
wget -N --no-check-certificate https://raw.githubusercontent.com/ToyoDAdoubiBackup/doubi/master/caddy_install.sh && chmod +x caddy_install.sh && bash caddy_install.sh
## 备用地址
wget -N --no-check-certificate https://www.moerats.com/usr/shell/Caddy/caddy_install.sh && chmod +x caddy_install.sh && bash caddy_install.sh
```
下载源码
```
https://github.com/zzzhan/markdown-notepad
wget https://www.moerats.com/usr/down/markdown-notepad.zip && unzip markdown-notepad.zip
```
配置Caddy
/usr/local/caddy/Caddyfile
指定/root/markdown/markdown-notepad 为web访问目录
```
#如果你想用IP登录，使用以下命令，注意端口80可以自己修改，登录地址为ip:port
echo ":11223 {
 root /root/markdown/markdown-notepad
 gzip
 browse
}" > /usr/local/caddy/Caddyfile

#如果你想域名登录，使用以下命令，记得修改域名
echo "http://moerats.com {
 root /root/markdown-notepad
 gzip
 browse
}" > /usr/local/caddy/Caddyfile

#如果你想使用域名https登录，使用以下命令
echo "https://moerats.com {
 root /root/markdown-notepad
 tls admin@moerats.com | /root/xxx.crt /root/xxx.key
 gzip
 browse
}" > /usr/local/caddy/Caddyfile
```


启动Caddy
```
/etc/init.d/caddy start

/etc/init.d/caddy restart
caddy start  #开启Caddy Web Server
caddy stop  #停止Caddy Web Server
caddy restart  #重启Caddy Web Server
caddy status  #查看Caddy Web Server状态
caddy install  #安装Caddy Web Server
```

# editor.md在线编辑器
1.搭建相关web服务器 eg:tomcat Caddy(参考上面安装)
2.下载editor.md源码 解压到指定web目录
3.利用example里面例子搭建环境
根据例子配置对应好路径 一般需要css  fonts  images  js  lib   plugins
https://blog.csdn.net/LoveJavaYDJ/article/details/73692917
https://www.jianshu.com/p/d0eed194db65
4.自己可DIY截图、自动保存、下载
https://www.codehui.net/info/39.html
https://www.codehui.net/info/40.html
只显示不能编辑等功能查看官方示例（ https://pandao.github.io/editor.md/examples/index.html ） 
看效果，从examples找对应的html，对应着添加配置文件即可
或者参考下其他的文档：
https://blog.csdn.net/LoveJavaYDJ/article/details/73692917

## 简易编写实现参数说明
```
<!DOCTYPE html>
<html lang="zh">
<head>
    <meta charset="utf-8" />
    <title>Simple example - Editor.md examples</title>
    <link rel="stylesheet" href="css/style.css" />
    <link rel="stylesheet" href="css/editormd.css" />
    <link rel="shortcut icon" href="css/favicon.ico" type="image/x-icon" />
</head>
<body>
    <div id="layout">
        <header>
            <h1>Markdown Online</h1>
        </header>
        <div class="btns">
            <button id="get-md-btn">Get Markdown</button>
        </div>
        <div id="test-editormd">
            <textarea style="display:none;">[TOC]
            </textarea>
        </div>
    </div>
    <script src="/markdown/examples/js/jquery.min.js"></script>
    <script src="/markdown/editormd.js"></script>
    <script type="text/javascript">
        $(function() {
            testEditor = editormd("test-editormd", {
            width        : "100%",//宽度
            height       : 720,//高度
            path         : '/markdown/lib/',//这块是lib的文件路径，必须设置，否则几个样式css，js访问不到的
            //以下内容选择性填写
            syncScrolling : "single", //滚动方式
            flowChart : true,//控制流程图编辑
            sequenceDiagram : true,//控制时序图编辑
            taskList : true,//任务列表
            tex  : true,//科学公式
            emoji : true,//moji表情
            htmlDecode : "style,script,iframe|on*", // 开启 HTML 标签解析，为了安全性，默认不开启
            codeFold : true,//ctrl+q代码折叠
//          saveHTMLToTextarea : true,//这个配置在simple.html中并没有，但是为了能够提交表单，使用这个配置可以让构造出来的HTML代码直接在第二个隐藏的textarea域中，方便post提交表单。这个转换好像有些缺陷，有些配置没有生效，目前还没想到怎么解决，我这里没有用,是在前台读取的时候转换的
            imageUpload : true,//开启本地图片上传
            imageFormats : ["jpg", "jpeg", "gif", "png", "bmp", "webp"],
            imageUploadURL : "/index.php/Admin/News/uploadFileMark",//图片上传地址 可以是后台上传地址
            /*上传图片成功后可以做一些自己的处理*/
            onload : function() {
                console.log('onload', this);
            },
            //设置标签显示内容
            toolbarIcons : function() {
            // Or return editormd.toolbarModes[name]; // full, simple, mini
            // Using "||" set icons align right.
            return ["undo", "redo", "|", "bold", "hr", "|", "preview", "watch", "|", "fullscreen", "info", "testIcon", "testIcon2", "file", "faicon", "||", "watch", "fullscreen", "preview", "testIcon"]
            }
            /**设置主题颜色*/
            // 下面三个选项是设置风格的，每个有什么风格，请看下载插件得examples/themes.html
            theme        : "lesser-dark",// 工具栏风格
            previewTheme : "dark",// 预览页面风格
            editorTheme  : "paraiso-dark",// 设置编辑页面风格
        });
        //设置可以添加目录，只需要在上面一行输入 [TOCM]，就会有目录，注意下面要空一行
        testEditor.config({
            tocm : true,
            tocContainer : "",
            tocDropdown   : false
        });
        //获取markdown文本 后续可调用xx函数 写入本地/传到服务器
        $("#get-md-btn").bind('click', function(){
            alert(testEditor.getMarkdown());
        });
    });
</script>
</body>
</html>
```
## localStorage自动保存
Editor.md自动保存插件的开发
https://www.codehui.net/info/40.html

## 自定义工具栏内容
```
t.toolbarModes={
    full:["undo","redo","|","bold","del","italic","quote","ucwords","uppercase","lowercase","|","h1","h2","h3","h4","h5","h6","|","list-ul","list-ol","hr","|","link","reference-link","image","code","preformatted-text","code-block","table","datetime","emoji","html-entities","pagebreak","|","goto-line","watch","preview","fullscreen","clear","search","|","help","info"],
    simple:["undo","redo","|","bold","del","italic","quote","uppercase","lowercase","|","h1","h2","h3","h4","h5","h6","|","list-ul","list-ol","hr","|","watch","preview","fullscreen","|","help","info"],
    mini:["undo","redo","|","watch","preview","|","help","info"]
}
```
图标查询
css/editormd.css 下面搜索 fa-  eg:fa-save
```
可能用到的图标 自定义下载
fa-save 保存到服务器  //fa-cloud 获得缓存 //fa-cloud-download 下载到本地
```


## 图片上传
首先testEditor 将imageUpload设置为True 且填写上传服务地址
editor.md期望你上传图片的服务返回如下json格式的内容
```
{
    success : 0 | 1, //0表示上传失败;1表示上传成功
    message : "提示的信息",
    url     : "图片地址" //上传成功时才返回
}
```

## 图片粘贴插入插件
https://www.codehui.net/info/39.html

## 动态内容文件下载到本地
借助HTML5 Blob实现文本信息文件下载
https://www.cnblogs.com/goloving/p/7651636.html


## Caddy反向代理 实现本机跨域目的
vim /usr/local/caddy/Caddyfile
/etc/init.d/caddy restart
将指定/markdown后缀的路径 转发到127.0.0.1:11224
```
:11223 {
root /root/markdown/markdownWeb/markdownWeb
gzip
browse
proxy  /markdown 127.0.0.1:11224
}
```

## ajax访问
表单数据contentType一定要选择
application/x-www-form-urlencoded;charset=utf-8
```
$.ajax({
    //请求方式
    type : "POST",
    // //请求的媒体类型
    contentType: "application/x-www-form-urlencoded;charset=utf-8",//这里
    //预期的服务器响应的数据类型
    dataType:'json',
    //请求地址
    url : "/markdown",
    //数据，json字符串
    data : {"method":"save_as_service","title":"case","content":textToWrite},
    //请求成功
    success : function(result) {
        $("#tipMsg").text(result["message"]);
    },
    // 请求失败，包含具体的错误信息
    error: function(XMLHttpRequest, textStatus, errorThrown) {
        alert(XMLHttpRequest.status);
        alert(XMLHttpRequest.readyState);
        alert(textStatus);
    }
});
```
