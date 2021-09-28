# 该脚本需要cd到测试文件所在目录下执行
method=$1
# 不传入method参数则运行全部测试
cmd="go test -v -coverprofile=test.out -gcflags=all=-l"
# 传入method参数则测试指定方法
if [ "$1" != "" ]; then
    fileList=$(ls)
    cmd="go test -v"
    for file in $fileList; do
        if echo "$file" | grep -q -E '\.go$'; then
            cmd="${cmd} $file"
        fi
    done
    cmd="$cmd -test.run $method -coverprofile=test.out -gcflags=all=-l"
fi
# 执行测试
$cmd
# 生成测试覆盖率html报告
go tool cover -html=test.out -o test.html
