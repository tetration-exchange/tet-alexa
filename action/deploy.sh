#!/bin/sh

rm -rf deploy
mkdir deploy
GOOS=linux go build -o ./deploy/tet-alexa-action
pushd deploy
zip tet-alexa-action.zip tet-alexa-action
popd
aws cloudformation package \
   --template-file action.yaml \
   --output-template-file deploy.yaml \
   --s3-bucket tet-alexa
aws cloudformation deploy\
  --template-file deploy.yaml\
  --stack-name tet-alexa-action-lambda\
  --capabilities CAPABILITY_IAM