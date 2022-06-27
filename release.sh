#! /bin/sh
version="1.0.0"

push() {
  git tag v$version
  git push origin v$version
}

test() {
  air
}

testbuildmac(){
  go build -o ./tmp/main ./test/test.go
}

testbuildlinux(){
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./tmp/main ./test/test.go
}

testbuildwindows(){
  CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./tmp/main ./test/test.go
}

main() {
  cmd_list=("push test testbuildlinux testbuildmac testbuildwindows")
  if echo "${cmd_list[@]}" | grep -wq "$1"; then
    "$1"
  else
    echo "Invalid command: $1"
  fi
}

main "$1"
