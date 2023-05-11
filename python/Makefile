
export SEP:=\n====================================================================================================================================\n
export BUILD_VER:=1.01
export WORKSPACE_DIR:=/workspace

all : 
	@echo ${SEP} all

clean : 
	@echo ${SEP} clean

get_host_dependancies:
	sudo apt-get update;
	sudo apt-get install -y docker.io 

build : 
	@echo ${SEP} build

test : 
	@echo ${SEP} test

get_build_shell : 
	@echo ${SEP} get_build_shell
	docker build -t python:${BUILD_VER} ./ ;
	xhost +local: ;
	# docker run -it --name="scripts-python" --network=host -w ${WORKSPACE_DIR} -v `pwd`:${WORKSPACE_DIR} python:${BUILD_VER} /bin/bash
	docker-compose up -d;
	docker exec -it python_dev-env_1 bash 


build : 
	@echo ${SEP} build


