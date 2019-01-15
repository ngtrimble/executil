#!/bin/bash

x=1
while [ $x -le 3 ]
do
  echo "Sleeping $x..."
  >&2 echo "This is a message on stderr"
  sleep 1
  x=$(( $x + 1 ))
done

echo "finishing runsleep"
>&2 echo "finishing runsleep stderr message"
