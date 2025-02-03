
export SEP:=\n====================================================================================================================================\n
export BUILD_VER:=1.01
export WORKSPACE_DIR:=/workspace
export CONTAINER_INSTANCE_NAME:=rust-1
export CONTAINER_IMAGE_NAME:=rust
export ENTRY_COMMAND:=sleep infinity

.phony: all clean get_host_dependancies build test get_build_shell 

all : 
	@echo ${SEP} all

clean : 
	@echo ${SEP} clean

get_host_dependancies:
	sudo apt-get update;
	sudo apt-get install -y docker.io 

test : 
	@echo ${SEP} test

get_build_shell : 
	@echo ${SEP} get_build_shell
	docker build -t ${CONTAINER_IMAGE_NAME}:${BUILD_VER} ./ ;
	xhost +local: ;
	# docker run -it --name="scripts-python" --network=host -w ${WORKSPACE_DIR} -v `pwd`:${WORKSPACE_DIR} python:${BUILD_VER} /bin/bash
	docker-compose up -d;
	docker exec -it ${CONTAINER_INSTANCE_NAME} bash 

cargobuild:
	cd ./hello-rust; cargo build ; cd ../''

build : 
	@echo ${SEP} build
	cd ./hello-rust/src; rustc main.rs; cd ../.. ; 

run:
	./hello-rust/target/debug/hello-rust