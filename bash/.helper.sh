#!/bin/bash

#This script is used to run a command with specific arguments. 
# Currently used to create a new minikube profile and use it
init_profile() {
    export MINIKUBE_PROFILE=$1
    cp $HOME/.kube/config $HOME/.kube/config-$MINIKUBE_PROFILE
    export KUBECONFIG=$HOME/.kube/config-$MINIKUBE_PROFILE
    minikube start
}

use_profile() {
    echo "Arguments : $@"
    export MINIKUBE_PROFILE=$1
    export KUBECONFIG=$HOME/.kube/config-$MINIKUBE_PROFILE
}

main() {
    args=${@:2}
    echo "Arguments to main : $args"
    $1 $args
}
echo "Script arguments : $@"
main $@