class Comment < ApplicationRecord
  before_save :sanitize_content

  belongs_to :post
  belongs_to :user
  has_many :likes

  validates :content, presence: true, length: { maximum: 500 }

  def liked_by?(user)
    likes.where(user_id: user.id).exists?
  end

  private

  def sanitize_content
    self.content = content.strip.gsub(/\n+/, "\n")
  end

end
