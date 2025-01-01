# twitter-clone
私のweb開発練習の一環としてtwitter-cloneを作ってみることにしました。

# 開発環境
- docker-compose
  - React (フロント)
  - GO (バック) 
  - MySQL (データベース)

# db設計
## users
| カラム名 | データ型 | 制約 |
| :---: | :---: | :---: |
| id | int | primary key, auto_increment |
| user_name | varchar(50) | not null, unique |
| created_at | datetime | not null |
| updated_at | datetime | not null |

## tweets
## likes
## retweets



# ディレクトリ構成
/
|- client
|- server
|-
|

