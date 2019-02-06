#!/bin/sh
git pull && make && echo "BUILD COMPLETE"
env PORT=10100 bin/manense