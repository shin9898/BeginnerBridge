class CommentsController < ApplicationController
  def create
    @comment = Comment.new(comment_params)
    if @comment.save
      html = render_to_string(partial: 'comments/comment', locals: { comment: @comment })
      ActionCable.server.broadcast "comment_channel", { html: html }
    end
  end
  
  private

  def comment_params
    params.require(:comment).permit(:content).merge(user_id: current_user.id, post_id: params[:post_id])
  end
end
