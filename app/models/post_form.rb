class PostForm
  include ActiveModel::Model

  attr_accessor :user_id, :title, :content, :post_category_id, :goal, :attempts, :source_code, :image

  validates :user_id, presence: true
  validates :title, presence: true, length: { maximum: 100 }
  validates :content, presence: true, length: { maximum: 500 }
  validates :post_category_id, numericality: { other_than: 1, message: "が無効なカテゴリーです"}
  validates :source_code, length: { maximum: 1000 }, allow_blank: true
  validates :goal, length: { maximum: 300 }
  validates :attempts, length: { maximum: 300 }

  validate :validate_goal_and_attempts

  def save
    Post.create(title: title, content: content, post_category_id: post_category_id, 
                              source_code: source_code, goal: goal, attempts: attempts, user_id: user_id)
  end

  private

  def question?
    post_category_id.to_i == 2
  end

  def validate_goal_and_attempts
    if question?
      errors.add(:goal, "を入力してください") if goal.blank?
      errors.add(:attempts, "を入力してください") if attempts.blank?
    end
  end  
end