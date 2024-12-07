class PostsController < ApplicationController
  before_action :authenticate_user!, only: [:new, :new_question, :new_opinion, :create]
  def index
    @posts = Post.all
  end

  def new
  end

  def new_question
    @post = Post.new
  end

  def new_opinion
    @post = Post.new
  end

  def create
    @post = Post.new(post_params)
    if @post.save
      redirect_to root_path
    else
      if params[:post_type] == 'question'
        render :new_question, status: :unprocessable_entity
      elsif params[:post_type] == 'opinion'
        render :new_opinion, status: :unprocessable_entity
      else
        redirect_to new_post_path, alert: '投稿タイプが不正です。'
      end  
    end
  end

  private
   def post_params
     params.require(:post).permit(:title, :content, :post_category_id, 
                                  :goal, :attempts, :source_code, :image).merge(user_id: current_user.id)
   end
end