#!/bin/bash

set -x
current=$(pwd)
echo "当前目录：[$current]"

branch=$1

if [ ! -n "$1" ]; then
    echo "必须传入分支名字执行"
    exit 0
fi

echo "收到参数"
echo $branch

su - www <<EOF
cd $(readlink -f $current)
git checkout .
git fetch
git checkout "$branch"
git pull
go build -o server cmd/server/main.go
exit
EOF

supervisorctl stop micro-merchant-basic:micro-merchant-basic_00
supervisorctl start micro-merchant-basic:micro-merchant-basic_00
echo "部署完成"
