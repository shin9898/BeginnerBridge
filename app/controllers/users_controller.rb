class UsersController < ApplicationController
    before_action :authenticate_user!, except: [:show]

    def show
      @user = User.find(params[:id])
      @username = @user.username
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
  
    def user_params
      params.require(:user).permit(:username, :email, :affiliation, :site_url, :bio)
    end
end
