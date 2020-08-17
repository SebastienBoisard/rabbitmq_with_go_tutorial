# RabbitMQ with Go tutorial
A simple RabbitMQ tutorial with Go


## Run RabbitMQ

For this tutorial, we'll use RabbitMQ from a Docker image.


Pull the latest RabbitMQ image from the Docker Hub repository (Cf. https://hub.docker.com/_/rabbitmq)
```
> docker pull rabbitmq
```

Run a detached instance of the RabbitMQ image ( named `rabbitMQ-tutorial`) listening on the default port of 5672 
with default login/password guest/guest
```
> docker run -d --hostname rabbitmq-tutorial-hostname --name rabbitmq-tutorial -p 5672:5672 rabbitmq
```

To run an instance with others credentials  
```
> docker run -d --hostname rabbitmq-tutorial-hostname --name rabbitmq-tutorial -p 5672:5672  \
  -e RABBITMQ_DEFAULT_USER=user -e RABBITMQ_DEFAULT_PASS=password rabbitmq
```

## Sending/Receiving messages



Open a bash on the docker instance
```
> docker exec -ti rabbitmq-tutorial  bash
```

List all the open queues
```
> rabbitmqctl list_queues

Timeout: 60.0 seconds ...
Listing queues for vhost / ...
name	messages
hello	1
```
 