test:
	sam local invoke -e ./sample.json

release:
	bundle install
	bundle install --deployment
	sam package --profile opszero --template-file template.yaml --output-template-file packaged-template.yaml --s3-bucket landingcrew-assets
	sam deploy  --profile opszero --region us-west-2 --template-file packaged-template.yaml --stack-name $(shell basename $(shell pwd)) --capabilities CAPABILITY_IAM
