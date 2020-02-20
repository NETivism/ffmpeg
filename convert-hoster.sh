#!/bin/bash
OUT="o_${1%%.*}.mp4"
docker run -v /var/www/sites/ffmpeg/input:/tmp/ffmpegi -v /var/www/sites/ffmpeg/output:/tmp/ffmpego --rm  netivism/ffmpeg \
  -stats \
  -i /tmp/ffmpegi/$1 \
  -vf 'scale=trunc(oh*a/2)*2:720' \
  -threads '4' -tune 'zerolatency' \
  -x264opts 'bitrate=2012:vbv-maxrate=2012:vbv-bufsize=2400' \
  -strict experimental \
  -vcodec 'h264' \
  -acodec 'aac' \
  -ac '2' \
  -b:a 220k \
  -y \
  /tmp/ffmpego/$OUT
