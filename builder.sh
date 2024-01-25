#!/bin/bash

docker build -t cluster-node .

docker image list

svc_names=('pikachu' 'squirtle' 'charmander')
for i in ${!svc_names[@]}
do
    portnum=$((5000+$i))
    echo "Name $i is ${svc_names[$i]}:${portnum}"
    docker run -dp 127.0.0.1:$portnum:$portnum --name ${svc_names[$i]} cluster-node -service_name=${svc_names[$i]} -port=$portnum
done

echo

#docker image rm -f cluster-node