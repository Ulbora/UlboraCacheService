Ulbora Cache Service
==============

An in-memory cache micro service based on freecache https://github.com/coocood/freecache.
Designed for use as a linked Docker container; therefore has no security added.


## Headers
- Content-Type: application/json (for POST and PUT)


## Set Value

```
POST:
URL: http://localhost:3010/rs/cache/set

Example Request
{
	"key": "testkey1",
	"value": "kens test value"
}
  
```

```
Example Response   

{
    "success": true
}

```



## Get Value

```
GET:
URL: http://localhost:3010/rs/cache/get/testkey1
  
```

```
Example Response   

{
    "success": true,
    "value": "kens test value"
}

```

## Set JSON Value as Base 64

```
POST:
URL: http://localhost:3010/rs/cache/set

Example Request
{
	"key": "testkey1",
	"value": "ew0KICAgInNvbWVBdHRyIjoic29tZVZhbCIsDQogICAiYW5vdGhlckF0dHIiOiJhbm90aGVyVmFsIg0KfQ=="
}
  
```

```
Example Response   

{
    "success": true
}

```



## Get JSON Value as Base 64

```
GET:
URL: http://localhost:3010/rs/cache/get/testkey1
  
```

```
Example Response   

{
    "success": true,
    "value": "ew0KICAgInNvbWVBdHRyIjoic29tZVZhbCIsDQogICAiYW5vdGhlckF0dHIiOiJhbm90aGVyVmFsIg0KfQ=="
}

```


## Get Value not found

```
GET:
URL: http://localhost:3010/rs/cache/get/testkey1
  
```

```
Example Response   

{
    "success": false,
    "value": ""
}

```


## Delete Value

```
DELETE:
URL: http://localhost:3010/rs/cache/delete/testkey1
  
```

```
Example Response   

{
    "success": true
}

```

# Docker use

```
docker run --network=ulbora_bridge --name cache --log-opt max-size=50m --restart=always -d ulboralabs/cache sh

```

