# == Schema Information
#
# Table name: users
#
#  id              :uuid             not null, primary key
#  email           :string
#  first_name      :string
#  password_digest :string
#  scheduler_token :uuid             not null
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
  has_secure_password

  has_many :applications

  def issue_token
    JWT.encode ({ user_id: id, exp: 24.hours.from_now.to_i }),
               Rails.application.secrets.secret_key_base
  end

  private

  def downcase_email
    email.downcase!
  end
end
