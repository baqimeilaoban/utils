# 数据库算法题

## 游戏玩法分析

题目：
活动表 Activity：
| Column Name  | Type    |
| :----:       | :----:  |
| player_id    | int     |
| device_id    | int     |
| event_date   | date    |
| games_played | int     |

表的主键是 (player_id, event_date)。
这张表展示了一些游戏玩家在游戏平台上的行为活动。
每行数据记录了一名玩家在退出平台之前，当天使用同一台设备登录平台后打开的游戏的数目（可能是 0 个）。

### 1.写一条SQL查询获取每位玩家第一次登录平台日期

查询结果的格式如下所示：
Activity 表：
| player_id | device_id | event_date | games_played |
|:----:|:----:|:----:|:----:|
| 1         | 2         | 2016-03-01 | 5            |
| 1         | 2         | 2016-05-02 | 6            |
| 2         | 3         | 2017-06-25 | 1            |
| 3         | 1         | 2016-03-02 | 0            |
| 3         | 4         | 2018-07-03 | 5            |

Result表：
| player_id | first_login |
|:----:|:----:|
| 1         | 2016-03-01  |
| 2         | 2017-06-25  |
| 3         | 2016-03-02  |

答案如下：

```SQL
SELECT 
    player_id, 
    min(event_date) AS first_login
FROM 
    Activity 
GROUP BY
    player_id 
```

### 2.请编写一个 SQL 查询，描述每一个玩家首次登陆的设备名称

查询结果格式在以下示例中：
Activity 表：
| player_id | device_id | event_date | games_played |
|:----:|:----:|:----:|:----:|
| 1         | 2         | 2016-03-01 | 5            |
| 1         | 2         | 2016-05-02 | 6            |
| 2         | 3         | 2017-06-25 | 1            |
| 3         | 1         | 2016-03-02 | 0            |
| 3         | 4         | 2018-07-03 | 5            |

Result表：
| player_id | device_id |
|:----:|:----:|
| 1         | 2  |
| 2         | 3  |
| 3         | 1  |

答案如下：

```SQL
SELECT
  player_id,
  device_id
FROM
  Activity
WHERE
  (player_id, event_date) IN (
    SELECT
      player_id,
      MIN(event_date)
    FROM
      Activity
    GROUP BY
      player_id
  )
```

### 3.编写一个 SQL 查询，同时报告每组玩家和日期，以及玩家到目前为止玩了多少游戏。也就是说，在此日期之前玩家所玩的游戏总数

查询结果格式在以下示例中：
Activity 表：
| player_id | device_id | event_date | games_played |
|:----:|:----:|:----:|:----:|
| 1         | 2         | 2016-03-01 | 5            |
| 1         | 2         | 2016-05-02 | 6            |
| 1         | 3         | 2017-06-25 | 1            |
| 3         | 1         | 2016-03-02 | 0            |
| 3         | 4         | 2018-07-03 | 5            |

Result table:
| player_id | event_date | games_played_so_far |
|:-:|:-:|:-:|
| 1         | 2016-03-01 | 5                   |
| 1         | 2016-05-02 | 11                  |
| 1         | 2017-06-25 | 12                  |
| 3         | 2016-03-02 | 0                   |
| 3         | 2018-07-03 | 5                   |

对于 ID 为 1 的玩家，2016-05-02 共玩了 5+6=11 个游戏，2017-06-25 共玩了 5+6+1=12 个游戏。
对于 ID 为 3 的玩家，2018-07-03 共玩了 0+5=5 个游戏。
请注意，对于每个玩家，我们只关心玩家的登录日期

答案如下：
```SQL

```
