@REM 使用脚本可以将json和ts文件复制到指定目录
LoadExcel.exe
@set jsonPath=D:\Workspace\fall_in_love\program\meizi1\assets\resources\data\
@set entityPath=D:\Workspace\fall_in_love\program\meizi1\assets\script\game\
@for /R ./out %%f in (*.json) do ( 
copy /y %%f %jsonPath%
)
@for /R ./out %%f in (*.ts) do ( 
xcopy /Y %%f %entityPath%
)
pause