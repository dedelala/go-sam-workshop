#!/bin/bash

die() { echo "oh noes: $*" >&2; exit 1; }
usage() { echo "usage: $0 <bucket> <stack>" >&2; exit 2; }

bucket=$1
[[ -n $bucket ]] || usage
stack=$2
[[ -n $stack ]] || usage

for x in go zip sam aws; do
    hash "$x" &>/dev/null || die "missing $x"
done

GO111MODULE=on GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o wc || die "build"
zip wc.zip wc || die "zip"

sam package --template-file template.yml --s3-bucket "$bucket" \
  --output-template-file sam.yml || die "package"

sam deploy --template-file sam.yml --stack-name "$stack" \
  --capabilities CAPABILITY_IAM || die "deploy"

echo "(ﾉ◕ヮ◕)ﾉ*:・ﾟ✧ $stack deployed!"
aws cloudformation describe-stacks --stack-name "$stack" \
  --query "Stacks[].Outputs[].OutputValue" --output text
