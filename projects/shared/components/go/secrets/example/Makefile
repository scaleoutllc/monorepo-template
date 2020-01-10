main:
	go build -o dist/main
	APP_SECRET_VALUE_TEST=explicit ./dist/main # check when only value is supplied
	APP_SECRET_PATH_TEST=aws-secretsmanager:test APP_SECRET_VALUE_TEST=override ./dist/main || true # check when both path and value are supplied
	AWS_PROFILE=YOUR_PROFILE_HERE AWS_REGION=us-east-1 APP_SECRET_PATH_TEST=bad://test ./dist/main || true # when invalid protocol handler is used
	APP_SECRET_PATH_TEST=aws-secretsmanager:test ./dist/main || true # check when path is supplied but region is missing
	AWS_REGION=us-east-1 APP_SECRET_PATH_TEST=aws-secretsmanager:test ./dist/main || true # check when path is supplied but creds are missing
	AWS_PROFILE=YOUR_PROFILE_HERE AWS_REGION=us-east-1 APP_SECRET_PATH_TEST=aws-secretsmanager:test ./dist/main || true # good secret retrieval
	AWS_PROFILE=YOUR_PROFILE_HERE AWS_REGION=us-east-1 APP_SECRET_PATH_TEST=aws-secretsmanager:missing ./dist/main || true # missing secret retrieval
	AWS_PROFILE=YOUR_PROFILE_HERE AWS_REGION=us-east-1 APP_SECRET_PATH_TEST=aws-ssm:test ./dist/main || true # good secret retrieval
	AWS_PROFILE=YOUR_PROFILE_HERE AWS_REGION=us-east-1 APP_SECRET_PATH_TEST=aws-ssm:missing ./dist/main || true # missing secret retrieval
