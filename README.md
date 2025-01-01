# twitter-clone
私のweb開発練習の一環としてtwitter-cloneを作ってみることにしました。

# 開発環境
- docker-compose
  - React (フロント)
  - GO (バック) 
  - MySQL (データベース)

# db設計
db設計には以下のサイトを参考にしました。
[X（旧Twitter）のデータベース設計をER図で可視化する](https://qiita.com/at-sacai/items/6ec62eac0b823da9aa9f#3-%E3%83%86%E3%83%BC%E3%83%96%E3%83%AB%E3%81%AE%E8%A9%B3%E7%B4%B0%E8%A8%AD%E8%A8%88)

## users
| カラム名 | データ型 | 制約 |
| :---: | :---: | :---: |
| id | int | primary key, auto_increment |
| user_name | varchar(50) | not null, unique |
| created_at | datetime | not null |
| updated_at | datetime | not null |

## follows
| カラム名 | データ型 | 制約 |
| :---: | :---: | :---: |
| id | int | primary key, auto_increment |
| follower_id | int | foreign key(users) , not null |
| followed_id | int | foreign key(users) , not null |
| created_at | datetime | not null |
| updated_at | datetime | not null |

## tweets
| カラム名 | データ型 | 制約 |
| :---: | :---: | :---: |
| id | int | primary key, auto_increment |
| user_id | int | foreign key(users) , not null |
| content | varchar(200) |  , not null |
| created_at | datetime | not null |
| updated_at | datetime | not null |

## likes
| カラム名 | データ型 | 制約 |
| :---: | :---: | :---: |
| id | int | primary key, auto_increment |
| user_id | int | foreign key(users) , not null |
| tweet_id | int | foreign key(tweets) , not null |
| created_at | datetime | not null |
| updated_at | datetime | not null |

## retweets
| カラム名 | データ型 | 制約 |
| :---: | :---: | :---: |
| id | int | primary key, auto_increment |
| user_id | int | foreign key(users) , not null |
| tweet_id | int | foreign key(tweets) , not null |
| created_at | datetime | not null |
| updated_at | datetime | not null |

#ER図


# ディレクトリ構成


