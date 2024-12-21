class PostsController < ApplicationController
  before_action :authenticate_user!, only: [ :new, :new_question, :new_opinion, :create, :edit, :update, :destroy ]
  before_action :set_post, only: [:show, :edit, :update, :destroy]
  before_action :check_user_permission, only: [:destroy]
  def index
    if params[:category] == "質問"
      @posts = Post.where(post_category_id: 2).order(created_at: :desc)
    elsif params[:category] == "意見交換"
      @posts = Post.where(post_category_id: 3).order(created_at: :desc)
    else
      @posts = Post.all.order(created_at: :desc)
    end
  end

  def new
  end

  def new_question
    @post_form = PostForm.new
  end

  def new_opinion
    @post_form = PostForm.new
  end

  def create
    @post_form = PostForm.new(post_form_params)
    if @post_form.valid?
      @post_form.save
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
    @comment = Comment.new
    @comments = @post.comments.includes(:user)
  end

  def edit
    post_attributes = @post.attributes
    @post_form = PostForm.new(post_attributes)
    unless @post.user == current_user
      redirect_to post_path(@post) and return
    end
    @post_form.tag_name = @post.tags.first&.tag_name
  end

  def update
    @post_form = PostForm.new(post_form_params)
    @post_form.image ||= @post.image.blob
    if @post_form.valid? 
      @post_form.update(post_form_params, @post)
      redirect_to post_path(@post.id)
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

  def post_form_params
    params.require(:post_form).permit(:title, :content, :post_category_id,
                                  :goal, :attempts, :source_code, :tag_name, :image).merge(user_id: current_user.id)
  end

  def post_params
     params.require(:post).permit(:title, :content, :post_category_id, 
                                  :goal, :attempts, :source_code, :image).merge(user_id: current_user.id)
  end
end