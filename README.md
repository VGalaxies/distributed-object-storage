# chapter3

## elasticsearch

```
yay -S elasticsearch
sudo elasticsearch-keystore create
sudo systemctl start elasticsearch.service
http PUT localhost:9200/metadata < mappings.json
```

## test

```
./setup.sh
cd test
./test.sh
cd ..
pkill apiServer
pkill dataServer
```