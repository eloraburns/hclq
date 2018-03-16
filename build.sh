go get -d -v && \
for GOOS in darwin linux; do
    for GOARCH in amd64; do
        export GOOS; export GOARCH; go build -v -o bin/hclq-$GOOS-$GOARCH
    done
done
