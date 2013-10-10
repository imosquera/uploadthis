#!/bin/bash

#this makes sure the script stops running if there are any failures
set -e

setGOPATH() {
    echo 
    echo "GOPATH is a workspace where your go files, libraries and projects will be placed."
    echo -e "\033[35mNOTE:\033[0m This must not be the same path as your Go installation. More info: http://golang.org/doc/code.html#GOPATH"
    echo -n "Please set your GOPATH (default: $HOME/gocode):"
    read SET_GOPATH 
    if [ "$SET_GOPATH" = "" ]
    then 
        SET_GOPATH="$HOME/gocode"
    fi
    export GOPATH="$SET_GOPATH"
    mkdir -p $GOPATH

   
    echo -n "Add GOPATH=$GOPATH to your ~/.bash_profile? [highly suggested] (y/n):"
    read ADD_TO_PROFILE 
    if [ "$ADD_TO_PROFILE" = "y" ]
    then
        echo "export GOPATH=$GOPATH" >> ~/.bash_profile
        echo "export PATH=$PATH:$GOPATH/bin" >> ~/.bash_profile
    fi
}

generateMocks() {
    set -x
    #we have to build the packages so we can use reflection to find them and generate the mocks
    go get ./...

    #lets start generating mocks
    mockgen -package=mocks -source=./execution/execution.go > util/mocks/execution_mocks.go
    mockgen -package=mocks -source=./commands/commands.go > util/mocks/commands_mocks.go
    mockgen -package=mocks -source=./hooks/prehooks.go > util/mocks/prehooks_mocks.go
    mockgen -package=mocks os FileInfo > util/mocks/os_mocks.go
    mockgen -package=mocks io Reader > util/mocks/io_mocks.go

    #we generate using reflection because it exports package names into the parameters
    #this is the preferred method to generate mocks
    mockgen -package=mocks github.com/imosquera/uploadthis/conf ConfigLoader > util/mocks/conf_mocks.go
    mockgen -package=mocks github.com/imosquera/uploadthis/util OSFile > util/mocks/conf_mocks.go
    mockgen -package=mocks github.com/imosquera/uploadthis/upload Uploader > util/mocks/upload_mocks.go
    set +x
    echo "Mocks have been generated"
}

usage() {
    cat << EOF
    usage: $0 options command

    This script manages this project

    OPTIONS:
       -h                     Show this message
       -n                     Non-interactive mode
EOF
}

interactive() {
    if [ -z $GOPATH ] 
    then
        echo
        echo -e '\033[33m'"Your GOPATH is not set!\033[0m";
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

build () {
    if [ "$NON_INTERACTIVE" == "true" ] 
    then
        interactive
    fi
    set -x
    go get -v github.com/imosquera/uploadthis
    go get -v github.com/axw/gocov/gocov
    go get -v code.google.com/p/gomock/gomock
    go get -v code.google.com/p/gomock/mockgen
    go get -v github.com/qur/withmock
    go get -v github.com/qur/withmock/mocktest
    go get -v launchpad.net/gocheck
    go get -v github.com/matm/gocov-html
    go get -v github.com/bradfitz/goimports
    set +x

    echo 
    echo -n "The packages have been downloaded and installed here: "
    echo -e "\033[32m $GOPATH/src/github.com/imosquera/uploadthis\033[0m"
    echo -e "Have fun gophing around! -- From your friends at \033[32mShareThis\033[0m"
}

NON_INTERACTIVE="true"
while getopts “h:n” OPTION
do
     case $OPTION in
         h)
             usage
             exit 1
             ;;
         n)
             NON_INTERACTIVE="false"
             ;;
         ?)
             usage
             exit
             ;;
     esac
done
shift $((OPTIND-1))
PROGRAM=$1

case "$1" in
        genmocks)
            generateMocks
            ;;
        *)
            build
            ;;
esac
