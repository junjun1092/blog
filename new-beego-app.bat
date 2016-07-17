set APP=coscms.com
set GOPATH=%~dp0..
set BEE=%GOPATH%\bin\bee
%BEE% new %APP%
cd %APP%
echo %BEE% run %APP%.exe > run.bat
echo pause >> run.bat
start run.bat
pause
start http://127.0.0.1:8080