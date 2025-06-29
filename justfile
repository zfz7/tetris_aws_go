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
    @cd frontend && yarn install && yarn build

# Cleans frontend, removes build folder
clean-frontend:
    @echo "Cleaning Frontend..."
    @cd frontend && rm -rf build

# Builds backend
build-backend:
    @echo "Building Backend..."
    @cd backend/api && go generate
    @cd backend && go fmt ./...
    @cd backend && go test ./...
    @cd backend && GOOS=linux GOARCH=arm64 go build -o bootstrap ./lambda/main.go && zip lambdaFunction.zip bootstrap

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
