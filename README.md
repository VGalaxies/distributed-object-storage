# chapter2

## rabbitmq

```
install rabbitmq
sudo systemctl start rabbitmq.service
bat /etc/rabbitmq/rabbitmq-env.conf
sudo rabbitmq-plugins enable rabbitmq_management
wget localhost:15672/cli/rabbitmqadmin
sudo rabbitmqctl add_user root root
python3 rabbitmqadmin declare exchange name=apiServers type=fanout
python3 rabbitmqadmin declare exchange name=dataServers type=fanout
sudo rabbitmqctl set_permissions -p / root ".*" ".*" ".*"
```
修改 `NODENAME=localhost`

初始用户名和密码均为 guest

## test

```
./setup.sh
http PUT localhost:11000/objects/test name=vgalaxy
http localhost:11001/objects/test
http localhost:11001/locate/test
pkill apiServer
pkill dataServer
```