for i in `seq 0 5`; do mkdir -p /tmp/$i/objects; done

export RABBITMQ_SERVER=amqp://root:root@localhost:5672

# go run in the corresponding module
cd dataServer
LISTEN_ADDRESS=localhost:10000 STORAGE_ROOT=/tmp/0 go run dataServer.go &
LISTEN_ADDRESS=localhost:10001 STORAGE_ROOT=/tmp/1 go run dataServer.go &
LISTEN_ADDRESS=localhost:10002 STORAGE_ROOT=/tmp/2 go run dataServer.go &
LISTEN_ADDRESS=localhost:10003 STORAGE_ROOT=/tmp/3 go run dataServer.go &
LISTEN_ADDRESS=localhost:10004 STORAGE_ROOT=/tmp/4 go run dataServer.go &
LISTEN_ADDRESS=localhost:10005 STORAGE_ROOT=/tmp/5 go run dataServer.go &

cd ../apiServer
LISTEN_ADDRESS=localhost:11000 go run apiServer.go &
LISTEN_ADDRESS=localhost:11001 go run apiServer.go &