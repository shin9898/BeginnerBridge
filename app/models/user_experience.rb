class UserExperience < ActiveHash::Base
  self.data = [
    { id: 1, name: '---' },
    { id: 2, name: '未経験者' },
    { id: 3, name: '初心者' },
    { id: 4, name: '初級者' },
    { id: 5, name: '中級者' },
    { id: 6, name: '実務経験1年〜5年' },
    { id: 7, name: '実務経験6年〜10年' },
    { id: 8, name: '実務経験10年以上' }
  ]

  include ActiveHash::Associations
  has_many :users
end