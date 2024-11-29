class CreatePosts < ActiveRecord::Migration[7.0]
  def change
    create_table :posts do |t|
      t.references :user, null: false, foreign_key: true # 投稿者
      t.string :title, null: false                       # タイトル
      t.text :content, null: false                       # 内容
      t.integer :post_category_id, null: false           # 質問・意見のカテゴリー（ActiveHash）
      t.text :goal                                       # 実現したいこと
      t.text :attempts                                   # 試したこと・調べたこと
      t.text :source_code                                # 該当ソースコード
      t.string :image                                    # 参考画像

      t.timestamps
    end
  end
end
