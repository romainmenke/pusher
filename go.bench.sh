#!/bin/bash
set -e

currentBranch=$(git rev-parse --abbrev-ref HEAD)

echo "bench current state"

go test -run=NONE -bench=. ./casper/... > casper-b.bench
go test -run=NONE -bench=. ./common/... > common-b.bench
go test -run=NONE -bench=. ./link/...   > link-b.bench
go test -run=NONE -bench=. ./parser/... > parser-b.bench
go test -run=NONE -bench=. ./rules/...  > rules-b.bench

echo "stash and checkout master"
git stash
git checkout master

echo "bench master"
go test -run=NONE -bench=. ./casper/... > casper-a.bench
go test -run=NONE -bench=. ./common/... > common-a.bench
go test -run=NONE -bench=. ./link/...   > link-a.bench
go test -run=NONE -bench=. ./parser/... > parser-a.bench
go test -run=NONE -bench=. ./rules/...  > rules-a.bench

echo "go back to current state"
git checkout $currentBranch
git stash pop

benchcmp casper-a.bench casper-b.bench
benchcmp common-a.bench common-b.bench
benchcmp link-a.bench link-b.bench
benchcmp parser-a.bench parser-b.bench
benchcmp rules-a.bench rules-b.bench
