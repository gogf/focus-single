# 1. Focus Single

A `MVC` project using `GoFrame`.

# 2. Quick start

## 2.1 If you use SQLite

1. Download the source code
```
git clone https://github.com/gogf/focus-single.git
cd focus-single
```

2. Run Focus Single
```
cp manifest/config/config.example.yaml manifest/config/config.yaml
gf run main.go
```

3. Enjoy it
then open http://127.0.0.1:8199/ and enjoy it.

```
user:     goframe
password: 123456
```

## 2.2 If you use MySQL 

1. Download the source code
```
git clone https://github.com/gogf/focus-single.git
cd focus-single
```

2. Update config
copy manifest/config/config.yaml and edit database.default config
```
cp manifest/config/config.example.yaml manifest/config/config.yaml
```

3. Import Db
Import manifest/document/focus.sql to your Mysql

4. Run Focus Single And Enjoy
```
gf run main.go
```

then open http://127.0.0.1:8199/ and enjoy it.
```
user:     goframe
password: 123456
```

