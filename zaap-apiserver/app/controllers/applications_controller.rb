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

  def deploy
    @current_application.deploy
    @current_application.state = :starting
    @current_application.tap(&:save!)
  end

  private

  def application_params
    params.permit(%i[name image environment])
  end

  def current_application!
    @current_application = Application.find_by id: params[:id], user_id: @current_user.id
    head 404 unless @current_application
  end
end
