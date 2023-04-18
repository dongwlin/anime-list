# AnimeList

## anime.json

| 字段   | 类型   | 默认值 | 说明         |
| ------ | ------ | ------ | ------------ |
| id     | int    |        | 编号         |
| name   | string |        | 名称         |
| status | bool   | true   | 状态         |
| type   | int    | -1     | 类型         |
| dir    | string |        | 目标dir      |
| url    | string |        | 目标url      |
| img    | string | none   | 番剧图像     |
| day    | int    | -1     | 番剧更新星期 |

### type

| key  | value   |
| ---- | ------- |
| -1   | none    |
| 0    | network |
| 1    | local   |

### day

| key  | value     |
| ---- | --------- |
| -1   | none      |
| 0    | Sunday    |
| 1    | Monday    |
| 2    | Tuesday   |
| 3    | Wednesday |
| 4    | Thursday  |
| 5    | Friday    |
| 6    | Saturday  |



