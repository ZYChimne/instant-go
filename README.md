# instant-go

This is the Back-end Project of Instant, and you can visit the Front-end Project at [instant-vue](https://github.com/ZYChimne/instant-vue).

## Features

* High Performance and Scalable
* Access: RESTful
* Logical: grpc
* Storage: Redis, MySQL

## Project setup

```bash
go run cmd/main.go
```

### MySQL setup
```sql
CREATE DATABASE instant
    DEFAULT CHARACTER SET = 'utf8mb4';

CREATE TABLE instant.accounts (
    userid BIGINT UNSIGNED AUTO_INCREMENT UNIQUE PRIMARY KEY,
    mailbox VARCHAR(63) NOT NULL,
    phone VARCHAR(31) NOT NULL,
    pass_word VARCHAR(31) NOT NULL,
    username VARCHAR(31) NOT NULL,
    create_time DATETIME NOT NULL,
    update_time DATETIME NOT NULL,
    avatar TINYINT UNSIGNED NOT NULL,
    gender TINYINT UNSIGNED,
    country SMALLINT UNSIGNED,
    province SMALLINT UNSIGNED,
    city SMALLINT UNSIGNED,
    birthday DATETIME,
    school VARCHAR(63),
    company VARCHAR(63),
    my_mode TINYINT UNSIGNED,
    job VARCHAR(63),
    introduction VARCHAR(127),
    profile_image TINYINT UNSIGNED NOT NULL,
    tag VARCHAR(127),
    follow_ing TEXT,
    follow_ed TEXT,
    black_list TEXT,
    ins_id_list TEXT,
    like_id_list TEXT,
    comment_id_list TEXT
) DEFAULT CHARSET UTF8 COMMENT '';
```