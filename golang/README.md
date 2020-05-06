# websocket

- server
当接收到连接请求后，将连接使用的http协议升级为websocket协议。后续通信过程中，使用websocket进行通信。对每个连接，server端等待读取数据，读到数据后，打印数据，然后，将数据又发送给client.