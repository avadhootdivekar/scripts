
export PROJECT_ROOT = $(shell pwd)
export GO_WORKSPACE_DIR:=/usr/src/app
export WORKSPACE_DIR_HOST:=${PROJECT_ROOT}


.phony: all clean build

SP======================================================================================================================\n

all:
	@echo ${SP} all

clean:
	@echo ${SP} clean

build:
	@echo ${SP} build

shell:
	@echo ${SP} shell
	docker build -t build-env-scripts:0.1 ./  ;											\
	docker-compose up -d;		
	docker exec -it go_dev-env_1 bash ;		

env: 
	@echo ${SP} env
	sudo apt-get install -y docker.io docker-compose ;
	xhost +local:
	sudo groupadd docker || echo "Group Added";
	sudo usermod -aG docker ${USER}  || echo "User Added";					
	newgrp docker  || echo "Group refreshed";										

