class AddReplicasToApplications < ActiveRecord::Migration[6.0]
  def change
    add_column :applications, :replicas, :integer, null: false, default: 1
  end
end
