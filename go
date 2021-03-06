#!/bin/sh

find /public/greg/nwr -name 'NWR-*.raw' | xargs rm &
./rec -s - | \
./demux \
    "./splitter \
        \"./streamer -c wwf91.mp3.conf\" \
        \"./streamer -c wwf91.speex.conf\" \
        \"./streamer -c wwf91.vorbis.conf\" \
        \"./capture -d /public/file/0/nwr/wwf91 -l wwf91.capture.log -s \\\"./alert wwf91\\\" -\" \
        \"./log /public/greg/nwr/wwf91\"" \
    "./splitter \
        \"./streamer -c wxk27.mp3.conf\" \
        \"./streamer -c wxk27.speex.conf\" \
        \"./streamer -c wxk27.vorbis.conf\" \
        \"./capture -d /public/file/0/nwr/wxk27 -l wxk27.capture.log -s \\\"./alert wxk27\\\" -\" \
        \"./log /public/greg/nwr/wxk27\""
