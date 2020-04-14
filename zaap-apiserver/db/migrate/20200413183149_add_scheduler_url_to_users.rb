class AddSchedulerUrlToUsers < ActiveRecord::Migration[6.0]
  def change
    add_column :users, :scheduler_url, :string
  end
end
