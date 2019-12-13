# Go x Kafka

Kafka is used for building real-time data pipelines and streaming apps. It is horizontally scalable, fault-tolerant, wicked fast, and runs in production in thousands of companies.

We will implement Kafka x Golang using Docker.
 
#### Tech & Requirements

* [Golang]  - version 1.13.4
* [Docker] - version 19.03.2
* [Kafka] - using docker images
* [Zookeeper] - using docker images
** [Golang Library] - github.com/Shopify/sarama
** [Golang Library] - github.com/sirupsen/logrus

### Setup Go x Kafka

`Install Golang` 
`Install Docker`

Clone this repo on your machine.
```sh
git clone https://github.com/naufalziyad/Go-Kafka.git
```

Create docker network, this example create docker network 'kafka'
```sh 
docker network create kafka
```

Run Zookeeper
```sh 
docker run --net=kafka -d --name=zookeeper -e ZOOKEEPER_CLIENT_PORT=2181 confluentinc/cp-zookeeper:4.1.0
```

Run Apache Kafka image
```sh 
docker run --net=kafka -d -p 9092:9092 --name=kafka -e KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://kafka:9092 -e KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1 confluentinc/cp-kafka:4.1.0
```

Add your DNS kafka
```sh 
127.0.0.1 kafka
```

Run Services Consumer & Publisher
* Terminal 1 Consumer
```sh
cd [Repo Project]/action/consumer
```
```sh
go Run main.go`
```

* Terminal 2 Publisher
`cd [Repo Project]/action/publisher`
`go Run main.go`

### Result
![alt text](https://github.com/naufalziyad/Go-Kafka/blob/master/img/kafka-naufal.gif)


#### Thanksfull
https://medium.com/easyread/implementasi-kafka-menggunakan-golang-testing-db183e0b3c29
https://kafka.apache.org/
https://golang.org/
