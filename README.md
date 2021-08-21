# LoadExcel

### 介绍
一个导表工具，可以将Excel表导出为json数据，并同时导出数据表interface文件方便TypeScript和C#环境使用。 

### 环境
go 1.16 + github.com/360EntSecGroup-Skylar/excelize

### 编译
#### Mac下编译Linux和Windows64位可执行程序  
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./dist/LoadExcel.exe  
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./dist/LoadExcel  

#### Linux下编译Mac和Windows64位可执行程序  
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./dist/LoadExcel.exe  
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./dist/LoadExcel  

#### Windows下编译Mac和Linux64位可执行程序  
SET CGO_ENABLED=0  
SET GOOS=darwin  
SET GOARCH=amd64  
go build -o ./dist/LoadExcel  
SET CGO_ENABLED=0  
SET GOOS=linux  
SET GOARCH=amd64  
go build -o ./dist/LoadExcel  

#### 注意
GOOS：目标平台的操作系统（darwin、freebsd、linux、windows）
GOARCH：目标平台的体系架构（386、amd64、arm）
交叉编译不支持CGO所以要禁用它
macos: 为LoadExcel和LoadExcelMac.sh授权可执行权限,修改LoadExcelMac.sh默认为终端打开,双击运行LoadExcelMac.sh  

### 运行  
windows: 双击运行LoadExcelWin.cmd  
macos: 为LoadExcel和LoadExcelMac.sh授权可执行权限,修改LoadExcelMac.sh默认为终端打开,双击运行LoadExcelMac.sh  

### 说明
1、工具会检测当前目录及子目录所有Excel文件，Excel名字中包含～的文件会被忽略  
2、每个Excel表所有sheet数据会导出在同一个Json文件，文件名使用Excel文件名  
3、在LoadExcel中配置 Type:字段类型所在行 Key:字段key所在行 Commit:字段批注所在行 DataStart:数据开始的行 (行数从0开始)  
4、TypeScript数据类型支持：number，string，boolean以及它们的一维二维数组  
5、C#数据类型支持：int，float，string，bool以及它们的一维二维数组  
6、一维数组字段类型后面加[],二维数组字段类型后面加[][]；boolean用1、0表示true、false；一维数组使用(;)分隔,二维数组使用(,)分隔；字段类型不填表示此列数据将不会被导出。    