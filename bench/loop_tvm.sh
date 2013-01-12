#!/bin/bash
for ((i=1; i <= $2 ; i++))
do
  ~/subwrkspc/tinyvm/bin/tvmi ~/subwrkspc/tinyvm/programs/$1.vm >> /dev/null
done
