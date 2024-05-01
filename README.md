# api-shop

### How to use

#### Read product

In GET request is required send `id`.
```
curl --location 'http://localhost:4000/products/1'
```

#### Read All products
```
curl --location 'http://localhost:4000/products'
```

#### Insert product
```
curl --location 'http://localhost:4000/products/' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Camisa",
    "description": "Uma camisa estilosa e bonita",
    "price": 209.95
}'
```

#### Update product
```
curl --location 'http://localhost:4000/products/1' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Regata",
    "description": "Uma regata estilosa e bonita",
    "price": 209.95
}'
```

#### delete product
```
curl --location --request DELETE 'http://localhost:4000/products/1'
```