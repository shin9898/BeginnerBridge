class Post < ApplicationRecord
  belongs_to :user
  has_many :comments
  has_one_attached :image
  has_many :post_tag_relations
  has_many :tags, through: :post_tag_relations
  has_many :likes

  extend ActiveHash::Associations::ActiveRecordExtensions
  belongs_to :post_category

  def question?
    post_category_id.to_i == 2
  end

  def opinion?
    post_category_id.to_i == 3
  end

  def self.ransackable_attributes(auth_object = nil)
    ["title"]
  end

  def self.ransackable_associations(auth_object = nil)
    ["user", "tags"]
  end

  def liked_by?(user)
    likes.where(user_id: user.id).exists?
  end
end