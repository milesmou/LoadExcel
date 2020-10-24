# LoadExcel

#### 介绍
一个导表工具，可以将当前目录以及子目录所有Excel表导出为json数据，并同时导出数据表interface文件方便TypeScript环境使用。  

#### 环境
go

#### 说明
1、脚本会将所有的Excel表数据导出在一个json文件中，并以Excel文件名区分。  
2、Excel数据格式：每一行为一条数据；每条数据第1列作为此条数据id，不可重复；每张Excel只读取第1个sheet，通过Excel文件名来区分数据。  
3、数据类型支持：any，number，string，boolean，number[]，string[]，boolean[]；any不对数据做任何处理，主要用于同一列数据类型不同时；boolean用0和1表示；数组使用(;)分隔。  