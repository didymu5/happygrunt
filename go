#!/bin/bash

function hexo ()
{
	echo 'hexo days are better'
}
function fun ()
{ # A somewhat more complex function.
  i=0
  REPEATS=30

  echo
  echo "And now the fun really begins."
  echo

  sleep $JUST_A_SECOND    # Hey, wait a second!
  while [ $i -lt $REPEATS ]
  do
    echo "----------FUNCTIONS---------->"
    echo "<------------ARE-------------"
    echo "<------------FUN------------>"
    echo
    let "i+=1"
  done
}
if [$# -eq 1]
	then echo 'running' $1
fi
$1
