## GO NOTES

### Run Go File
Run Go File
```
$go run hello.go
```
Compile Go File
```
$go build hello.go
$./hello
```
Specify Different Program Name
```
$go build -o myprogram hello.go
$./myprogram
```
### Creating Go Module
```
$mkdir test_module
$go mod init sample_module
$go build .
$go run .
$./sample_module
```
### Creating Go Workspace
```
$mkdir go_workspace
$cd go_workspace
$mkdir module_one && mkdir module_two
$go work init ./module_one
$go work use ./module_two
$go run ./module_one/module1
$go build ./module_one/module1
$./module1
```
### Framework
```
Gin (Web Framework)
GORM (ORM)
Gorilla Mux (Http Router)
```

### Run Go Executable using PM2
```
$pm2 start go_simple_api --interpreter none --watch
$pm2 list

┌────┬──────────────────┬─────────────┬─────────┬─────────┬──────────┬────────┬──────┬───────────┬──────────┬──────────┬──────────┐
│ id │ name             │ namespace   │ version │ mode    │ pid      │ uptime │ ↺    │ status    │ cpu      │ mem      │ watching │
├────┼──────────────────┼─────────────┼─────────┼─────────┼──────────┼────────┼──────┼───────────┼──────────┼──────────┼──────────┼
│ 1  │ go_simple_api    │ default     │ N/A     │ fork    │ 13032    │ 20m    │ 0    │ online    │ 0%       │ 1.6mb    | enabled  │
└────┴──────────────────┴─────────────┴─────────┴─────────┴──────────┴────────┴──────┴───────────┴──────────┴──────────┴──────────┴

$pm2 delete 1
```
```
Browser access:
http://localhost:8081/articles
```
### Go Lang Official Package Community

[Package Go Dev](https://pkg.go.dev/)

### Resources
[Go Dev](https://go.dev/dl/)

[Programiz Golang](https://www.programiz.com/golang/getting-started)

[W3school Golang](https://www.w3schools.com/go/go_formatting_verbs.php)

