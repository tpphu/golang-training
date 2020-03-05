
## MySQL

```
// Should show error
mysql -uroot -proot -h127.0.0.1 -p3306
// Should correctly
mysql -uroot -proot -h127.0.0.1 -P3306
create database crawler;
GRANT ALL PRIVILEGES ON crawler.* TO 'default'@'%' WITH GRANT OPTION;
```