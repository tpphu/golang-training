## Command

```sh
echo $(pwd) | pbcopy
```

```sh
curl http://localhost:8080/register -d '{"code":"ABC","discount":0.05,"start":"2019-08-09T15:08:37.060Z","end":"2019-08-11T15:08:37.060Z"}' | jq .

curl http://localhost:8080/register -d '{"code":"ABC","discount":0.05,"start":"2019-08-10T15:08:37.060Z","end":"2019-08-12T15:08:37.060Z"}' | jq .
```