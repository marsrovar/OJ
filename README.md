# OJ
Online Judge(設計給電腦教室考試用)

目前為Demo版，功能只有一些，但是後面會陸續增加新功能。

目前操作環境只測試在 macbook 上。

執行環境以 Docker 來執行，主要有三個容器，python mysql richarvey/nginx-php-fpm。

Docker 指令:</br>
docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=admin -d mysql
docker run -d -p 80:80 --name web --link mysql -v ~/OJ/php/OJ/:/var/www/html/ richarvey/nginx-php-fpm
docker run -d -i -p 1121:1121 --name python --link mysql -v ~/OJ/python/server/:{your path} python

golang 為前端環境執行，在 macbook 上。

TUI 的 terminal 大小為 155X32。

前端的TUI圖:</br>
Login
![image](https://github.com/alfie1121/OJ/blob/master/Login.png)

Question Pool List
![image](https://github.com/alfie1121/OJ/blob/master/QPlist.png)

Code
![image](https://github.com/alfie1121/OJ/blob/master/Code.png)

Code Tool
![image](https://github.com/alfie1121/OJ/blob/master/CodeTool.png)




