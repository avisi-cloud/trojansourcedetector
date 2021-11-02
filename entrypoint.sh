#!/bin/sh

if [ -n "$1" ]; then
    exec /trojansourcedetector -config $1
else
    exec /trojansourcedetector
fi
