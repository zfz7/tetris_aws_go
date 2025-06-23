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
  @echo "Building TS Client..."

clean-tsclient:
    @echo "Cleaning TS Client..."
    @cd tsclient && find . -mindepth 1 ! -name 'yarn.lock' -exec rm -rf {} + > /dev/null
