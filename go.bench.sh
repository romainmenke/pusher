#!/bin/bash
set -e

echo "bench current state"

go test -run=NONE -bench=. ./common/... > common-b.bench
go test -run=NONE -bench=. ./casper/... > casper-b.bench
go test -run=NONE -bench=. ./link/...   > link-b.bench
go test -run=NONE -bench=. ./parser/... > parser-b.bench
go test -run=NONE -bench=. ./rules/...  > rules-b.bench

echo "stash and checkout master"
git stash
$currentBranch = git rev-parse --abbrev-ref HEAD
git checkout master

echo "bench master"
go test -run=NONE -bench=. ./common/... > common-a.bench
go test -run=NONE -bench=. ./casper/... > casper-a.bench
go test -run=NONE -bench=. ./link/...   > link-a.bench
go test -run=NONE -bench=. ./parser/... > parser-a.bench
go test -run=NONE -bench=. ./rules/...  > rules-a.bench

echo "go back to current state"
git checkout $currentBranch
git stash pop
