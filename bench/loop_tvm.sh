#!/bin/sh
for ((i=1; i <= $2 ; i++))
do
  ~/subwrkspc/tinyvm/bin/tvmi $GOPATH/src/github.com/PuerkitoBio/specter/cmd/examples/$1.vm >> /dev/null
done
