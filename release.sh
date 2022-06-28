#! /bin/sh
version="1.0.0"

push() {
  git tag v$version
  git push origin v$version
}

test() {
  air
}


# .\release.sh testbuild
testbuild(){
   testbuildmacx86
   testbuildmacm1
   testbuildlinux
   testbuildwindows
}

# .\release.sh testbuildmacx86
testbuildmacx86(){
   CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./tmp/macx86 ./test/test.go
}
# .\release.sh testbuildmacm1
testbuildmacm1(){
   CGO_ENABLED=0 GOOS=darwin GOARCH=arm go build -o ./tmp/macm1 ./test/test.go
}

# .\release.sh testbuildlinux
testbuildlinux(){
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./tmp/linux ./test/test.go
}
# .\release.sh testbuildwindows && ./tmp/win.exe --count 50
testbuildwindows(){
  CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./tmp/win.exe ./test/test.go
}
 v
main() {
  cmd_list=("push test testbuildlinux testbuild testbuildmacx86 testbuildmacm1 testbuildwindows")
  if echo "${cmd_list[@]}" | grep -wq "$1"; then
    "$1"
  else
    echo "Invalid command: $1"
  fi
}

main "$1"
