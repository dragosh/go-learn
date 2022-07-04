start/dev:
	CompileDaemon -directory="." -include="*.js" -include="*.css" -include="*.tmpl" -build="go build -o ./target/go-learn cmd/go-learn/main.go" -command="./target/go-learn" -graceful-kill=true -log-prefix=false



start/web:
	@go run cmd/bundle.go
	@go run cmd/web.go
