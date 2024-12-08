class PostsController < ApplicationController
  before_action :authenticate_user!, only: [ :new, :new_question, :new_opinion, :create, :edit, :update, :destroy ]
  before_action :set_post, only: [:show, :edit, :update, :destroy]
  before_action :check_user_permission, only: [:edit, :update, :destroy]
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
  end

  def edit
    unless @post.user == current_user
      redirect_to post_path(@post) and return
    end
  end

  def update
    if @post.update(post_params)
      redirect_to post_path(@post)
    else
      render :edit, status: :unprocessable_entity
    end
  end

  def destroy
    @post.destroy
    redirect_to root_path
  end

  private

  def set_post
    @post = Post.find(params[:id])
  end

  def check_user_permission
    redirect_to post_path(@post) unless @post.user == current_user
  end  

   def post_params
     params.require(:post).permit(:title, :content, :post_category_id, 
                                  :goal, :attempts, :source_code, :image).merge(user_id: current_user.id)
   end
end