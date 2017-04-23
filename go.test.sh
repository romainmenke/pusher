#!/usr/bin/env bash

set -e
echo "" > coverage.txt

if [[ $TRAVIS_GO_VERSION == 1.6* ]]
then

  for d in $(go list ./... | grep -v vendor | grep -v http2 | grep -v example); do
      go test -v -race -coverprofile=profile.out -covermode=atomic $d
      if [ -f profile.out ]; then
          cat profile.out >> coverage.txt
          rm profile.out
      fi
  done

elif [[ $TRAVIS_GO_VERSION == 1.7* ]]
then

  for d in $(go list ./... | grep -v vendor | grep -v http2 | grep -v example); do
      go test -v -race -coverprofile=profile.out -covermode=atomic $d
      if [ -f profile.out ]; then
          cat profile.out >> coverage.txt
          rm profile.out
      fi
  done

else

  for d in $(go list ./... | grep -v vendor); do
      go test -race $d
  done

  for d in $(go list ./... | grep -v vendor | grep -v http2 | grep -v example); do
      go test -v -race -coverprofile=profile.out -covermode=atomic $d
      if [ -f profile.out ]; then
          cat profile.out >> coverage.txt
          rm profile.out
      fi
  done

fi
