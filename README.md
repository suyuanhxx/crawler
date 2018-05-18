# golang tumblr爬虫
1. 在`sites.txt`中添加指定tumblr主页，比如`http://allthingseurope.tumblr.com/`只需要添加`allthingseurope`
2. 多条分行添加
3. 代理，默认使用本地代理127.0.0.1:1080端口代理。如需修改请将`common/proxyHttp.go`中的`127.0.0.1:1080`改成自己代理。
4. 目前已完成图片和视频的下载
5. 采用channel控制并发，goroutine有点少，下载速度一般
6. 代码进一步优化，修复空指针问题
7. 已采用waitGroup控制并发，最大程度开启goroutine，下载速度飞快，但一般情况下会将机器直接卡死......
