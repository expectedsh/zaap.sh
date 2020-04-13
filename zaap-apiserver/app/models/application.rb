# == Schema Information
#
# Table name: applications
#
#  id          :uuid             not null, primary key
#  environment :json
#  image       :string           not null
#  name        :string           not null
#  state       :integer          default(0), not null
#  created_at  :datetime         not null
#  updated_at  :datetime         not null
#  user_id     :uuid             not null
#
# Indexes
#
#  index_applications_on_user_id  (user_id)
#
class Application < ApplicationRecord
  after_save :deploy

  validates :name, presence: true, length: { minimum: 2, maximum: 32 }
  validates :image, presence: true

  enum state: %i[unknown stopped starting running]

  belongs_to :user, dependent: :destroy

  def deploy
    payload = { scheduler_token: user.scheduler_token, application: self }.to_json

    conn = Bunny.new.tap(&:start)
    ch = conn.create_channel
    ch.direct('deployment', durable: true).publish payload, routing_key: "deployment-consumer-#{user.scheduler_token}"
    ch.close
    conn.close
  end
end
