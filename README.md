# golang tumblr爬虫
1. 在`sites.txt`中添加指定tumblr主页，比如`http://allthingseurope.tumblr.com/`只需要添加`allthingseurope`
2. 多条分行添加
3. 代理，默认使用本地代理127.0.0.1:1080端口代理。如需修改请将`common/proxyHttp.go`中的`127.0.0.1:1080`改成自己代理。
4. 目前已完成图片和视频的下载
5. 采用channel控制并发，goroutine有点少，下载速度一般
