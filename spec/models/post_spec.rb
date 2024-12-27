require 'rails_helper'

RSpec.describe Post, type: :model do
  before do
    @post = FactoryBot.build(:post)
    @tag = FactoryBot.create(:tag, tag_name: 'Ruby')
  end
  describe '質問・意見交換投稿の保存' do
    context '質問・意見交換が投稿できる場合' do
      it 'すべての情報が正しく入力されている場合保存できる' do
        expect(@post).to be_valid
      end
      it 'goalとattemptsが空でもpost_category_idが3なら保存できる' do
        @post.post_category_id = 3
        @post.goal = ''
        @post.attempts = ''
        expect(@post).to be_valid
      end
      it 'source_codeが空でも投稿できる' do
        @post.source_code = ''
        expect(@post).to be_valid
      end
      it 'imageが空でも投稿できる' do
        @post.image = nil
        expect(@post).to be_valid
      end
      it '投稿とタグを関連づけて投稿できる' do
        PostTagRelation.create(post: @post, tag: @tag)

        expect(@post.tags.count).to eq(1)
        expect(@post.tags.first.tag_name).to eq('Ruby')
      end
    end

    context '質問・意見交換が投稿できない場合' do
      it 'goalが空でpost_category_idが2の場合保存できない' do
        @post.post_category_id = 2
        @post.goal = ''
        @post.valid?
        expect(@post.errors.full_messages).to include("実現したいことを入力してください")
      end
      it 'attemptsが空でpost_category_idが2の場合保存できない' do
        @post.post_category_id = 2
        @post.attempts = ''
        @post.valid?
        expect(@post.errors.full_messages).to include("試したこと・調べたことを入力してください")
      end
      it 'titleが空では投稿できない' do
        @post.title = ''
        @post.valid?
        expect(@post.errors.full_messages).to include("タイトルを入力してください")
      end
      it 'contentが空では投稿できない' do
        @post.content = ''
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

#bundle exec rspec spec/models/post_spec.rb