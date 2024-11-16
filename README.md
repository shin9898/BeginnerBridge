## Users テーブル

| Column              | Type       | Options                   | 説明                       |
|---------------------|------------|---------------------------|----------------------------|
| username            | string     | null: false               | ユーザー名                  |
| email               | string     | null: false, unique: true | メールアドレス              |
| encrypted_password  | string     | null: false               | 暗号化されたパスワード       |
| affiliation         | string     |                           | 所属                       |
| site_url            | string     |                           | サイト・ブログのURL          |
| bio                 | text       |                           | 自己紹介                   |
| user_experience_id  | integer    | null: false               | ユーザーの経験（ActiveHashで管理） |
| created_at          | datetime   | null: false               | 作成日時                   |
| updated_at          | datetime   | null: false               | 更新日時                   |

### Association
- has_many :posts
- has_many :answers
- has_many :comments
- has_many :promotions
- has_many :likes_posts
- has_many :likes_answers
- has_many :likes_comments
- has_many :likes_promotions
- has_many :reported_posts
- has_many :reported_answers
- has_many :reported_comments
- has_many :reported_promotions
- has_many :follows
- has_many :followers, through: :follows

---

## Posts テーブル

| Column             | Type       | Options                   | 説明                       |
|--------------------|------------|---------------------------|----------------------------|
| user               | references | null: false, foreign_key: true | 投稿者                     |
| title              | string     | null: false               | タイトル                   |
| content            | text       | null: false               | 内容                       |
| post_category_id   | integer    | null: false               | 質問・意見のカテゴリー（ActiveHash）|
| goal               | text       |                           | 実現したいこと             |
| attempts           | text       |                           | 試したこと・調べたこと      |
| source_code        | text       |                           | 該当ソースコード           |
| image              | string     |                           | 添付画像                   |
| created_at         | datetime   | null: false               | 作成日時                   |
| updated_at         | datetime   | null: false               | 更新日時                   |

### Association
- belongs_to :user
- has_many :answers
- has_many :comments
- has_many :likes_posts
- has_many :reported_posts

---

## Answers テーブル

| Column             | Type       | Options                   | 説明                       |
|--------------------|------------|---------------------------|----------------------------|
| post               | references | null: false, foreign_key: true | 対応する投稿               |
| user               | references | null: false, foreign_key: true | 回答者                     |
| content            | text       | null: false               | 回答内容                   |
| created_at         | datetime   | null: false               | 作成日時                   |
| updated_at         | datetime   | null: false               | 更新日時                   |

### Association
- belongs_to :post
- belongs_to :user
- has_many :likes_answers
- has_many :reported_answers

---

## Comments テーブル

| Column             | Type       | Options                   | 説明                       |
|--------------------|------------|---------------------------|----------------------------|
| post               | references | foreign_key: true         | 対応する投稿（任意）        |
| promotion          | references | foreign_key: true         | 対応するプロモーション（任意）|
| user               | references | null: false, foreign_key: true | コメント投稿者             |
| content            | text       | null: false               | コメント内容               |
| created_at         | datetime   | null: false               | 作成日時                   |
| updated_at         | datetime   | null: false               | 更新日時                   |

### Association
- belongs_to :post, optional: true
- belongs_to :promotion, optional: true
- belongs_to :user
- has_many :likes_comments
- has_many :reported_comments

---

## Promotions テーブル

| Column             | Type       | Options                   | 説明                       |
|--------------------|------------|---------------------------|----------------------------|
| user               | references | null: false, foreign_key: true | 宣伝投稿者                 |
| title              | string     | null: false               | 宣伝タイトル               |
| content            | text       | null: false               | 宣伝内容                   |
| promotion_category_id | integer  | null: false               | 宣伝カテゴリー（ActiveHashで管理）|
| url                | string     |                           | リンク先URL                 |
| image              | string     |                           | 添付画像                   |
| created_at         | datetime   | null: false               | 作成日時                   |
| updated_at         | datetime   | null: false               | 更新日時                   |

### Association
- belongs_to :user
- has_many :comments
- has_many :likes_promotions
- has_many :reported_promotions

---

## Likes テーブル

### LikesPosts テーブル
| Column             | Type       | Options                   | 説明                       |
|--------------------|------------|---------------------------|----------------------------|
| user               | references | null: false, foreign_key: true | いいねをしたユーザー         |
| post               | references | null: false, foreign_key: true | 対応する投稿                |

### LikesAnswers テーブル
| Column             | Type       | Options                   | 説明                       |
|--------------------|------------|---------------------------|----------------------------|
| user               | references | null: false, foreign_key: true | いいねをしたユーザー         |
| answer             | references | null: false, foreign_key: true | 対応する回答                |

### LikesComments テーブル
| Column             | Type       | Options                   | 説明                       |
|--------------------|------------|---------------------------|----------------------------|
| user               | references | null: false, foreign_key: true | いいねをしたユーザー         |
| comment            | references | null: false, foreign_key: true | 対応するコメント            |

### LikesPromotions テーブル
| Column             | Type       | Options                   | 説明                       |
|--------------------|------------|---------------------------|----------------------------|
| user               | references | null: false, foreign_key: true | いいねをしたユーザー         |
| promotion          | references | null: false, foreign_key: true | 対応するプロモーション       |

---

## Reports テーブル

### ReportedPosts テーブル
| Column             | Type       | Options                   | 説明                       |
|--------------------|------------|---------------------------|----------------------------|
| user               | references | null: false, foreign_key: true | 報告をしたユーザー           |
| post               | references | null: false, foreign_key: true | 報告対象の投稿              |
| reason             | text       | null: false               | 報告理由                   |

### ReportedAnswers テーブル
| Column             | Type       | Options                   | 説明                       |
|--------------------|------------|---------------------------|----------------------------|
| user               | references | null: false, foreign_key: true | 報告をしたユーザー           |
| answer             | references | null: false, foreign_key: true | 報告対象の回答              |
| reason             | text       | null: false               | 報告理由                   |

### ReportedComments テーブル
| Column             | Type       | Options                   | 説明                       |
|--------------------|------------|---------------------------|----------------------------|
| user               | references | null: false, foreign_key: true | 報告をしたユーザー           |
| comment            | references | null: false, foreign_key: true | 報告対象のコメント           |
| reason             | text       | null: false               | 報告理由                   |

### ReportedPromotions テーブル
| Column             | Type       | Options                   | 説明                       |
|--------------------|------------|---------------------------|----------------------------|
| user               | references | null: false, foreign_key: true | 報告をしたユーザー           |
| promotion          | references | null: false, foreign_key: true | 報告対象のプロモーション     |
| reason             | text       | null: false               | 報告理由                   |

---

## Follows テーブル

| Column             | Type       | Options                                        | 説明                        |
|--------------------|------------|------------------------------------------------|----------------------------|
| follower_id        | references | null: false, foreign_key: { to_table: :users } | フォローするユーザー          |
| followed_id        | references | null: false, foreign_key: { to_table: :users } | フォローされるユーザー        |
| created_at         | datetime   | null: false                                    | フォローした日時              |

### Association
- belongs_to :follower, class_name: "User"
- belongs_to :followed, class_name: "User"

---

## Categories テーブル（ActiveHash）

### user_experience
| id | name          |
|----|---------------|
| 1  | 未経験者       |
| 2  | 初心者         |
| 3  | 中級者         |
| 4  | 上級者         |

### post_category
| id | name          |
|----|---------------|
| 1  | プログラミング |
| 2  | キャリア       |
| 3  | その他         |

### promotion_category
| id | name          |
|----|---------------|
| 1  | 企業           |
| 2  | 学習塾         |
| 3  | その他         |
