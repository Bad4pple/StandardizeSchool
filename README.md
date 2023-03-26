# Order Service

## **How to run**
```
docker-compose up -d 
```
## **How to stop**
```
docker-compose down -d 
```
## **Docs**

## **To create new order**
```js
POST http://localhost:8000/api/orders

{
  "customer": {
    "customer_id": "292929292",
    "organization_name": "Standardize School",
    "phone": "0987654321",
    "email": "standardize.school@stds.com",
    "social_links": ["https://www.facebook.com/standardize.group"],
    "address": "Chiang Mai"
  },
  "options": [
    {
      "option_id": "option-abc123",
      "price": 255
    },
    {
      "option_id": "option-def456",
      "price": 2
    }
  ]
}
``` 
## **Submission**
```js
POST http://localhost:8000/api/orders/{order_id}/submission
```

## **Get Order**
```js
GET http://localhost:8000/api/orders/{order_id}
```

## **Update Options**
```js
PUT http://localhost:8000/api/orders/{order_id}
{
    options: [
        {
             "option_id": "option-def456",
             "price": 2
      }
    ]
}

```