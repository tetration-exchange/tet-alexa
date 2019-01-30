#!/bin/sh

rm -rf deploy
mkdir deploy
GOOS=linux go build -o ./deploy/tet-alexa
pushd deploy
zip tet-alexa.zip tet-alexa
popd
aws cloudformation package \
   --template-file tet-alexa.yaml \
   --output-template-file serverless-deploy_tet_alexa.yaml \
   --s3-bucket tet-alexa
aws cloudformation deploy\
  --template-file serverless-deploy_tet_alexa.yaml\
  --stack-name tet-alexa-lambda\
  --capabilities CAPABILITY_IAM