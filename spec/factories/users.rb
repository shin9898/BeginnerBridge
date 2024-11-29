FactoryBot.define do
  factory :user do
    username do
      # 日本語文字列と英数字を含むセットから3〜15文字をランダム生成
      charset = ('ぁ'..'ん').to_a + ('ァ'..'ン').to_a + ('一'..'龯').to_a + ('a'..'z').to_a + ('A'..'Z').to_a + ('0'..'9').to_a
      Array.new(Faker::Number.between(from: 3, to: 15)) { charset.sample }.join
    end
    email {Faker::Internet.email}
    password { Faker::Alphanumeric.alphanumeric(number: 8, min_alpha: 1, min_numeric: 1) }
    password_confirmation {password}
    user_experience_id { Faker::Number.between(from: 2, to: 8) }
  end
end