#!/bin/bash


receive()
{
    var=`echo "$1" | base64 --decode`
    echo "$var"
}

while [ 1 ]
do
    read -p "Enter text: " var
    var=`receive $var`
    echo "Entered text : $var "
    sleep 2
done
