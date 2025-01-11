class PostsController < ApplicationController
  before_action :authenticate_user!, only: [ :new, :new_question, :new_opinion, :new_other_question, :create, :edit, :update, :destroy ]
  before_action :set_post, only: [:show, :edit, :update, :destroy]
  before_action :check_user_permission, only: [:edit, :update, :destroy]

  def index
    if params[:category] == "質問"
      @posts = @q.result.where(post_category_id: [2, 4]).order(created_at: :desc)
    elsif params[:category] == "意見交換"
      @posts = @q.result.where(post_category_id: 3).order(created_at: :desc)
    else
      @posts = @q.result.order(created_at: :desc)
    end
  end

  def new
  end

  def new_question
    @post_form = PostForm.new
  end

  def new_other_question
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
      elsif params[:post_type] == 'other_question'
        render :new_other_question, status: :unprocessable_entity
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

  def search_tags
    return nil if params[:keyword] == ""
    tag = Tag.where(['tag_name LIKE ?', "%#{params[:keyword]}%"] )
    render json:{ keyword: tag }
  end

  def search
    keyword = params.dig(:q, :title_or_tags_tag_name_or_user_username_cont) || ""
    @posts = @q.result.includes(:tags, :user).distinct
    @posts = case params[:sort_order]
             when 'new'
              @posts.order(created_at: :desc)
             when 'old'
              @posts.order(created_at: :asc)
             else
              @posts
             end
             
    @keyword = keyword
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