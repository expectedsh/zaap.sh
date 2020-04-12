class ApplicationController < ActionController::API
  protected

  def authenticate_user!
    token = request.headers['Authorization']&.split(' ')&.last
    return head 403 unless token

    payload = JWT.decode token, Rails.application.secrets.secret_key_base
    @current_user = User.find_by! id: payload[0]['user_id']
  rescue JWT::DecodeError, ActiveRecord::RecordNotFound
    head 403
  end
end
