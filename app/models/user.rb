class User < ApplicationRecord
  # Include default devise modules. Others available are:
  # :confirmable, :lockable, :timeoutable, :trackable and :omniauthable
  devise :database_authenticatable, :registerable,
         :recoverable, :rememberable, :validatable
  # Deviseのバリデーション
  validates :email, presence: true, uniqueness: true, format: { with: URI::MailTo::EMAIL_REGEXP, message: "メールアドレスが無効です" }
  validates :username, presence: true, uniqueness: true, length: { in: 3..15 }, format: { with: /\A[ぁ-んァ-ン一-龯a-zA-Z0-9]+\z/, message: "ユーザー名は漢字、ひらがな、英数字（組み合わせでも単体でも可）で入力してください" }
  validates :user_experience_id, numericality: { other_than: 1, message: "無効なカテゴリです"}

  # パスワードと確認用パスワードのバリデーション
  validates :password, presence: true, length: { minimum: 8 }, format: { with: /\A[a-zA-Z0-9]+\z/, message: "パスワードは半角英数字で8文字以上にしてください" }, confirmation: { case_sensitive: true }
  validates :password_confirmation, presence: true, if: :password_required?
       
  # その他のフィールドに対するバリデーション
  validates :affiliation, length: { maximum: 100 }, allow_blank: true
  validates :site_url, format: { with: URI::DEFAULT_PARSER.make_regexp, message: "URLが無効です" }, allow_blank: true
  validates :bio, length: { maximum: 500 }, allow_blank: true

  # ユーザー名とメールアドレスの一意性
  validates_uniqueness_of :email, :username

  # パスワードのバリデーションが必要かどうかを判定するメソッド
  private

  def password_required?
    new_record? || password.present? || password_confirmation.present?
  end
end
