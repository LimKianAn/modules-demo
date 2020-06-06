# Build
```
go build .
# Run the genereated binary
# Open localhost:3000 with a browser
```

# Go modules Commands
```
go build .
go build -mod=readonly .
go build -mod=vendor .
go get rsc.io/quote
go get rsc.io/quote/v2
go get -u github.com/gorilla/mux
go get github.com/gorilla/mux@latest
go get github.com/gorilla/mux@master
go get "github.com/gorilla/mux@<v1.7.0" # Closest match wins, not the latest!
go list
go list all
go list -m all
go list -m -versions github.com/gorilla/mux # versions plural
go mod download github.com/gorilla/context@master
go mod edit -module github.com/LimKianAn/modules-examples
go mod edit -go 1.12
go mod edit -require github.com/gorilla/context@v1.0.0
go mod edit -droprequire github.com/gorilla/context
go mod edit -exclude github.com/gorilla/mux@v1.7.0
go mod edit -dropexclude github.com/gorilla/mux@v1.7.0
go mod edit -replace rsc.io/quote@v1.5.2=../quote
go mod edit -dropreplace rsc.io/quote@v1.5.2
go mod edit -print
go mod edit -json
go mod graph
go mod help edit
go mod init github.com/LimKianAn/modules-demo
go mod tidy # "go build ." doesn't run this command automatically.
go mod verify # "go build ." still works even if the verification fails.
go mod why -m github.com/gorilla/context
```