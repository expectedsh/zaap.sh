# == Schema Information
#
# Table name: applications
#
#  id          :uuid             not null, primary key
#  environment :json
#  image       :string           not null
#  name        :string           not null
#  replicas    :integer          default(1), not null
#  state       :integer          default("unknown"), not null
#  created_at  :datetime         not null
#  updated_at  :datetime         not null
#  user_id     :uuid             not null
#
# Indexes
#
#  index_applications_on_user_id  (user_id)
#
class Application < ApplicationRecord
  after_save :request_deployment
  before_destroy :request_deletion

  validates :name, presence: true, length: { minimum: 2, maximum: 32 }
  validates :image, presence: true

  enum state: %i[unknown stopped starting running]

  belongs_to :user

  def to_grpc
    Protocol::Application.new id: id, name: name, image: image,
                          replicas: replicas,
                          environment: environment
  end

  def request_deployment
    req = Protocol::DeployApplicationRequest.new application: to_grpc
    user.scheduler_connection.deploy_application req
  end

  def request_deletion
    req = Protocol::DeleteApplicationRequest.new id: id
    user.scheduler_connection.delete_application req
  end
end
