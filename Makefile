build_oauth2_introspection:
	cd ./oauth2-introspection && go build

package_oauth2_introspection:
	cd ./oauth2-introspection && GOOS=linux GOARCH=amd64 go build -o oauth2-introspection.linux-amd64
	cd ./oauth2-introspection && docker build -t secure-api-gateway:oauth2-introspection .
