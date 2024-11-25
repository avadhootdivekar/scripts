#!/bin/bash

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