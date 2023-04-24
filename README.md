# OJ
Online Judge(目標是給考場使用){目前為練習程式語言用}{在學校時練習的專案}

Web 端:</br>
php-laravel 5.6

Server 端:</br>
php7.0</br>
python3</br>

Client 端:</br>
golang

目前功能:</br>
Web 端教授可以新增題目，學生可以看有哪些題目</br>
python3 為 Client 與 資料庫溝通的中介</br>
golang 為跨平台，教室為任意作業系統都可直接上機考試</br>

目前為Demo版，功能只有一些，但是後面會陸續增加新功能。

操作環境只在 macbook 上測試過。

執行環境以 Docker 來執行，主要有三個容器，python mysql richarvey/nginx-php-fpm。

Docker 指令:</br>
docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=admin -d mysql</br>
docker run -d -p 80:80 --name web --link mysql -v ~/OJ/php/OJ/:/var/www/html/ richarvey/nginx-php-fpm</br>
docker run -d -i -p {any port}:{any port} --name python --link mysql -v ~/OJ/python/server/:{your path} python</br>

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




