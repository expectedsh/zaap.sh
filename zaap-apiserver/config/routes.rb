Rails.application.routes.draw do
  post 'auth/login'

  get 'me', to: 'me#index'
  put 'me', to: 'me#update'
  patch 'me', to: 'me#update'

  resources :users, except: %i[edit new]
  resources :applications, except: %i[edit new] do
    member do
      get 'logs'
      post 'deploy'
    end
  end
end

