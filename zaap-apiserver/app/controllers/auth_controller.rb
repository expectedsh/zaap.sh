class AuthController < ApplicationController
  def login
    @user = User.find_by email: login_params[:email]&.downcase
    unless @user&.authenticate(params[:password])
      render status: :not_found, json: { message: 'Invalid email or password.' }
      return
    end
    @token = @user.issue_token
  end

  private

  def login_params
    params.permit %i[email password]
  end
end
