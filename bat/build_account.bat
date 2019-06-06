set GOPATH=%cd%/../../../../..

go build -i -o ../bin/account.exe github.com/huoshan017/ib_server/src/account
if errorlevel 1 goto exit

goto ok

:exit
echo build failed !!!

:ok
echo build ok