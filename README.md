# LoadExcel

#### 介绍
一个导表工具，可以将Excel表导出为json数据，并同时导出数据表interface文件方便TypeScript和C#环境使用。 

#### 环境
python3 + xlrd1.2.0

#### 说明
1、工具会检测当前目录及子目录所有Excel文件，Excel名字中包含～的文件会被忽略
2、每个Excel表所有sheet数据会导出在同一个Json文件，文件名使用Excel文件名  
3、在LoadToTS或LoadToCS中配置 字段key所在行 字段类型所在行 数据开始行 (行数从0开始)  
4、TypeScript数据类型支持：none，number，string，boolean，number[]，string[]，boolean[]；boolean用1、0表示true、false；数组使用(;)分隔；none或不填表示此列数据将不会被导出。  
5、C#数据类型支持：none，int，float，string，bool，int[]，float[]，string[]，bool[]；bool用1、0表示true、false；数组使用(;)分隔；none或不填表示此列数据将不会被导出。  