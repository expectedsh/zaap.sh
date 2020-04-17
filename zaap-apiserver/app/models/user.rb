# == Schema Information
#
# Table name: users
#
#  id              :uuid             not null, primary key
#  email           :string           not null
#  first_name      :string           not null
#  password_digest :string           not null
#  scheduler_token :uuid             not null
#  scheduler_url   :string
#  created_at      :datetime         not null
#  updated_at      :datetime         not null
#
# Indexes
#
#  index_users_on_email            (email) UNIQUE
#  index_users_on_scheduler_token  (scheduler_token) UNIQUE
#
class User < ApplicationRecord
  before_save :downcase_email

  validates :email, presence: true, uniqueness: true
  validates :first_name, presence: true
  validate :scheduler_url, :check_scheduler_connection, on: :update
  has_secure_password

  has_many :applications, dependent: :destroy

  def issue_token
    JWT.encode ({ user_id: id, exp: 24.hours.from_now.to_i }),
               Rails.application.secrets.secret_key_base
  end

  def scheduler_connection
    Protocol::Scheduler::Stub.new scheduler_url, :this_channel_is_insecure
  end

  def check_scheduler_connection
    req = Protocol::TestConnectionRequest.new token: scheduler_token
    res = scheduler_connection.test_connection req
    errors.add :scheduler_url, 'invalid scheduler token' unless res.ok
  rescue StandardError
    errors.add :scheduler_url, 'connection failed'
  end

  private

  def downcase_email
    email.downcase!
  end
end
