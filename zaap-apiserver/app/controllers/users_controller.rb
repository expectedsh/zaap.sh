class UsersController < ApplicationController
  def create
    @user = User.new registration_params
    unless @user.save
      render status: :unprocessable_entity, json: @user.errors
      return
    end
    @token = @user.issue_token
    render status: :created
  end

  private

  def registration_params
    params.permit %i[first_name email password]
  end
end
