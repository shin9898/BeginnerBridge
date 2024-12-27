FactoryBot.define do
  factory :post_tag_relation do
    association :post
    association :tag
  end
end
