#!/bin/bash
set -e

trap 'kill -HUP 0' EXIT;

# warm up
go test -run=NONE -bench=. ./link/... > link-new.bench
go test -run=NONE -bench=. ./link/... > link-new.bench

currentBranch=$(git rev-parse --abbrev-ref HEAD)

echo "bench current state"

go test -run=NONE -bench=. ./casper/... > casper-new.bench
go test -run=NONE -bench=. ./common/... > common-new.bench
go test -run=NONE -bench=. ./link/...   > link-new.bench
go test -run=NONE -bench=. ./parser/... > parser-new.bench
go test -run=NONE -bench=. ./rules/...  > rules-new.bench

git diff-index --quiet HEAD -- &&
{
  echo "checkout master"
  git checkout master
};

git diff-index --quiet HEAD -- ||
{
  echo "stash and checkout master"
  git stash
  stashed=true
  git checkout master
};

echo "bench master"
go test -run=NONE -bench=. ./casper/... > casper-old.bench
go test -run=NONE -bench=. ./common/... > common-old.bench
go test -run=NONE -bench=. ./link/...   > link-old.bench
go test -run=NONE -bench=. ./parser/... > parser-old.bench
go test -run=NONE -bench=. ./rules/...  > rules-old.bench

echo "go back to current state"
git checkout $currentBranch

if [ "$stashed" = true ]; then
  echo "pop stash"
  git stash pop
fi

echo "--- casper ---"
benchcmp casper-old.bench casper-new.bench || true
echo "--- common ---"
benchcmp common-old.bench common-new.bench || true
echo "--- link ---"
benchcmp link-old.bench   link-new.bench || true
echo "--- parser ---"
benchcmp parser-old.bench parser-new.bench || true
echo "--- rules ---"
benchcmp rules-old.bench  rules-new.bench || true
