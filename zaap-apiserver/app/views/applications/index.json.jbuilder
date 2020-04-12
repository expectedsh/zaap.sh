json.applications @applications do |application|
  json.partial! 'applications/application', application: application
end
