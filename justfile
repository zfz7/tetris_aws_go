[private]
alias bm := build-model
[private]
alias btsc := build-tsclient
[private]
alias bf := build-frontend
[private]
alias bb := build-backend
[private]
alias bc := build-cdk
[private]
alias b := build
[private]
alias d := deploy
[private]
alias dq := deploy-quick
[private]
alias s := serve
[private]
alias l := lint
[private]
alias ut := unit-test


_default:
  just --list

#Builds all packages
build:
    @echo "Building..."
    @just build-model
    @just build-backend
    @just build-tsclient
    @just build-frontend
    @just build-cdk

# Cleans all packages
clean:
    @echo "Building..."
    @just clean-model
    @just clean-backend
    @just clean-tsclient
    @just clean-frontend
    @just clean-cdk

# Deploys all packages (assumes packages are built)
deploy:
    @echo "Deploying..."
    @cd cdk && yarn deploy

# Builds the backend go service binary and deploys it without using CDK
[confirm("Are you sure no changes have been made to the CDK / Smithy / APIGateway?")]
deploy-quick:
    @cd backend && \
            GOOS=linux \
            GOARCH=arm64 \
            CGO_ENABLED=0 \
            go build -ldflags="-s -w" -o bootstrap ./lambda/main.go
    @cd backend && zip lambdaFunction.zip bootstrap
    @bash -c ' \
    FUNCTION_NAME=$(aws lambda list-functions \
    --region "us-west-2" \
    --query "Functions[?starts_with(FunctionName, \`Tetris-Go-Service\`)].FunctionName" \
    --output text); \
    echo "Lambda function: $FUNCTION_NAME"; \
    if aws lambda update-function-code --region "us-west-2" --function-name "$FUNCTION_NAME" --zip-file fileb://./backend/lambdaFunction.zip > /dev/null; then \
    echo "✅ Lambda function code updated successfully"; \
    else \
    echo "❌ Failed to update Lambda function"; \
    fi \
    '
# Build Smithy Models
build-model:
  @echo "Building Model..."
  @cd model && ./gradlew build

# Cleans model packages (removes ./build folder)
clean-model:
    @echo "Cleaning Model..."
    @cd model && rm -rf build

# Build Ts client from model
build-tsclient:
  just clean-tsclient
  @echo "Building TS Client..."
  @cd tsclient && cp -r ../model/build/smithyprojections/model/source/typescript-codegen/* .
  @cd tsclient && yarn install && yarn build

# Cleans TS Client, removes all files expect yarn.lock and .gitignore
clean-tsclient:
    @echo "Cleaning TS Client..."
    bash -O extglob -c 'cd tsclient && rm -rf -- !("yarn.lock"|".gitignore"|"node_modules")'

# Builds frontend, requires TS Client to be built
build-frontend:
    @echo "Building Frontend..."
    @cd frontend && yarn build

# Runs frontend locally
serve:
 @cd frontend && yarn vite serve

# Cleans frontend, removes build folder
clean-frontend:
    @echo "Cleaning Frontend..."
    @cd frontend && rm -rf build

build-backend:
    @cd backend/api && go generate ./...
    @cd backend && go fmt ./...
    @cd backend && goimports -w .
    @cd backend && golangci-lint run ./...
    @cd backend && go test -v -race -cover ./...
    @cd backend && gosec ./...
    @cd backend && govulncheck ./...
    @cd backend && \
        GOOS=linux \
        GOARCH=arm64 \
        CGO_ENABLED=0 \
        go build -ldflags="-s -w" -o bootstrap ./lambda/main.go
    @cd backend && zip lambdaFunction.zip bootstrap

# Cleans backend, removes bootstrap and lambdaFunction.zip
clean-backend:
    @echo "Cleaning Backend..."
    @cd backend && rm -rf bootstrap && rm -rf lambdaFunction.zip

# Builds cdk
build-cdk:
    @echo "Building CDK..."
    @cd cdk && yarn install && yarn build

# Cleans cdk, removes build folder
clean-cdk:
    @echo "Cleaning CDK..."
    @cd cdk && yarn clean && rm -rf dist

lint:
    @cd backend && golangci-lint run ./...
    @cd cdk && yarn format
    @cd frontend && yarn lint --fix

unit-test:
    @cd backend && go test ./...
    @cd frontend && yarn vitest --run