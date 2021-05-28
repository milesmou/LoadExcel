@REM 使用脚本可以将json和ts,cs文件复制到指定目录

@set jsonPath=..\assets\resources\data\
@set entityPath=..\assets\script\game\

.\LoadExcel.exe ts

@for /R ./out %%f in (*.json) do ( 
xcopy /Y %%f %jsonPath%
)

@for /R ./out %%f in (*.ts) do ( 
xcopy /Y %%f %entityPath%
)

@for /R ./out %%f in (*.cs) do ( 
xcopy /Y %%f %entityPath%
)

pause