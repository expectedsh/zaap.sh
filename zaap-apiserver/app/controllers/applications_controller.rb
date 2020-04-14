class ApplicationsController < ApplicationController
  before_action :authenticate_user!
  before_action :current_application!, except: %i[index create]

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

  def update
    return if @current_application.update application_params

    render status: :unprocessable_entity, json: @current_application.errors
  end

  def deploy
    @current_application.request_deployment
  rescue StandardError => e
    render status: :service_unavailable, json: { message: e.message }
  end

  private

  def application_params
    params.permit(%i[name image environment replicas])
  end

  def current_application!
    @current_application = Application.joins(:user).find_by(id: params[:id], user_id: @current_user.id)
    head 404 unless @current_application
  end
end
