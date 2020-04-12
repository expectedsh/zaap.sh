json.application do
  json.partial! 'applications/application', application: @current_application
end
