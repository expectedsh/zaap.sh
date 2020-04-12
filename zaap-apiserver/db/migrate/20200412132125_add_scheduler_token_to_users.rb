class AddSchedulerTokenToUsers < ActiveRecord::Migration[6.0]
  def change
    add_column :users, :scheduler_token, :uuid, null: false, default: 'gen_random_uuid()'
    add_index :users, :scheduler_token, unique: true
  end
end
