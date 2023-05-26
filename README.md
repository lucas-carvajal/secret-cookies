# Secret Cookies
A simple API to store and retrieve cookie recipes.
The endpoints are protected by **JWT Authentication** and the accounts
are managed by **Firebase Authentication**.
The API is realised with **Gin** and the data is persisted in **MongoDB**.

## Endpoints
`POST /api/signup`  
Create a new user account. The request body should contain the following fields:
- `email` (string)
- `password` (string)
- `usermame` (string)

`POST /api/login`  
Login with an existing user account. The request body should contain the following fields:
- `email` (string)
- `password` (string)  

The response contains a JWT token that should be used for authentication.


`GET /api/recipes`  
Requires read permission.
Get all recipes. The response contains an array of recipes.

`GET /api/recipes/:id`  
Requires read permission.
Get a recipe by id. The response contains a single recipe.

`POST /api/recipes`  
Requires write permission.
Create a new recipe. The request body should contain the following fields:
- `name` (string)
- `ingredients` (map of strings)
- `instructions` (ordered array of strings)
- `notes` (string)

`PUT /api/recipes/:id`  
Requires write permission.
Update a recipe by id. The request body should contain the following fields:
- `name` (string)
- `ingredients` (map of strings)
- `instructions` (ordered array of strings)
- `notes` (string)

