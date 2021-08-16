#!/bin/bash

go build -o mise.exe cmd/server/*.go && ./mise -p=6379 -m=true