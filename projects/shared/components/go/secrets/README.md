# secrets
> get them from anywhere (aws for now) using go

## Example Output
This illustrates the operator-friendly responses for running the example code in
various contexts.

```
cd example && make

go build -o dist/main
APP_SECRET_VALUE_TEST=explicit ./dist/main # check when only value is supplied
2019/12/11 17:42:01 secret system initialized with support for protocol(s): aws-secretsmanager:, aws-ssm:
2019/12/11 17:42:01 using secret value explicitly provided in the environment at APP_SECRET_VALUE_TEST
APP_SECRET_PATH_TEST=aws-secretsmanager:test APP_SECRET_VALUE_TEST=override ./dist/main || true # check when both path and value are supplied
2019/12/11 17:42:01 secret system initialized with support for protocol(s): aws-secretsmanager:, aws-ssm:
2019/12/11 17:42:01 using secret value provided in the environment at APP_SECRET_VALUE_TEST, ignoring APP_SECRET_PATH_TEST
AWS_PROFILE=YOUR_PROFILE_HERE AWS_REGION=us-east-1 APP_SECRET_PATH_TEST=bad://test ./dist/main || true # when invalid protocol handler is used
2019/12/11 17:42:01 secret system initialized with support for protocol(s): aws-secretsmanager:, aws-ssm:
2019/12/11 17:42:01 unable to locate protocol handler for APP_SECRET_PATH_TEST=bad://test
APP_SECRET_PATH_TEST=aws-secretsmanager:test ./dist/main || true # check when path is supplied but region is missing
2019/12/11 17:42:01 secret system initialized with support for protocol(s): aws-secretsmanager:, aws-ssm:
2019/12/11 17:42:01 using path "test" at APP_SECRET_PATH_TEST in the environment to retrieve secret value using aws-secretsmanager: (*secrets.AWSSecretStore)
2019/12/11 17:42:01 secret retrieval error: MissingRegion: could not find region configuration
AWS_REGION=us-east-1 APP_SECRET_PATH_TEST=aws-secretsmanager:test ./dist/main || true # check when path is supplied but creds are missing
2019/12/11 17:42:01 secret system initialized with support for protocol(s): aws-secretsmanager:, aws-ssm:
2019/12/11 17:42:01 using path "test" at APP_SECRET_PATH_TEST in the environment to retrieve secret value using aws-secretsmanager: (*secrets.AWSSecretStore)
2019/12/11 17:42:22 secret retrieval error: NoCredentialProviders: no valid providers in chain. Deprecated.
        For verbose messaging see aws.Config.CredentialsChainVerboseErrors
AWS_PROFILE=YOUR_PROFILE_HERE AWS_REGION=us-east-1 APP_SECRET_PATH_TEST=aws-secretsmanager:test ./dist/main || true # good secret retrieval
2019/12/11 17:42:22 secret system initialized with support for protocol(s): aws-secretsmanager:, aws-ssm:
2019/12/11 17:42:22 using path "test" at APP_SECRET_PATH_TEST in the environment to retrieve secret value using aws-secretsmanager: (*secrets.AWSSecretStore)
2019/12/11 17:42:22 APP_SECRET_PATH_TEST secret at path "test" retrieved successfully using aws-secretsmanager: (*secrets.AWSSecretStore)
AWS_PROFILE=YOUR_PROFILE_HERE AWS_REGION=us-east-1 APP_SECRET_PATH_TEST=aws-secretsmanager:missing ./dist/main || true # missing secret retrieval
2019/12/11 17:42:22 secret system initialized with support for protocol(s): aws-secretsmanager:, aws-ssm:
2019/12/11 17:42:22 using path "missing" at APP_SECRET_PATH_TEST in the environment to retrieve secret value using aws-secretsmanager: (*secrets.AWSSecretStore)
2019/12/11 17:42:23 secret retrieval error: ResourceNotFoundException: Secrets Manager can't find the specified secret.
        status code: 400, request id: a04a5d3a-09f3-423b-adb1-3eca43731bb3
AWS_PROFILE=YOUR_PROFILE_HERE AWS_REGION=us-east-1 APP_SECRET_PATH_TEST=aws-ssm:test ./dist/main || true # good secret retrieval
2019/12/11 17:42:23 secret system initialized with support for protocol(s): aws-secretsmanager:, aws-ssm:
2019/12/11 17:42:23 using path "test" at APP_SECRET_PATH_TEST in the environment to retrieve secret value using aws-ssm: (*secrets.AWSParameterStore)
2019/12/11 17:42:24 APP_SECRET_PATH_TEST secret at path "test" retrieved successfully using aws-ssm: (*secrets.AWSParameterStore)
AWS_PROFILE=YOUR_PROFILE_HERE AWS_REGION=us-east-1 APP_SECRET_PATH_TEST=aws-ssm:missing ./dist/main || true # missing secret retrieval
2019/12/11 17:42:24 secret system initialized with support for protocol(s): aws-secretsmanager:, aws-ssm:
2019/12/11 17:42:24 using path "missing" at APP_SECRET_PATH_TEST in the environment to retrieve secret value using aws-ssm: (*secrets.AWSParameterStore)
2019/12/11 17:42:24 secret retrieval error: parameter not found
```
