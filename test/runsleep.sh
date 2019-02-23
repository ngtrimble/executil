#!/bin/bash

x=1
while [ $x -le 3 ]
do
  echo "Sleeping $x..."
  sleep 1
  x=$(( $x + 1 ))
done

echo "finishing runsleep"
