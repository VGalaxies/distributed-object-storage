module dataServer

go 1.18

require lib v0.0.0

replace lib => ../lib

require github.com/streadway/amqp v1.0.0 // indirect
