# node

## ??nats?????ws

```
go run ws-tcp-relay.go
```

## node connection gnatsd
```
const NATS = require('nats')
const nc = NATS.connect('nats://testservicenats.myun.tv:4222')
nc.publish('foo', 'Hello World!')
nc.subscribe('foo', function(msg) {
    console.log('Received a message: ' + msg)
})
```

## node connection websocket-nats