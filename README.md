## GO NOTES

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
### Go Lang Official Package Community

[Package Go Dev](https://pkg.go.dev/)

### Resources

[Programiz Golang](https://www.programiz.com/golang/getting-started)

[W3school Golang](https://www.w3schools.com/go/go_formatting_verbs.php)
