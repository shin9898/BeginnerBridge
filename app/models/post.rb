class Post < ApplicationRecord
  belongs_to :user
  has_many :comments
  has_one_attached :image

  extend ActiveHash::Associations::ActiveRecordExtensions
  belongs_to :post_category

  def question?
    post_category_id.to_i == 2
  end

  def opinion?
    post_category_id.to_i == 3
  end
end