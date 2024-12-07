FactoryBot.define do
  factory :post do
    title {Faker::Lorem.characters(number: 50)}
    content {Faker::Lorem.characters(number: 500)}
    post_category_id { Faker::Number.between(from: 2, to: 3) }
    goal {Faker::Lorem.characters(number: 300)}
    attempts {Faker::Lorem.characters(number: 300)}
    source_code {Faker::Lorem.characters(number: 1000)}
    association :user

    trait :opinion do
      post_category_id { 3 }
      goal { nil }
      attempts { nil }
    end
  end
end
