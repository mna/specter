#!/bin/bash
for ((i=1; i <= $2 ; i++))
do
  $SUBWRKSPC/tinyvm/bin/tvmi $GOPATH/src/github.com/mna/specter/cmd/examples/$1.vm
done
