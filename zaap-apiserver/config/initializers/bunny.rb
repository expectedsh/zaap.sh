require 'bunny'

AMQP = Bunny.new Rails.configuration.rabbitmq.url
AMQP.start
