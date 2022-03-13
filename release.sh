#! /bin/sh
name="meow-whisper-sfu-server"
port=(15302 15303 3478)
branch="main"

gitpull() {
  echo "-> 正在拉取远程仓库"
  git reset --hard
  git pull origin $branch
}

dependencies() {
  echo "-> 正在准备相关资源"
  # 删除无用镜像
  docker rm $(docker ps -q -f status=exited)
  docker rmi -f $(docker images | grep '<none>' | awk '{print $3}')
}

build() {
  echo "-> 准备构建Docker"
  docker build -t $name .
}

run() {
  echo "-> 准备运行Docker"
  docker stop $name
  docker rm $name
  runcmd="docker run --name=$name "
  for v in ${port[*]}; do
    runcmd="${runcmd}-p ${v}:${v} "
  done

  runcmd="${runcmd} --restart=always -d ${name}"
  $runcmd
}

logs() {
  docker logs -f $name
}

start() {
  echo "-> 正在启动「${name}」服务"
  # gitpull
  dependencies
  build
  run
  logs
}

main() {
  cmd_list=("logs start")
  if echo "${cmd_list[@]}" | grep -wq "$1"; then
    "$1"
  else
    echo "Invalid command: $1"
  fi
}

main "$1"
