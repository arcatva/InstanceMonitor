#!/bin/bash

#This script is to help user install the python-based collector, i.e. influxdb-client & psutil in the target monitored node

echo "============START INSTALLATION============"

OS_VERSION=$(cat /etc/os-release | grep -e "^NAME=" | tr -d "NAME=\"")

if [[ ${OS_VERSION} == "Ubuntu" ]]
then
    echo "Current OS is Ubuntu, OK"
elif [[ ${OS_VERSION} != "Ubuntu" ]]
then
    echo "Current OS is NOT Ubuntu! Unsupported script!"
	exit
fi

sudo apt-get update

if [[ $? -ne 0 ]]

then
        echo "============INSTALL FAILED============"
        exit
fi

sudo apt-get install -y python3 python3-pip 

if [[ $? -ne 0 ]]

then
        echo "============INSTALL FAILED============"
        exit
fi

pip3 install psutil influxdb-client


if [[ $? -ne 0 ]]

then
	echo "============INSTALL FAILED============"
	exit
else 
	echo "============INSTALL COMPLETED============"
	python3 --version
	pip3 --version
	pip3 list | grep psutil
	pip3 list | grep influxdb-client

fi

python3 collector.py