# 安装errcheck
go install github.com/kisielk/errcheck@latest

errcheck github.com/kisielk/errcheck/testdata

# 要检查当前目录下的所有软件包
errcheck ./...

# 检查所有软件包：$GOPATH $GOROOT
errcheck all

# 使用标志指定包含函数列表的文件的路径 被排除在外 -exclude
errcheck -exclude errcheck_excludes.txt path/to/package
