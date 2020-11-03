# LoadExcel

#### 介绍
一个导表工具，可以将Excel表导出为json数据，并同时导出数据表interface文件方便TypeScript环境使用。  


#### 环境
go in go mod

#### 说明
1、工具会检测当前目录及子目录所有Excel文件  
2、每个Excel表所有sheet数据会导出在同一个Json文件，文件名使用Excel文件名 
4、字段key默认第2行 字段类型默认第1行 数据默认从第4行开始 (行数从0开始) 
5、数据类型支持：none，number，string，boolean，number[]，string[]，boolean[]；boolean用0和1表示；数组使用(;)分隔；none表示此列数据将不会被导出。  