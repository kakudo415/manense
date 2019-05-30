#!/bin/sh
cd `dirname $0`
git pull && make && echo "BUILD COMPLETE"
env PORT=10100 bin/manense
