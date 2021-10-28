# geneve
to compile the plugin:
1. delete all the required packages from the mod file
2. go get github.com/openziti/ziti@v0.22.7
3. go mod tidy -go=1.16 && go mod tidy -go=1.17
4. go build -buildmode=plugin