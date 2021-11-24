#!/bin/bash
echo downloading dependencies
cd src ; go get . ; cd ..
echo dependencies installed, now run ./ymp3cli to start
