Variables

```bash
export REDIS=localhost:6379
export RMQ=amqp://guest:guest@localhost:5672/todo
export RMQ_API=http://guest:guest@localhost:15672/api/
```

Run

```bash
go build && ./create
```