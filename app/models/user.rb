class User < ApplicationRecord
  # Include default devise modules. Others available are:
  # :confirmable, :lockable, :timeoutable, :trackable and :omniauthable
  devise :database_authenticatable, :registerable,
         :recoverable, :rememberable, :validatable, :omniauthable, omniauth_providers: [:facebook, :google_oauth2]
  # アソシエーション
  has_many :posts
  has_many :comments
  has_many :active_relationships, class_name: "Relationship", foreign_key: :following_id
  has_many :followings, through: :active_relationships, source: :follower

  has_many :passive_relationships, class_name: "Relationship", foreign_key: :follower_id
  has_many :followers, through: :passive_relationships, source: :following
  has_many :likes
  has_many :liked_posts, through: :likes, source: :post
  has_many :sns_credentials

  extend ActiveHash::Associations::ActiveRecordExtensions
  belongs_to :user_experience

  validates :email, presence: true, uniqueness: true, format: { with: URI::MailTo::EMAIL_REGEXP, message: "が無効です" }
  validates :username, presence: true, uniqueness: true, length: { in: 3..15 }, 
                     format: { with: /\A[\p{Hiragana}\p{Katakana}\p{Han}\w]+\z/, 
                               message: "は漢字、ひらがな、カタカナ、英数字（組み合わせでも単体でも可）で入力してください" }
  validates :user_experience_id, numericality: { other_than: 1, message: "が無効なカテゴリーです"}
  validates :password, presence: true, length: { minimum: 8 }, on: :create
  validates :password_confirmation, presence: true, on: :create
  validates :affiliation, length: { maximum: 100 }, allow_blank: true
  validates :site_url, format: { with: URI::DEFAULT_PARSER.make_regexp, message: "URLが無効です" }, allow_blank: true
  validates :bio, length: { maximum: 500 }, allow_blank: true
  validates_uniqueness_of :email, :username

  def self.from_omniauth(auth)
    sns = SnsCredential.where(provider: auth.provider, uid: auth.uid).first_or_create
    user = User.where(email: auth.info.email).first_or_initialize(
     username: auth.info.name,
       email: auth.info.email
   )

    if user.persisted?
      sns.user = user
      sns.save
    end
    { user: user, sns: sns }
  end

  def self.ransackable_attributes(auth_object = nil)
    ["username"]
  end

  def followed_by?(user)
    follower =  passive_relationships.find_by(following_id: user.id)
    return follower.present?
  end

  def following?(user)
    active_relationships.exists?(follower_id: user.id)
  end
end