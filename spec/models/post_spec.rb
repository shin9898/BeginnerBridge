require 'rails_helper'

RSpec.describe Post, type: :model do
  before do
    @post = FactoryBot.build(:post)
  end
  describe '質問・意見交換投稿の保存' do
    context '質問・意見交換が投稿できる場合' do
      it 'title,content,post_category_id,goal,attempts,source_code,imageを投稿できる' do
        expect(@post).to be_valid
      end
      it 'goalが空でも投稿できる' do
        @post.goal = ""
        expect(@post).to be_valid
      end
      it 'attemptsが空でも投稿できる' do
        @post.attempts = ""
        expect(@post).to be_valid
      end
      it 'source_codeが空でも投稿できる' do
        @post.source_code = ""
        expect(@post).to be_valid
      end
      it 'imageが空でも投稿できる' do
        @post.image = nil
        expect(@post).to be_valid
      end
    end

    context '質問・意見交換が投稿できない場合' do
      it 'titleが空では投稿できない' do
        @post.title = ""
        @post.valid?
        expect(@post.errors.full_messages).to include("タイトルを入力してください")
      end
      it 'contentが空では投稿できない' do
        @post.content = ""
        @post.valid?
        expect(@post.errors.full_messages).to include("内容を入力してください")
      end
      it 'post_category_idが空では投稿できない' do
        @post.post_category_id = nil
        @post.valid?
        expect(@post.errors.full_messages).to include("Post categoryが無効なカテゴリーです")
      end
      it 'ユーザーが紐づいてなければ投稿できない' do
        @post.user = nil
        @post.valid?
        expect(@post.errors.full_messages).to include("Userを入力してください")
      end
    end
  end
end