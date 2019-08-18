## Todo
- [ ] Make a version to demo how data can duplicate data.
- [ ] Write k6 script to see

## Command

```sh
echo $(pwd) | pbcopy
```

```sh
curl http://localhost:8080/register -d '{"code":"ABC","discount":0.05,"start":"2019-08-09T15:08:37.060Z","end":"2019-08-11T15:08:37.060Z"}' | jq .

curl http://localhost:8080/register -d '{"code":"ABC","discount":0.05,"start":"2019-08-10T15:08:37.060Z","end":"2019-08-12T15:08:37.060Z"}' | jq .
```

```sql
select count(*)
from voucher;

select v1.id, v1.code, v1.`start`, v1.`end`, v2.id, v2.`start`, v2.`end`
from voucher as v1
join voucher as v2
on v1.code = v2.code
where v2.id > v1.id
and  v2.`start`<= v1.`end` AND v2.`end` >= v1.`start` 
LIMIT 10;
```