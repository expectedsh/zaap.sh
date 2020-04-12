class CreateApplications < ActiveRecord::Migration[6.0]
  def change
    create_table :applications, id: :uuid do |t|
      t.string :name, null: false
      t.string :image, null: false
      t.json :environment
      t.integer :state, null: false, default: 0
      t.belongs_to :user, null: false, type: :uuid, index: true

      t.timestamps
    end
  end
end
