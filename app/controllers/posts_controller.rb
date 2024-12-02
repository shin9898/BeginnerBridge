class PostsController < ApplicationController
  before_action :authenticate_user!, only: [:new, :new_question, :new_opinion, :create]
  def index
    @posts = Post.all
  end

  # 質問 or 意見交換の選択ページ
  def new
    # 「質問する」「意見交換する」ボタンを表示する画面
  end

  def new_question
    @post = Post.new
    # 質問投稿用のフォームを表示するテンプレートをレンダリング
  end

  def new_opinion
    @post = Post.new
    # 意見交換用のフォームを表示するテンプレートをレンダリング
  end

  def create
    @post = Post.new(post_params)
    if @post.save
      redirect_to root_path
    else
      render :new, status: :unprocessable_entity
    end
  end

  private
   def post_params
     params.require(:post).permit(:title, :content, :post_category_id, 
                                  :goal, :attempts, :source_code, :image).merge(user_id: current_user.id)
   end
end