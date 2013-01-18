#!/bin/bash
#
# Copied from the blog post Profiling Go programs
# http://blog.golang.org/2011/06/profiling-go-programs.html
#
# On Mac OSX you have to install gnu-time (via Macports or Homebrew).
#
os=$(uname)
if [[ "$os" == 'Darwin' ]]; then
  echo user-time sys-time real-time max-mem cmd...
  gtime -f '%Uu %Ss %er %MkB %C' "$@"
else
  echo user-time sys-time real-time max-mem cmd...
  /usr/bin/time -f '%Uu %Ss %er %MkB %C' "$@"
fi
