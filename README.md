# Housing gRPC

This is a gRPC microservice that performs Break-Even analysis

## Estimation Parameters

For the purpose of this example, we will store our estimation parameters as environment variables.
You can set your preferred values here. All environment variables are store in the `.env.sample` file.

NOTE: Ensure you created a `.env` file in the root directory and add these variables to the file

Sample estimation parameters are defined as follows:

```
PROPERTY_APPRECIATION_RATE=2.00
ANNUAL_INFLATION=2.40
CLOSING_COST=1.5
REALTOR_FEES=3
LOAN_DURATION=30
LOAN_INTEREST_RATE=5
ANNUAL_RENT_INCREASE_RATE=2
PROPERTY_TAX=2.40
PROPERTY_TRANSFER_TAX_RATE=1
MAINTENANCE_RATE=1
UTILITIES_RATE=0.05
HOME_OWNER_ASSOCIATION_RATE=0.25
HOME_OWNER_ASSOCIATION_INSURANCE_RATE=0.5
```

Other environment variables (with sample values) needed are as follows:

```
GO_ENV=development
APP_PORT=8080
CALCULATOR_SERVICE_PORT=7990
```

## Run gRPC service and REST API gateway

To run the calculator gRPC service:

```
$ go run grpc/server/*.go
```

To run the REST API gateway:

```
$ go run gateway/app.go
```

Sample Request:

```
curl http://localhost:8080/v1/breakeven/compute
-H "Content-Type: application/json"
-D '{ "propertyValue": 300000.00, "initialDepositRate": 20.00, "monthlyRent": 3000.00, "occupancyDuration": 8 }'
-X POST
```

Sample response:

```
{
  "success": true,
  "data": {
    "message": "If you stay in this property for 8 years, buying is cheaper than renting.",
    "purchase": 322490.3,
    "rent": 345477.38,
    "verdict": "Year 7 is the break-even year"
  }
}
```