# instant-go

This is the Back-end Project of Instant, and you can visit the Front-end Project at [instant-vue](https://github.com/ZYChimne/instant-vue).

## Features

* High Performance, High Availability and Scalable
* Access: RESTful
* Storage: Redis, Mongodb (https://www.mongodb.com/developer/products/mongodb/mongodb-schema-design-best-practices/)
* Fan out on Write

## Project setup

```bash
sudo service redis-server start # wsl
gofmt ./ -w
golines ./ -w
```
```
$env:GOPROXY = "https://proxy.golang.com.cn,direct"
set http_proxy=socks5://127.0.0.1:7890
set https_proxy=%http_proxy%
```

## TODO
  
* Read Escape Analysis
* [Some High Availability Improvement](https://mp.weixin.qq.com/s/4T_S7nR8-HXJ59IbK4FBWQ)
* AVIF Support
* [Logging System](https://mp.weixin.qq.com/s/dUs7WUKUDOf9lLG6tzdk0g)
* Hot spot detection and local cache

## MongoDB 

### Users

#### Validator

```
{
  $jsonSchema: {
    bsonType: 'object',
    required: [
      'mailbox',
      'phone',
      'username',
      'password',
      'created',
      'lastModified',
      'avatar',
      'gender',
      'country',
      'province',
      'city',
      'birthday',
      'school',
      'company',
      'job',
      'myMode',
      'introduction',
      'coverPhoto',
      'tags',
      'followings',
      'followers'
    ],
    properties: {
      mailbox: {
        bsonType: 'string',
        pattern: '[a-zA-Z]+@[a-zA-Z]+.[a-zA-Z]+',
        maxLength: 64,
        description: 'must be a valid email address'
      },
      phone: {
        bsonType: 'string',
        maxLength: 11,
        minLength: 10,
        pattern: '^[0-9]+$',
        description: 'must be a valid phone number'
      },
      username: {
        bsonType: 'string',
        maxLength: 16,
        description: 'must be in length [0, 16]'
      },
      password: {
        bsonType: 'string',
        description: 'must be a valid password'
      },
      created: {
        bsonType: 'date',
        description: 'must be a valid date'
      },
      lastModified: {
        bsonType: 'date',
        description: 'must be a valid date'
      },
      avatar: {
        bsonType: 'int',
        maximum: 10,
        minimum: 0,
        description: 'must be in range [0, 10]'
      },
      gender: {
        bsonType: 'int',
        maximum: 2,
        minimum: 0,
        description: 'must be in range [0, 2]'
      },
      country: {
        bsonType: 'int',
        maximum: 64,
        minimum: 0,
        description: 'must be a valid country'
      },
      province: {
        bsonType: 'int',
        maximum: 64,
        minimum: 0,
        description: 'must be a valid province'
      },
      city: {
        bsonType: 'int',
        maximum: 64,
        minimum: 0,
        description: 'must be a valid city'
      },
      birthday: {
        bsonType: 'date',
        description: 'must be a valid date'
      },
      school: {
        bsonType: 'string',
        maxLength: 16,
        description: 'must be in length [0, 16]'
      },
      company: {
        bsonType: 'string',
        maxLength: 16,
        description: 'must be in length [0, 16]'
      },
      myMode: {
        bsonType: 'int',
        maximum: 10,
        minimum: 0,
        description: 'must be in range [0, 10]'
      },
      introduction: {
        bsonType: 'string',
        maxLength: 256,
        description: 'must be in length [0, 256]'
      },
      coverPhoto: {
        bsonType: 'int',
        maximum: 10,
        minimum: 0,
        description: 'must be in range [0, 10]'
      },
      tags: {
        bsonType: 'array',
        minItems: 0,
        maxItems: 10,
        description: 'must be in length [0, 256]'
      },
      followings: {
        bsonType: 'int',
        minimum: 0,
        description: 'must be greater than 0'
      },
      followers: {
        bsonType: 'int',
        minimum: 0,
        description: 'must be greater than 0'
      }
    }
  }
}
```

### Instants

#### Validator

```
{
  $jsonSchema: {
    bsonType: 'object',
    required: [
      'userID',
      'created',
      'lastModified',
      'content',
      'likes',
      'shares'
    ],
    properties: {
      userID: {
        bsonType: 'objectId',
        description: 'must be a valid userID'
      },
      created: {
        bsonType: 'date',
        description: 'must be a valid date'
      },
      lastModified: {
        bsonType: 'date',
        description: 'must be a valid date'
      },
      content: {
        bsonType: 'string',
        maxLength: 256,
        description: 'must be in length [0, 256]'
      },
      refOriginID: {
        bsonType: 'objectId',
        description: 'must be a valid refOriginID'
      },
      likes: {
        bsonType: 'int',
        minimum: 0,
        description: 'must be greater than 0'
      },
      shares: {
        bsonType: 'int',
        minimum: 0,
        description: 'must be greater than 0'
      }
    }
  }
}
```

### followings

#### Validator

```
{
  $jsonSchema: {
    bsonType: 'object',
    required: [
      'userID',
      'followingID',
      'lastModified'
    ],
    properties: {
      userID: {
        bsonType: 'objectId',
        description: 'must be a valid userID'
      },
      followingID: {
        bsonType: 'objectId',
        description: 'must be a valid followingID'
      },
      lastModified: {
        bsonType: 'date',
        description: 'must be a valid date'
      }
    }
  }
}
```