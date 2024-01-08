#!/bin/bash

pid=-1
input=0

send()
{
    v=`echo "$@"|base64`
    echo $v>input
}

main()
{
    mkfifo input 
    ./loop.sh < input &
    pid=$!
    v="Text 
    12345
    " 
    send "$v"
    # vb=`echo "$v"|base64 `
    # echo $vb > input
    sleep 5
    send "Text 109876" 
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

