class Users::RegistrationsController < Devise::RegistrationsController
  before_action :configure_permitted_parameters

  respond_to :html

  protected

  def configure_permitted_parameters
    devise_parameter_sanitizer.permit(:sign_up) do |user|
      user.permit(%i[first_name email password])
    end
  end
end
