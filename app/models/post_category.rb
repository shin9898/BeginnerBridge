class PostCategory < ActiveHash::Base
  self.data = [
    { id: 1, name: '選択してください' },
    { id: 2, name: '質問' },
    { id: 3, name: '意見交換' },
    { id: 4, name: 'その他質問'}
  ]

  include ActiveHash::Associations
  has_many :posts
end