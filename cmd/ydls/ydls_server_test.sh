#!/bin/bash
set -ex

TEMPDIR=`mktemp -d`
trap "rm -rf $TEMPDIR" EXIT
cd "$TEMPDIR"

ydls -server -listen :1234 -config "$CONFIG" &
sleep 1

curl -OJ "http://0:1234/mp3+1s/https://www.youtube.com/watch?v=C0DPdy98e4c"
ffprobe -show_format -hide_banner -i "TEST VIDEO.mp3" 2>&1 | grep format_name=mp3

curl -OJ "http://0:1234/https://www.youtube.com/watch?v=C0DPdy98e4c"
ffprobe -show_format -hide_banner -i "TEST VIDEO.mp4" 2>&1 | grep format_name=mov

kill %1
