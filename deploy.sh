#!/bin/bash
# shellcheck disable=SC2164
set -x
current=$(pwd)
# shellcheck disable=SC2034
now=$(date "+%Y-%m-%d %H:%M:%S")
# shellcheck disable=SC2034
branch=$1
if [ ! -n "$1" ] ;then
    branch="dev"
else
    # shellcheck disable=SC2034
    branch=$1
fi
echo "当前执行目录：[$current]"
echo "当前执行分支：[$branch]"

function deal() {
    echo "开始执行[$1]"
    cd services/"$1"
    git add .
    git reset --hard
    git checkout -b "$branch"
    git pull origin "$branch"
    cd "$current"
    # shellcheck disable=SC2086
    protoc --php_out=composer --proto_path=services/"$1"/ services/"$1"/proto/$1/*.proto
}


function initComposer() {
    echo "初始化[composer]"
    cd composer
    git add .
    git reset --hard
    git checkout -b "$branch"
    git pull origin "$branch"
    cd "$current"
}
initComposer

# shellcheck disable=SC2006
servers=`ls services`
# shellcheck disable=SC2034
for s in $servers
do
  server="services/$s/proto/$s"
  if [ -d "$current/$server" ];then
    echo "[$server] 存在"
    deal "$s"
    else
    echo "[$server] 不存在 跳过处理"
  fi
done


function pushComposer() {
    echo "提交[composer]"
    cd composer
    git add .
    git commit -m "$now 提交"
    git push origin "$branch"
    cd "$current"
}
pushComposer
