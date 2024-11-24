class Genre < ActiveHash::Base
  self.data = [
    { id: 1, name: '---' },
    { id: 2, name: '未経験者' },
    { id: 3, name: '経験者' },
  ]

  include ActiveHash::Associations
  has_many :users
end