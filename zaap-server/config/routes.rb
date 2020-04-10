Rails.application.routes.draw do
  devise_for :users,
             controllers: { registrations: 'users/registrations', sessions: 'users/sessions'  },
             skip: [:passwords, :confirmations, :registrations, :unlocks]
  devise_scope :user do
    get 'sign_up', to: 'users/registrations#new'
    post 'sign_up', to: 'users/registrations#create'
    get 'login', to: 'users/sessions#new'
    post 'login', to: 'users/sessions#create'
  end
  # For details on the DSL available within this file, see https://guides.rubyonrails.org/routing.html
end
