#!/bin/bash
for ((i=1; i <= $2 ; i++))
do
  $GOPATH/bin/cmd $GOPATH/src/github.com/mna/specter/cmd/examples/$1.vm
done
