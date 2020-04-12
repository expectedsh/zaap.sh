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
require 'test_helper'

class ApplicationTest < ActiveSupport::TestCase
  # test "the truth" do
  #   assert true
  # end
end
