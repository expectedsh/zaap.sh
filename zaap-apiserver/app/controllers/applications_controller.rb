class ApplicationsController < ApplicationController
  before_action :authenticate_user!

  def index
    @applications = @current_user.applications
  end

  def create
    @application = Application.new application_params
    @application.environment ||= {}
    @application.user = @current_user
    unless @application.save
      render status: :unprocessable_entity, json: @application.errors
      return
    end

    render status: :created
  end

  private

  def application_params
    params.permit(%i[name image environment])
  end
end
