#!/bin/bash

pid=-1

main()
{
    mkfifo input 
    ./loop.sh < input &
    pid=$!
    echo "Text 12345" > input
    sleep 5
    echo "Text 109876" > input
    sleep 2 
}

exit()
{
    kill -9 $pid
    echo "Killed child."
}

trap 'exit' 9
trap 'exit' 2 EXIT

main
exit

