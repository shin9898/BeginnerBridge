class PostCategory < ActiveHash::Base
  self.data = [
    { id: 1, name: '選択してください', description: '選択してください' },
    { id: 2, name: '質問', description: '実現方法がわからない、エラーが解決できない' },
    { id: 3, name: '意見交換', description: 'プログラミングやITエンジニアについて、意見を募りたいテーマがある場合には「意見交換」を選択' }
  ]

  include ActiveHash::Associations
  has_many :posts
end