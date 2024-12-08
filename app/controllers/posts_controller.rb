class PostsController < ApplicationController
  before_action :authenticate_user!, only: [:new, :new_question, :new_opinion, :create, :edit, :update]
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
        redirect_to new_post_path
      end  
    end
  end

  def show
    @post = Post.find(params[:id])
  end

  def edit
    @post = Post.find(params[:id])
    redirect_to post_path(@post) unless @post.user == current_user
  end

  def update
    @post = Post.find(params[:id])

    unless @post.user == current_user
    redirect_to root_path and return
    end

    if @post.update(post_params)
      redirect_to post_path(@post), notice: '投稿を更新しました。'
    else
      flash.now[:alert] = '更新に失敗しました。'
      render :edit, status: :unprocessable_entity
    end
  end

  private
   def post_params
     params.require(:post).permit(:title, :content, :post_category_id, 
                                  :goal, :attempts, :source_code, :image).merge(user_id: current_user.id)
   end
end