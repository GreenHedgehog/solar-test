# API Documentation

## **Open Endpoints**

Endpoints that do not require authentication

### **Login** 

Getting a token for a user with the right to edit.

**URL** : `/login`

**METHOD** : `POST`

**Example Request** :
```json
{
  "login": "login",
  "hash": "hash"
}
```

**Example Responce**:
##### **code** : `200 ok`

##### **Body example**

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MzkxOTQ0MDZ9.SStozlh8nuZQLa8ek5f4wpnO4s7snBp1jI8g219fVZM"
}
```

### **All vacancy** 

Getting info for all vacancy

**URL** : `/vacancy`

**METHOD** : `GET`

**Example Responce**:
##### **code** : `200 ok`

##### **Body example**

```json
[
	{
		"id": 1,
		"name": "best vacancy",
		"salary": 120000,
		"experience": "3-5",
		"place": "msc"
	}
]
```


### **Get vacancy** 

Get vacancy info by id

**URL** : `/vacancy/:id`

**METHOD** : `GET`

**Example Request** : `/vacancy/1`

**Example Responce**:
##### **code** : `200 ok`

##### **Body example**

```json
{
	"id": 1,
	"name": "best vacancy",
	"salary": 120000,
	"experience": "3-5",
	"place": "msc"
}
```

## **Secured Endpoints**

Endpoints that require authentication.

**Required Header**:

`Authorization: Bearer <token from login>`

### **Delete vacancy**

Deletes vacancy with id

**URL** : `/vacancy/:id`

**METHOD** : `DELETE`

**Example Request** : `/vacancy/1`

**Example Responce**:
##### **code** : `204 No Content`

### **Add vacancy**

Creates new vacancy with given info

**URL** : `/vacancy`

**METHOD** : `PUT`

**Example Request** : 
```json
{
	"name": "best vacancy",
	"salary": 120000,
	"experience": "3-5",
	"place": "msc"
}
```

**Example Responce**:
##### **code** : `201 Created`

##### **Body example**

```json
{
	"id": 1,
	"name": "best vacancy",
	"salary": 120000,
	"experience": "3-5",
	"place": "msc"
}
```
