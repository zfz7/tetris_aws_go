[private]
alias bm := build-model
[private]
alias btsc := build-tsclient

build-model:
  @echo "Building Model..."
  @cd model && ./gradlew build

clean-model:
    @echo "Cleaning Model..."
    @cd model && rm -rf build

build-tsclient:
  just clean-tsclient
  @echo "Building TS Client..."
  @cd tsclient && cp -r ../model/build/smithyprojections/model/source/typescript-codegen/* .
  @cd tsclient && yarn install && yarn build

clean-tsclient:
    @echo "Cleaning TS Client..."
    bash -O extglob -c 'cd tsclient && rm -rf -- !("yarn.lock"|".gitignore")'
