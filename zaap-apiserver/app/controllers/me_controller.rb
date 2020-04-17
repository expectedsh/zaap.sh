class MeController < ApplicationController
  before_action :authenticate_user

  def update
    return if @current_user.update user_params

    render status: :unprocessable_entity, json: @current_user.errors
  end

  private

  def user_params
    params.permit %i[first_name email scheduler_url]
  end
end
