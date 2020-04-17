class CreateDeployments < ActiveRecord::Migration[6.0]
  def change
    create_table :deployments do |t|
      t.string :image
      t.json :environment
      t.integer :replicas

      t.timestamps
    end
  end
end
