<?php

$json = '{"args":["-hide_banner","-loglevel error","-stats","-i /tmp/ffmpegi/abc.mov","-vf scale=trunc(oh*a/2)*2:720","-threads 4 -tune zerolatency","-x264opts bitrate=2012:vbv-maxrate=2012:vbv-bufsize=2400","-strict experimental","-vcodec h264","-acodec aac","-ac 2","-b:a 220k","-y","/tmp/ffmpego/abc.mp4"]}';
$o = json_decode($json);
if ($o) {
  var_export($o);
  $data_json = $json;
  $ch = curl_init('http://127.0.0.1:32468/execute');
  curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
  curl_setopt($ch, CURLOPT_CUSTOMREQUEST, "POST");
  curl_setopt($ch, CURLOPT_HTTPHEADER, array('Content-Type: application/json','Content-Length: ' . strlen($data_json)));
  curl_setopt($ch, CURLOPT_POSTFIELDS, $data_json);
  $response = curl_exec($ch);
  $httpcode = curl_getinfo($ch, CURLINFO_HTTP_CODE);
  print "\n==============\n";
  echo "HTTP status: ". $httpcode."\n\n";
  print_r($response);
  curl_close($ch);
}
