#!/bin/bash
ps aux | grep ffmpeg-router | grep "docker logs" | awk '{ print $2 }' | xargs kill
echo "" > log/error.log
echo "" > log/output.log
docker rm -f ffmpeg-router
docker run -d \
  --name ffmpeg-router \
  --restart=always \
  -v /var/www/sites/ffmpeg/input:/tmp/ffmpegi \
  -v /var/www/sites/ffmpeg/output:/tmp/ffmpego \
  -v $(PWD)/docker/ffmpeg-router /usr/local/bin/ffmpeg-router \
  --entrypoint /usr/local/bin/ffmpeg-router \
  -p 127.0.0.1:32468:32468 \
  jrottenberg/ffmpeg:7.0.2-ubuntu2204 bash
sleep 1

bash -c "docker logs --follow ffmpeg-router 2> log/error.log" &
bash -c "docker logs --follow ffmpeg-router 1> log/output.log" &

