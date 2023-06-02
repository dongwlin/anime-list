# Anime数据存储

## 字段一览

| 字段   | 类型   | 默认值 | 说明         |
| ------ | ------ | ------ | ------------ |
| id     | string |        | 编号         |
| name   | string |        | 名称         |
| status | bool   | true   | 状态         |
| type   | int    | -1     | 类型         |
| dir    | string |        | 目标dir      |
| url    | string |        | 目标url      |
| img    | string |        | 番剧图像     |
| day    | int    | -1     | 番剧更新星期 |

## 特殊字段说明

### type

| key  | value   | 说明         |
| ---- | ------- | ------------ |
| -1   | none    | 默认值       |
| 0    | network | 网络跳转     |
| 1    | local   | 打开本地目录 |

### day

| key  | value     | 说明   |
| ---- | --------- | ------ |
| -1   | none      | 默认值 |
| 0    | Sunday    | 周日   |
| 1    | Monday    | 周一   |
| 2    | Tuesday   | 周二   |
| 3    | Wednesday | 周三   |
| 4    | Thursday  | 周四   |
| 5    | Friday    | 周五   |
| 6    | Saturday  | 周六   |

