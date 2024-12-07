class User < ApplicationRecord
  # Include default devise modules. Others available are:
  # :confirmable, :lockable, :timeoutable, :trackable and :omniauthable
  devise :database_authenticatable, :registerable,
         :recoverable, :rememberable, :validatable
  # アソシエーション
  has_many :posts

  extend ActiveHash::Associations::ActiveRecordExtensions
  belongs_to :user_experience

  validates :email, presence: true, uniqueness: true, format: { with: URI::MailTo::EMAIL_REGEXP, message: "が無効です" }
  validates :username, presence: true, uniqueness: true, length: { in: 3..15 }, 
                     format: { with: /\A[\p{Hiragana}\p{Katakana}\p{Han}\w]+\z/, 
                               message: "は漢字、ひらがな、カタカナ、英数字（組み合わせでも単体でも可）で入力してください" }
  validates :user_experience_id, numericality: { other_than: 1, message: "が無効なカテゴリーです"}
  validates :password, presence: true, length: { minimum: 8 },
                     format: { with: /\A(?=.*[a-zA-Z])(?=.*\d)[a-zA-Z0-9]+\z/, message: "は半角英数字混合で8文字以上にしてください" },
                     confirmation: { case_sensitive: true }, on: :create
  validates :password_confirmation, presence: true, on: :create
  validates :affiliation, length: { maximum: 100 }, allow_blank: true
  validates :site_url, format: { with: URI::DEFAULT_PARSER.make_regexp, message: "URLが無効です" }, allow_blank: true
  validates :bio, length: { maximum: 500 }, allow_blank: true
  validates_uniqueness_of :email, :username
end