Rails.application.routes.draw do
  post 'auth/login'

  get 'me', to: 'me#index'
  put 'me', to: 'me#update'
  patch 'me', to: 'me#update'

  resources :users, except: %i[edit new]
end

