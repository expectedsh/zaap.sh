class ApplicationController < ActionController::API
  protected

  def user_from_token(token)
    payload = JWT.decode token, Rails.application.secrets.secret_key_base
    User.find_by! id: payload[0]['user_id']
  end

  def authenticate_user
    token = request.headers['Authorization']&.split(' ')&.last
    return head :forbidden unless token

    @current_user = user_from_token token
  rescue JWT::DecodeError, ActiveRecord::RecordNotFound
    head :forbidden
  end

  def authenticate_user_from_params
    token = params[:token]
    return head :forbidden unless token

    @current_user = user_from_token token
  rescue JWT::DecodeError, ActiveRecord::RecordNotFound
    head :forbidden
  end
end
