#!/usr/bin/env bash

docker run -it -v $PWD:/app -w /app personcentredcare_web bundle update
