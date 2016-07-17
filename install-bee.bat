set GOPATH=%~dp0..
go build github.com\beego\bee
copy bee.exe %GOPATH%\bin\bee.exe
del bee.exe
pause