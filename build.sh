#!/bin/bash

setGOPATH() {
    SET_GOPATH=""
    while [ "$SET_GOPATH" == "" ]
    do
        echo -n "Please set your GOPATH (default: $HOME/gocode):"
        read SET_GOPATH 
    done

    export GOPATH="$SET_GOPATH"

    echo -n "Would you like to add GOPATH=$GOPATH to your ~/.bash_profile? (y/n):"
    read ADD_TO_PROFILE 
    if [ "$ADD_TO_PROFILE" = "y" ]
    then
        echo "export GOPATH=$GOPATH" >> ~/.bash_profile
    fi
}

interactive() {
    if [ -z $GOPATH ] 
    then
        echo "Your GOPATH is empty!"
        setGOPATH
    else
        echo -n "Using $GOPATH, Would you like to set a new one? (y/n): "
        read SET_NEWONE
        if [ "$SET_NEWONE" = "y" ]
        then
            setGOPATH
        fi
    fi
}

if [ "$1" != "--non-interactive" ]
then
    interactive
fi

go get -v github.com/imosquera/uploadthis
go get -v github.com/axw/gocov/gocov
go get -v launchpad.net/gocheck

echo 
echo -n "The packages have been downloaded and installed here: "
echo -e "\033[32m $GOPATH/src/github.com/imosquera/uploadthis\033[0m"
echo -e "Have fun gophing around! -- From your friends at \033[32mShareThis\033[0m"
