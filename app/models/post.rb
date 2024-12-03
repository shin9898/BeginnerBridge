class Post < ApplicationRecord
  belongs_to :user
  has_one_attached :image

  extend ActiveHash::Associations::ActiveRecordExtensions
  belongs_to :post_category

  validates :title, presence: true, length: { maximum: 100 }
  validates :content, presence: true, length: { maximum: 500 }
  validates :post_category_id, numericality: { other_than: 1, message: "が無効なカテゴリーです"}
  validates :goal, length: { maximum: 300 }, allow_blank: true
  validates :attempts, length: { maximum: 300 }, allow_blank: true
  validates :source_code, length: { maximum: 1000 }, allow_blank: true
end