class UsersController < ApplicationController
  before_action :authenticate_user!, except: [:show]
  before_action :redirect_if_signed_in

  def show
    @user = User.find(params[:id])
    @posts = case params[:filter]
    when "likes"
      @user.liked_posts
    when "comments"
       Post.joins(:comments).where(comments: { user_id: @user.id }).distinct
    else
      @user.posts
    end
  end

  def edit
    @user = current_user
  end
  
  def update
    @user = current_user
    if @user.update(user_params)
      redirect_to user_path(@user), notice: 'ユーザー情報を更新しました。'
    else
      render :edit, status: :unprocessable_entity
    end
  end

  private

  def redirect_if_signed_in
    redirect_to root_path if user_signed_in? && request.path == new_user_session_path
  end

  def user_params
    params.require(:user).permit(:username, :email, :user_experience_id, :affiliation, :site_url, :bio)
  end
end
