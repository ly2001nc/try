go env -w GOPROXY=https://goproxy.cn,direct
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/
go env -w GO111MODULE=on
go env -w GOPRIVATE=*.gitlab.com,*.gitee.com,*.github.com,*.gopkg.in
go env -w GOSUMDB=off
或者 go env -w GOSUMDB="sum.golang.google.cn"
go env -w GOMODCACHE=自己设置一个go mod的目录  //配置go mod 的缓存目录