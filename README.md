# DBO Backend test API documentation

* **End Point:**
/dbo/login

* **Description:**
This API is used to login using basic auth. 

* **Method:**
POST

* **Success Response:**
{
    "code": 200,
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTg1MDc3MzksImlzcyI6IkRCTyBURVNUIiwiVXNlcm5hbWUiOiJjb2JhIDEyMyJ9.LlKMYhY1yfRpJtLEra2SzH5zZkGRRWEGAEAKTwsVKUA",
    "token_type": "Bearer",
    "expires_in": 3600
}

* **Error Response:**
  * **code:** 400 <br />
  * **Content:** {
        "code": 400,
        "message": "Invalid username or password"
    }

  * **code:** 500 <br />
  * **Content:** {
        "code": 500,
        "message": "Internal Server Error"
    }

------------------------------------------------------------------------------------------------------------------------------------------------

* **End Point:**
/dbo/customers/create-customer

* **Description:**
This API is used for creating user/similiar to register user

* **Method:**
POST

* **URL Query Params:**
none

* **URL Params:**
none

* **JSON Body:**
{
    "username" : "coba123",
    "password" : "Jakarta1!",
    "firstname": "coba",
    "lastname" : "123",
    "email" : "coba9@mail.com",
    "address": "Jl Margonde"
}

* **Success Response:**
{
    "code": 201,
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTg1MDYzNjksImlzcyI6IkRCTyBURVNUIiwiVXNlcm5hbWUiOiJjb2JhMTIzIn0.YfO_5RftReRC0vYTwDo0-ewHIkoMzIF3gWNAPo_Lqh8",
    "token_type": "Bearer",
    "expires_in": 3600
}

* **Error Response:**
  * **code:** 400 <br />
  * **Content:** {
        "code": 400,
        "message": "your email is already exist"
    }

  * **code:** 400 <br />
  * **Content:** {
        "code": 400,
        "message": "your username is already exist"
    }  

  * **code:** 500 <br />
  * **Content:** {
        "code": 500,
        "message": "Failed to create account. Please call customer service"
    }

------------------------------------------------------------------------------------------------------------------------------------------------

* **End Point:**
/dbo/customers/get-all-customer

* **Description:**
This API is used to get all customer data. It's searchable too using email key word and using pagination too

* **Method:**
GET

* **URL Query Params:**
email=...
page=...

* **URL Params:**
none

* **JSON Body:**
none

* **Success Response:**
[
    {
        "ID": 1,
        "Username": "coba 123",
        "FirstName": "coba",
        "LastName": "123",
        "Email": "coba123@mail.com",
        "Address": "Jl Margonde"
    },
    {
        "ID": 2,
        "Username": "coba 123",
        "FirstName": "coba",
        "LastName": "123",
        "Email": "coba1223@mail.com",
        "Address": "Jl Margonde"
    },
    {
        "ID": 4,
        "Username": "coba123",
        "FirstName": "coba",
        "LastName": "123",
        "Email": "coba2@mail.com",
        "Address": "Jl Margonde"
    }
]

* **Error Response:**
  * **code:** 500 <br />
  * **Content:** {
        "code": 500,
        "message": "internal server error"
    }
------------------------------------------------------------------------------------------------------------------------------------------------

* **End Point:**
/dbo/customers/get-customer-by-id

* **Description:**
This API is used to get detail customer data by customer ID

* **Method:**
GET

* **URL Query Params:**
custId=...

* **URL Params:**
none

* **JSON Body:**
none

* **Success Response:**
{
    "ID": 1,
    "Username": "coba 123",
    "FirstName": "coba",
    "LastName": "123",
    "Email": "coba123@mail.com",
    "Address": "Jl Margonde"
}

* **Error Response:**
  * **code:** 400 <br />
  * **Content:** {
        "code": 400,
        "message": "missing customer id"
    }

  * **code:** 500 <br />
  * **Content:** {
        "code": 500,
        "message": "internal server error"
    }

------------------------------------------------------------------------------------------------------------------------------------------------

* **End Point:**
/dbo/customers/update-customer

* **Description:**
This API is used to update user data. You need to pass customer id query param to specify which user you want to update

* **Method:**
PUT

* **URL Query Params:**
custId=...

* **URL Params:**
none

* **JSON Body:**
You can add whatever payload as long as it's property registered in customer struct. The json form of the customer struct is: <br/>
{
    "username" : "coba123",
    "password" : "Jakarta1!",
    "firstname": "coba",
    "lastname" : "123",
    "email" : "coba9@mail.com",
    "address": "Jl Margonde"
}

* **Success Response:**
{
    "code": 200,
    "message": "Success update data"
}

* **Error Response:**
  * **code:** 400 <br />
  * **Content:** {
        "code": 400,
        "message": "missing customer id"
    }

  * **code:** 500 <br />
  * **Content:** {
        "code": 500,
        "message": "Internal Server Error"
    }

------------------------------------------------------------------------------------------------------------------------------------------------

* **End Point:**
/dbo/customers/delete-customer

* **Description:**
This API is used to delete user data. You need to pass customer id query param to specify which user you want to delete

* **Method:**
DELETE

* **URL Query Params:**
custId=...

* **URL Params:**
none

* **JSON Body:**
none

* **Success Response:**
{
    "code": 200,
    "message": "User deleted"
}

* **Error Response:**
  * **code:** 400 <br />
  * **Content:** {
        "code": 400,
        "message": "missing customer id"
    }

  * **code:** 500 <br />
  * **Content:** {
        "code": 500,
        "message": "Internal Server Error"
    }

------------------------------------------------------------------------------------------------------------------------------------------------

* **End Point:**
/dbo/orders/create-order

* **Description:**
This API is used to create the order

* **Method:**
POST

* **URL Query Params:**
none

* **URL Params:**
none

* **JSON Body:**
{
    "name" : "Sirloin",
    "quantity" : 1
}

* **Success Response:**
{
    "CreatedAt": "2024-06-16T12:14:46.5979086+07:00",
    "UpdatedAt": "2024-06-16T12:14:46.5979086+07:00",
    "DeletedAt": null,
    "ID": 12,
    "Name": "Sirloin",
    "Quantity": 1,
    "CustomerID": 7,
    "Customer": {
        "CreatedAt": "2024-06-16T07:56:41.241+07:00",
        "UpdatedAt": "2024-06-16T07:56:41.241+07:00",
        "DeletedAt": null,
        "ID": 7,
        "Username": "coba6",
        "Password": "$2a$10$wyPdbGPq9Fh/9xPmAEmoeOJ9uNt2Q/6AZZMzFYrmAiL6.vwtIm3q6",
        "FirstName": "coba",
        "LastName": "123",
        "Email": "coba8@mail.com",
        "Address": "Jl Margonde",
        "Order": null
    }
}

* **Error Response:**
  * **code:** 500 <br />
  * **Content:** {
        "code": 500,
        "message": "Failed to create order. Please call customer service"
    }

------------------------------------------------------------------------------------------------------------------------------------------------

* **End Point:**
/dbo/orders/get-all-order

* **Description:**
This API is used to get all orders data. It's searchable too using name key word and using pagination too

* **Method:**
GET

* **URL Query Params:**
name=...
page=...

* **URL Params:**
none

* **JSON Body:**
none

* **Success Response:**
[
    {
        "id": 11,
        "name": "bumbu bebek",
        "quantity": 3
    },
    {
        "id": 12,
        "name": "Sirloin",
        "quantity": 5
    }
]

* **Error Response:**
  * **code:** 500 <br />
  * **Content:** {
        "code": 500,
        "message": "internal server error"
    }

* **code:** 500 <br />
  * **Content:** {
        "code": 500,
        "message": "Failed fetch data"
    } 
------------------------------------------------------------------------------------------------------------------------------------------------

* **End Point:**
/dbo/orders/get-order-by-id

* **Description:**
This API is used to get detail order data by order ID

* **Method:**
GET

* **URL Query Params:**
orderId=...

* **URL Params:**
none

* **JSON Body:**
none

* **Success Response:**
{
    "id": 12,
    "name": "Sirloin",
    "quantity": 5
}

* **Error Response:**
  * **code:** 400 <br />
  * **Content:** {
        "code": 400,
        "message": "missing order id"
    }

  * **code:** 500 <br />
  * **Content:** {
        "code": 500,
        "message": "internal server error"
    }

------------------------------------------------------------------------------------------------------------------------------------------------

* **End Point:**
/dbo/orders/update-order

* **Description:**
This API is used to update order data. You need to pass order id query param to specify which order you want to update

* **Method:**
PUT

* **URL Query Params:**
orderId=...

* **URL Params:**
none

* **JSON Body:**
You can add whatever payload as long as it's property registered in order struct. The json form of the order struct is: <br/>
{
    "CreatedAt": "2024-06-16T12:14:46.5979086+07:00",
    "UpdatedAt": "2024-06-16T12:14:46.5979086+07:00",
    "DeletedAt": null,
    "ID": 12,
    "Name": "Sirloin",
    "Quantity": 1,
    "CustomerID": 7,
    "Customer": {...}
}

* **Success Response:**
{
    "code": 200,
    "message": "Success update order data"
}

* **Error Response:**
  * **code:** 400 <br />
  * **Content:** {
        "code": 400,
        "message": "missing customer id"
    }order

  * **code:** 500 <br />
  * **Content:** {
        "code": 500,
        "message": "Internal Server Error"
    }

------------------------------------------------------------------------------------------------------------------------------------------------

* **End Point:**
/dbo/orders/delete-order

* **Description:**
This API is used to delete order data. You need to pass order id query param to specify which order you want to delete

* **Method:**
DELETE

* **URL Query Params:**
custId=...

* **URL Params:**
none

* **JSON Body:**
none

* **Success Response:**
{
    "code": 200,
    "message": "Order deleted"
}

* **Error Response:**
  * **code:** 400 <br />
  * **Content:** {
        "code": 400,
        "message": "missing order id"
    }

  * **code:** 500 <br />
  * **Content:** {
        "code": 500,
        "message": "Internal Server Error"
    }

------------------------------------------------------------------------------------------------------------------------------------------------