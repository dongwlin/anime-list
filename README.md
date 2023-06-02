<div align="center">

# AnimeList

<div>
    <img alt="license" src="https://img.shields.io/github/license/lin-Dongw/AnimeList">
    <img alt="commit" src="https://img.shields.io/github/commit-activity/m/lin-Dongw/AnimeList?color=%23ff69b4">
    <img alt="stars" src="https://img.shields.io/github/stars/lin-Dongw/AnimeList?style=social">
</div>

一款简单的番剧管理工具

基于vue和cpp

菜菜，还请指指点点

</div>

## 简介

通过它你可以快速地打开你存放在本地或者网络上的番剧，不用再慢慢翻目录了。

## 背景

因为平时使用rss订阅自动下载番剧，但由于看的番剧越来越多，以至于保存番剧的目录也越来越多，甚至有的番剧是下载在nas里，所以导致有时难以找到想看的番剧的目录。我在想能不能有一个工具可以帮我打开番剧存放的目录，为了实现这个需求从而有了这个项目。（为什么不用AList来管理？我就要自己写一个）

## 功能

- 番剧信息创建
- 番剧信息修改
- 番剧信息删除
- 番剧信息状态管理

（搜索还没做，有空再说）

## 使用说明

很简单的自己摸索一下就会用了

### 关闭服务

在Setting/server里有状态显示，执行停止后就会结束运行。

## 番剧信息数据

相关的数据会存储在同个目录下的data子目录里的anime.json文件，具体请看[anime.json说明](./docs/AnimeData.md)

## 致谢

### 开源库

- 前端框架： [vue3](https://github.com/vuejs/core)
- vue 状态管理库：[pinia](https://github.com/vuejs/pinia)
- UI组件：[element-plus](https://github.com/element-plus/element-plus)
- 请求库：[axios](https://github.com/axios/axios)
- c++ json库：[json](https://github.com/nlohmann/json)
- c++ http库：[cpp-httplib](https://github.com/yhirose/cpp-httplib)



