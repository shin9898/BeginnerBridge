class Tag < ApplicationRecord
  has_many :post_tag_relations
  has_many :posts, through: :post_tag_relations
  validates :tag_name,  uniqueness: true

  def self.ransackable_attributes(auth_object = nil)
    ["tag_name"]
  end
end
