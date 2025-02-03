#!/bin/bash

#This script is used to run a command with specific arguments. 
# Currently used to create a new minikube profile and use it
init_profile() {
    export MINIKUBE_PROFILE=$1
    cp $HOME/.kube/config $HOME/.kube/config-$MINIKUBE_PROFILE
    export KUBECONFIG=$HOME/.kube/config-$MINIKUBE_PROFILE
    minikube start --cpu 8 --memory 6144
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
apply_this_helper(){
    cp ./.helper.sh $HOME/.helper.sh
    source $HOME/.helper.sh
}

copy_exo_bins_to_gsite(){
    pushd $EXO_REPO/SatOS_BIN/truetwin_software/janus2_truetwin
    docker cp ./ e2e-gsite-1:/tmp/bins
    docker exec e2e-gsite-1 bash -c "chmod +x /tmp/bins/*"
    popd
}

main() {
    args=${@:3}
    echo "Arguments to main : $args"
    echo "Environment file : $1"
    source $1 
    echo "Function to run : $2"
    $2 $args
}
echo "Script arguments : $@"
main $@