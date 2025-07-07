# tetris_aws_go

## Goals

The goal of this project was to create a full stack AWS native project template. The secondary goal is to keep the DTO
types in sync between the frontend and backend using code generation. Here is the achieved stack:

* Backend:
    * [Go](https://go.dev/)
    * [Lambda](https://aws.amazon.com/lambda/)
    * [ApiGateway](https://aws.amazon.com/api-gateway/)
* Frontend:
    * [Typescript](https://www.typescriptlang.org/)
    * [React](https://react.dev/)
* Infrastructure:
    * Build Tool: [Gradle](https://gradle.org/) and [Yarn](https://classic.yarnpkg.com/lang/en/)
    * Deployments: [CDK](https://aws.amazon.com/cdk/)
    * Authentication: [Cognito](https://aws.amazon.com/cognito/)
    * Api Definition: [Smithy](https://smithy.io/2.0/index.html)
    * Go Type Generation: [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen)
    * Typescript Type Generation: [smithy-typescript](https://github.com/awslabs/smithy-typescript)

## Warnings

[smithy-go](https://github.com/awslabs/smithy-go) Does not yet work. Track this [issue](https://github.com/aws/smithy-go/issues/458) to see updated status

The go and typescript type generation rely on [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen)
and [smithy-typescript](https://github.com/awslabs/smithy-typescript), both products are not yet GA. As such they still
have some quirks to work through. Overall it seems that the typescript code generation tool is a bit more stable.

If you stumbled across this guide another resource that wasn't GA but looks promising
is: [aws-prototyping-sdk](https://github.com/aws/aws-prototyping-sdk), specifically
the [type-safe-api](https://github.com/aws/aws-prototyping-sdk/blob/mainline/packages/type-safe-api/README.md).

^last updated 11/05/2023

## Getting started

### Prerequisites

This project assumes the following are installed on your system:

* [Gradle](https://gradle.org/)
* [Yarn](https://classic.yarnpkg.com/lang/en/)
    * concurrently `yarn global add concurrently`
* [AWS CLI](https://aws.amazon.com/cli/)
* [AWS CDK](https://docs.aws.amazon.com/cdk/v2/guide/cli.html)
* An AWS will SSO setup:
    * https://docs.aws.amazon.com/cdk/v2/guide/getting_started.html#getting_started_auth
    * https://docs.aws.amazon.com/sdkref/latest/guide/access-sso.html

### Setup
1. Clone the repo: `git clone git@github.com:zfz7/tetris_aws.git`
2. Export the following variables:
```
export AWS_PROFILE=AdministratorAccess-123456789012
export AWS_ACCOUNT=123456789012
export ROOT_HOSTED_ZONE_ID_GO=ABCDEFGHIJKLIMOP
export ROOT_HOSTED_ZONE_NAME_GO=example.com
```
3. Build the project (from root): `.just build`
4. Login to AWS: `aws sso login`
5. Deploy the project (from root): `just deploy`

### Other Commands

```
just --list
Available recipes:
    build          # Builds all packages
    build-backend  # Builds backend
    build-cdk      # Builds cdk
    build-frontend # Builds frontend, requires TS Client to be built
    build-model    # Build Smithy Models
    build-tsclient # Build Ts client from model
    clean          # Cleans all packages
    clean-backend  # Cleans backend, removes bootstrap and lambdaFunction.zip
    clean-cdk      # Cleans cdk, removes build folder
    clean-frontend # Cleans frontend, removes build folder
    clean-model    # Cleans model packages (removes ./build folder)
    clean-tsclient # Cleans TS Client, removes all files expect yarn.lock and .gitignore
    deploy         # Deploys all packages (assumes packages are built)

###Testing the endpoint with Cognito
export C_TOKEN="$(aws cognito-idp initiate-auth --region us-west-2 --auth-flow USER_PASSWORD_AUTH --client-id <YOUR_CLIENT_ID> --auth-parameters USERNAME=<USERNAME>,PASSWORD=<PASSWORD> | jq -r .AuthenticationResult.IdToken)"'
curl -H "Authorization: Bearer $C_TOKEN" https://api.daniel-eichman.com/hello\?name\=hi
```
