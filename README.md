# LoadExcel

#### 介绍
一个导表工具，可以将Excel表导出为json数据，并同时导出数据表interface文件方便TypeScript和C#环境使用。 

#### 环境
nodejs + typescript + exceljs + pkg(导出为各系统执行文件)  

#### 说明
1、工具会检测当前目录及子目录所有Excel文件，Excel名字中包含～的文件会被忽略  
2、每个Excel表所有sheet数据会导出在同一个Json文件，文件名使用Excel文件名  
3、在LoadExcel中配置 Type:字段类型所在行 Key:字段key所在行 Commit:字段批注所在行 DataStart:数据开始的行 (行数从1开始)  
4、TypeScript数据类型支持：number，string，boolean以及它们的一维二维数组(一维[],二维[][])；boolean用1、0表示true、false；一维数组使用(;)分隔,二维数组使用(,)分隔 ；字段类型不填表示此列数据将不会被导出。  
5、C#数据类型支持：int，float，string，bool以及它们的一维二维数组(一维[],二维[][])；bool用1、0表示true、false；一维数组使用(;)分隔,二维数组使用(,)分隔 ；字段类型不填表示此列数据将不会被导出。  