# Anime List

## 简介

Anime List的前端部分

## 依赖安装

```shell
npm i
```

## dev

```shell
npm run dev
```

## build

```shell
npm run build
```

下面的指令是在windows系统下的运行，其他系统需要另外修改

```shell
npm run build
&&
md "./../x64/Debug/www/assets"
||
del /q ".\..\x64\Debug\www\assets\*"
&&
Copy /y "./dist" "./../x64/Debug/www"
&&
Copy /y "./dist/assets" "./../x64/Debug/www/assets"
```
