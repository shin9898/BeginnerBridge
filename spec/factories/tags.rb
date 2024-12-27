FactoryBot.define do
  factory :tag do
    tag_name { Faker::Lorem.word }
  end
end
