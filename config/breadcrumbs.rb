crumb :root do
  link "トップ", root_path
end

crumb :post do |post|
  link post.title, post_path(post)
  parent :root
end

crumb :edit_post do |post|
  link "編集", edit_post_path(post)
  parent :post, post
end

crumb :new_post do
  link "投稿形式を選択する", new_post_path
  parent :root
end

crumb :new_question_post do
  link "プログラミングについて質問する", new_question_posts_path
  parent :new_post
end

crumb :new_other_question_post do
  link "その他について質問する", new_other_question_posts_path
  parent :new_post
end

crumb :new_opinion_post do
  link "意見交換する", new_opinion_posts_path
  parent :new_post
end

crumb :search do
  link "検索結果", search_posts_path
  parent :root
end

crumb :new_user_registration do
  link "新規登録する", new_user_registration_path
  parent :root
end

crumb :login do
  link "ログインする", new_user_session_path
  parent :root
end

crumb :mypage do |user|
  link "#{user.username}さんのマイページ", user_path(user)
  parent :root
end

crumb :index_relationships do |user|
  link "#{user.username}さんのフォロー・フォロワー一覧", user_relationships_path(user)
  parent :mypage, user
end

crumb :edit_user do |user|
  link "プロフィール編集", edit_user_registration_path(user)
  parent :mypage, user
end
