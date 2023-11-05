# Kafka_in_Go
Apache Kafka is a distributed data store for processing streaming data in real time.
The streaming platform needs to process these data continous inflowing data and process them step by step in sequence.

Kafka operates publish-subscribe model.
Producers publish data and consumers subscribe receive and process these data.

Kafka is primarily used to build  real time streaming data pipelines and applications that adapt to data streams.

docker compose up -d
cd /consumer/cmd -> go run main.go
cd /producer/cmd -> go run main.go