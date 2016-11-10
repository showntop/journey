# api specification
==========

apps share and perfect

##create user

```
method: post
url:    /users
request body:
    {
        mobile: "",
        password: ""
    }
response body:
    {
        state_code:
        message:
        data:{
            id: 
            username:
            mobile:
        }
    }
```

## user login

```
method: post
url:    /sessions
request body:
    {
        field_value: username/mobile/email,
        password: ""
    }
response body:
    {
        state_code:
        message:
        data:{
            id: 
            username:
            mobile:
        }
    }
```

## get category list

```
method: get
url:    /categories
response body:
    {
        state_code:
        message:
        data:{
            {
              "id": 16110000,
              "name": ""
            },
        }
    }
```

## get app list

```
method: get
url:    /apps?category_id={:id}&page_no={:no}&page_num={:num}
response body:
    {
        state_code:
        message:
        data:{
            {
              "id": 16110000,
              "category_id": 16110000,
              "name": "",
              "version": "1.09",
              "description": "",
              "size": 16714301,
              "logo_url": "",
              "download_url": "",
              "Assets": null,
              "tags": [
                "t1",
                "t2",
              ]
            },
        }
    }
```

## get app detail info

```
method: get
url:    /apps/{:app_id}
response body:
    {
        state_code:
        message:
        data:{
            {
              "id": 16110000,
              "category_id": 16110000,
              "name": "",
              "version": "1.09",
              "description": "",
              "size": 16714301,
              "logo_url": "",
              "download_url": "",
              "Assets": null,
              "tags": [
                "t1",
                "t2",
              ]
            },
        }
    }
```

