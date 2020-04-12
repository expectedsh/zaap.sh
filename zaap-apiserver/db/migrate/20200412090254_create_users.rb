class CreateUsers < ActiveRecord::Migration[6.0]
  def change
    create_table :users, id: :uuid do |t|
      t.string :email, null: false
      t.string :password_digest, null: false
      t.string :first_name, null: false

      t.timestamps
    end

    add_index :users, :email, unique: true
  end
end
