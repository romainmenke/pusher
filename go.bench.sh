#!/bin/bash
set -e

go test -run=NONE -bench=. ./common/... > common-b.bench
go test -run=NONE -bench=. ./casper/... > casper-b.bench
go test -run=NONE -bench=. ./link/...   > link-b.bench
go test -run=NONE -bench=. ./parser/... > parser-b.bench
go test -run=NONE -bench=. ./rules/...  > rules-b.bench

git stash

$currentBranch = git rev-parse --abbrev-ref HEAD

git checkout master

go test -run=NONE -bench=. ./common/... > common-a.bench
go test -run=NONE -bench=. ./casper/... > casper-a.bench
go test -run=NONE -bench=. ./link/...   > link-a.bench
go test -run=NONE -bench=. ./parser/... > parser-a.bench
go test -run=NONE -bench=. ./rules/...  > rules-a.bench

git checkout $currentBranch
git stash pop
