#!/bin/sh

docker build -t image-magick . && docker run -v `pwd`/tests/:/tests -p 9191:8080 image-magick