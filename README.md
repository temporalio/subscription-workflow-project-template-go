# Temporal Subscription Workflow Template - Go

Temporal customer subscription Workflow example. 

### Setup

#### Run Temporal server

```bash
git clone https://github.com/temporalio/docker-compose.git
cd docker-compose
docker-compose up
```

#### Start the example

Start the Worker:

```text
go run worker/main.go
```

Start the Workflow executions.
This will start the Subscription Workflow for 5 customers with ids:

```text
Id-0
Id-1
Id-2
Id-3
Id-4
```

```text
go run starter/main.go
```

##### Querying billing information:

You can query billing information for all customers after the workflows have started with:

```text
go run querybillinginfo/main.go    
```
This will return the current Billing Period and the current Billing Period Charge amount for each of the customers.

You can run this multiple times to see the billing period number increase during 
workflow execution