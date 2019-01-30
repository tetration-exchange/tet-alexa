#!/bin/sh

rm -rf deploy
mkdir deploy
GOOS=linux go build -o ./deploy/tet-alexa-notification
pushd deploy
zip tet-alexa-notification.zip tet-alexa-notification
popd
aws cloudformation package \
   --template-file notification.yaml \
   --output-template-file deploy.yaml \
   --s3-bucket tet-alexa
aws cloudformation deploy\
  --template-file deploy.yaml\
  --stack-name tet-alexa-notification-lambda\
  --capabilities CAPABILITY_IAM