# Calculating the price of the trip

Run docker compose:
````
docker-compose up
````

Request:
````
curl -X POST http://localhost:3300/api/v1/prices/calculate \
     -H "Content-Type: application/json" \
     -d '{
            "start_point":{
                "lat":55.739060,
                "lng":37.622691
            },
            "end_point":{
                "lat":55.808116,
                "lng":37.581609
            }
        }'
````

Response:
````
{"price":3997.0}
````

Default tariff data:
````
min_price: 500.0
order_price: 250.0
minute_price: 20.0
km_price: 20.0
````

Calculation formula
````
order_price + minute_price * time_in_minutes + km_price * distance_in_km
````

You can override tariff data:

Request:
````
curl -X POST http://localhost:3300/api/v1/prices/calculate \
     -H "Content-Type: application/json" \
     -d '{
            "start_point":{
                "lat":55.739060,
                "lng":37.622691
            },
            "end_point":{
                "lat":55.808116,
                "lng":37.581609
            },
            "tariff":{
                "min_price":1000.0,
                "order_price":1250.0
            }
        }'
````

Response:
````
{"price":4997.0}
````
