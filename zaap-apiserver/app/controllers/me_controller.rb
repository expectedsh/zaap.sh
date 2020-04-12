class MeController < ApplicationController
  before_action :authenticate_user!

  def update
    @current_user.update user_params
  end

  private

  def user_params
    params.permit(%i[first_name email])
  end
end
