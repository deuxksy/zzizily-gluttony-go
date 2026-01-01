#!/bin/bash

# Mac에서 Chrome을 원격 디버깅 모드로 실행
# 포트: 12222
# 프로필: /tmp/chrome_dev_profile (기존 크롬과 별개로 실행)

"/Applications/Google Chrome.app/Contents/MacOS/Google Chrome" \
  --remote-debugging-port=12222 \
  --user-data-dir="/tmp/chrome_dev_profile" \
  --no-first-run \
  --no-default-browser-check
