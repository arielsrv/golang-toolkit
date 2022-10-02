module github.com/arielsrv/golang-toolkit/examples/restclient

go 1.19

require (
	github.com/arielsrv/golang-toolkit/common v0.0.6
	github.com/arielsrv/golang-toolkit/restclient v0.2.1
	github.com/go-http-utils/headers v0.0.0-20181008091004-fed159eddc2a
	github.com/ldez/mimetype v0.1.0
	github.com/stretchr/testify v1.8.0
)

replace github.com/arielsrv/golang-toolkit/common => ../../common

replace github.com/arielsrv/golang-toolkit/restclient => ../../restclient

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.4.0 // indirect
	github.com/tjarratt/babble v0.0.0-20210505082055-cbca2a4833c1 // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
