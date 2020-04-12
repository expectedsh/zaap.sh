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
  validates :name, presence: true, length: { minimum: 2, maximum: 32 }
  validates :image, presence: true

  enum state: %i[unknown stopped starting running]

  belongs_to :user, dependent: :destroy
end
