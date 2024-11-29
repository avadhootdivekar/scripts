#!/bin/bash

#This script is used to run a command with specific arguments. 
# Currently used to create a new minikube profile and use it
init_profile() {
    export MINIKUBE_PROFILE=$1
    cp $HOME/.kube/config $HOME/.kube/config-$MINIKUBE_PROFILE
    export KUBECONFIG=$HOME/.kube/config-$MINIKUBE_PROFILE
    minikube start
}
rm_container(){
    __cont=$1
    docker stop $__cont
    docker rm $__cont
}
use_profile() {
    echo "Arguments : $@"
    export MINIKUBE_PROFILE=$1
    export KUBECONFIG=$HOME/.kube/config-$MINIKUBE_PROFILE
}
minikube_setup(){
    pushd $CODE/PC-playground/
    ./minikube-setup.sh $@
    popd
}

help(){
    echo "Usage : "
    echo "init_profile <profile_name> : Create a new minikube profile and use it"
    echo "use_profile <profile_name> : Use an existing minikube profile"
}
main() {
    args=${@:2}
    echo "Arguments (2:) to main : $args"

    $1 $args
}
echo "Script arguments : $@"
main $@