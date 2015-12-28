#!/bin/bash

function hexo ()
{
	echo 'hexo days are better'
}
hexo() {
  echo "running Hexo"
  "$(npm bin)/hexo" "$@"
}

if [ "$(type -t "$1")" == "function" ]; then
  "$@"
else
  echo $1 "not enough"
fi