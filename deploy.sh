make

aws cloudformation package \
  --template-file demo.sam.yaml \
  --output-template-file demo.out.yaml \
  --s3-bucket <bucket name>

aws cloudformation deploy \
  --template-file demo.out.yaml \
  --capabilities CAPABILITY_IAM \
  --stack-name <stack name>
