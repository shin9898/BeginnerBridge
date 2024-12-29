class RelationshipsController < ApplicationController
  before_action :set_user, only: [:index]
  def index
    if params[:tab] == 'followers'
      @followers = @user.followers.page(params[:page]).per(20)
    else
      @followings = @user.followings.page(params[:page]).per(20)
    end
  end

  def create
    follow = current_user.active_relationships.new(follower_id: params[:user_id])
    follow.save
    redirect_to request.referer
  end

  def destroy
    follow = current_user.active_relationships.find_by(follower_id: params[:user_id])
    follow.destroy
    redirect_to request.referer
  end

  private
  def set_user
    @user = User.find_by(id: params[:user_id])
  redirect_to root_path, alert: 'ユーザーが見つかりません' unless @user
  end
end
