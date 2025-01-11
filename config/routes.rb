Rails.application.routes.draw do
  devise_for :users, controllers: {
   omniauth_callbacks: 'users/omniauth_callbacks',
   registrations: 'users/registrations'
 }
  root 'posts#index'
  resources :users, only: [:show, :edit, :update] do
    resource :relationships, only: [:create, :destroy]
    resources :relationships, only: [:index]
  end

  resources :posts do
    collection do
      get 'new_question', to: 'posts#new_question'
      get 'new_other_question' , to: 'posts#new_other_question'
      get 'new_opinion', to: 'posts#new_opinion'
      get 'search_tags', to: 'posts#search_tags'
      get 'search'
    end
    resources :comments, only: :create
    resource :likes, only: [:create, :destroy]
  end
end