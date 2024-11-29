require 'rails_helper'

RSpec.describe User, type: :model do
  before do
    @user = FactoryBot.build(:user)
  end

  describe 'ユーザー新規登録' do
    context '新規登録できる場合' do
      it "usernameとemail、passwordとpassword_confirmationが存在し、user_experience_idが1以外なら登録できる" do
        expect(@user).to be_valid
      end
    end
    context '新規登録できない場合' do
      it "usernameが空では登録できない" do
        @user.username = ''
        @user.valid?
        expect(@user.errors.full_messages).to include('ユーザーネームを入力してください')
      end
      it 'usernameが2文字以下なら登録できない' do
        @user.username = 'aa'
        @user.valid?
        expect(@user.errors.full_messages).to include('ユーザーネームは3文字以上で入力してください')
      end
      it 'usernameが16文字以上なら登録できない' do
        @user.username = 'aabbccddeeffgghh'
        @user.valid?
        expect(@user.errors.full_messages).to include('ユーザーネームは15文字以内で入力してください')
      end
      it 'usernameが重複してたら登録できない' do
        user1 = FactoryBot.create(:user, username: 'taro')
        @user.username = 'taro'
        @user.valid?
        expect(@user.errors.full_messages).to include('ユーザーネームはすでに存在します')
      end
      it "emailが空では登録できない" do
        @user.email = ''
        @user.valid?
        expect(@user.errors.full_messages).to include('Eメールを入力してください')
      end
      it 'emailが@を含む形式でないと登録できない' do
        @user.email = 'hogehoge'
        @user.valid?
        expect(@user.errors.full_messages).to include('Eメールは不正な値です', 'Eメールが無効です')
      end
      it 'emailが重複していたら登録できない' do
        user1 = FactoryBot.create(:user, email: 'hoge@hoge.jp')
        @user.email = 'hoge@hoge.jp'
        @user.valid?
        expect(@user.errors.full_messages).to include('Eメールはすでに存在します')
      end
      it "passwordが空では登録できない" do
        @user.password = ''
        @user.valid?
        expect(@user.errors.full_messages).to include('パスワードを入力してください')
      end
      it 'passwordが7文字以下では登録できない' do
        @user.password = '1234aaa'
        @user.valid?
        expect(@user.errors.full_messages).to include('パスワードは8文字以上で入力してください')
      end
      it 'passwordが英数字混合でないと登録できない' do
        @user.password = '12345678'
        @user.valid?
        expect(@user.errors.full_messages).to include('パスワードは半角英数字混合で8文字以上にしてください')
      end
      it 'passwordが全角英数字では登録できない' do
        @user.password = '１２３４ＡＢＣＤ'
        @user.valid?
        expect(@user.errors.full_messages).to include('パスワードは半角英数字混合で8文字以上にしてください')
      end
      it 'passwordとpassword_confirmationが一致していないと登録できない' do
        @user.password = '1111aaaa'
        @user.password_confirmation = '1234aaaa'
        @user.valid?
        expect(@user.errors.full_messages).to include('パスワード（確認用）とパスワードの入力が一致しません')
      end
      it 'user_experience_idが1なら登録できない' do
        @user.user_experience_id = 1
        @user.valid?
        expect(@user.errors.full_messages).to include('経験レベルカテゴリーが無効なカテゴリです')
      end
    end
  end
end