# notification server

This example demonstrates a more real-world application consisting of multiple microservices.

## Description

Used microservices architecture.

The implementation is based on the [Domain Driven Design](http://www.amazon.com/Domain-Driven-Design-Tackling-Complexity-Software/dp/0321125215) book by Eric Evans.

### Organization

The project consists of multiple folders.
- __core__ - contains domain packages that contain some intricate business-logic.
- __application__ - contains micro services.
- __inmem__ - contains in-memory implementations for the repositories found in the domain packages.
- __cmd__ - contains commands that used to start services.

The application folder consists of three micro services. 

- __publishing__ - used by the user to publish messages.
- __managing__ - used by our to retrieve user's messages, update message state(set 'read' or 'unread') etc.
- __subscriptioning__ - used by user to subscription category of messages.



## Usage

Install dependencies & run application.

```
dep ensure

cd cmd/notificationserverd
go run main.go

```

Publish message.
```
curl -X POST \
  http://localhost:8080/publishing/v1/publish \
  -H 'Content-Type: application/json' \
  -d '{
	"notificationName": "create",
	"message": "{'\''msg'\'': '\''wiki'\''}",
	"severity": "info",
	"userIds": [1, 2],
	"excludedUserIds": []
}'
```

Get user notifications.
```
curl -X GET http://localhost:8080/managing/v1/user/1
```