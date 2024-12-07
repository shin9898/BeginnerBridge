class Post < ApplicationRecord
  belongs_to :user
  has_one_attached :image

  extend ActiveHash::Associations::ActiveRecordExtensions
  belongs_to :post_category

  validates :title, presence: true, length: { maximum: 100 }
  validates :content, presence: true, length: { maximum: 500 }
  validates :post_category_id, numericality: { other_than: 1, message: "が無効なカテゴリーです"}
  validates :source_code, length: { maximum: 1000 }, allow_blank: true
  validates :goal, length: { maximum: 300 }
  validates :attempts, length: { maximum: 300 }

  validate :validate_goal_and_attempts

  def question?
    post_category_id == 2
  end

  def opinion?
    post_category_id == 3
  end

  private

  def validate_goal_and_attempts
    if post_category_id == 2
      errors.add(:goal, "を入力してください") if goal.blank?
      errors.add(:attempts, "を入力してください") if attempts.blank?
    end
  end
end