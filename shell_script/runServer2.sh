#!/bin/bash

npm install -g http-server
cd ~/my_website

start /b http-server -a 0.0.0.0 -p 8000
