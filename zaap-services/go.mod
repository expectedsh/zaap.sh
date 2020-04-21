module github.com/expected.sh/zaap.sh/zaap-services

go 1.14

replace github.com/expected.sh/zaap.sh/zaap-scheduler => ../zaap-scheduler

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/expected.sh/zaap.sh/zaap-scheduler v0.0.0-00010101000000-000000000000
	github.com/go-chi/chi v4.1.1+incompatible
	github.com/go-chi/cors v1.1.1
	github.com/go-ozzo/ozzo-validation/v4 v4.1.0
	github.com/jinzhu/gorm v1.9.12
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/lib/pq v1.3.0
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.5.0
	github.com/streadway/amqp v0.0.0-20200108173154-1c71cc93ed71
	google.golang.org/grpc v1.28.1
)
