Rails.application.routes.draw do
  devise_for :users
  root 'posts#index'
  resources :users, only: [:show, :edit, :update]
  resources :posts do
    collection do
      get 'new_question', to: 'posts#new_question'
      get 'new_opinion', to: 'posts#new_opinion'
      get 'search'
    end
    resources :comments, only: :create
  end
end