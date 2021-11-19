# !/usr/bin/bash

mkdir /opt/asa_feeler
mv ../Scanner-API /opt/asa_feeler
cd /opt/asa_feeler/Scanner-API/

read -p "请输入服务启动IP：" ASA_FEELER_SERVER_IP
read -p "请输入服务启动端口：" ASA_FEELER_SERVER_PORT
sed -i "s/{ASA_FEELER_SERVER_IP}/${ASA_FEELER_SERVER_IP}/g" /opt/asa_feeler/Scanner-API/config/config.yaml
sed -i "s/{ASA_FEELER_SERVER_PORT}/${ASA_FEELER_SERVER_PORT}/g" /opt/asa_feeler/Scanner-API/config/config.yaml
go build -o scanner
nohup ./scanner > log.log &

