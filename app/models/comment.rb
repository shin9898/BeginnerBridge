class Comment < ApplicationRecord
  before_save :sanitize_content

  belongs_to :post
  belongs_to :user

  validates :content, presence: true, length: { maximum: 500 }

  private

  def sanitize_content
    self.content = content.strip.gsub(/\n+/, "\n")
  end

end
