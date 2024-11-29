class UserExperience < ActiveHash::Base
  self.data = [
    { id: 1, name: '---', description: '選択してください' },
    { id: 2, name: '未経験者', description: 'これからプログラミングを学び始めたいもしくは学び始めたばかりの方' },
    { id: 3, name: '初心者', description: '基本的な文法や簡単なプログラムが理解できる方' },
    { id: 4, name: '初級者', description: '小規模なプロジェクトを自力で構築できる方' },
    { id: 5, name: '中級者', description: '複数人での開発経験があり、フレームワークを使いこなせる方' },
    { id: 6, name: '実務経験1年〜5年' },
    { id: 7, name: '実務経験6年〜10年' },
    { id: 8, name: '実務経験10年以上' }
  ]

  include ActiveHash::Associations
  has_many :users
end