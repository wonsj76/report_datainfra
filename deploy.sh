#! /bin/sh
rm -rf ./report_datainfra
rm -rf ./logs/*
git pull
go get -u github.com/ndcinfra/report_datainfra
killall -9 ./report_datainfra
go build
nohup ./report_datainfra &
cd logs


