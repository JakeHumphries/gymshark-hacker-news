## Hacker News Onboarding Project

This is a project to help get up to speed with the WoW, stack and tooling the gymshark software team uses. :rocket:  

It is a monorepo that is comprised of the following services:

- **RabbitMQ Publisher Service:** to get information from the Hacker News API and publish messages to a queue. It pulls from the hackernews api periodically via a cron job. 
- **RabbitMQ Consumer Service:** to spawn worker processes to read from the queue and save to a database.
- **API Service:** to forward requests to GRPC service 
- **GRPC Service:** to read and return items from the database

:warning: *Although this project has got some unit tests it is by no means fully tested*:warning:

## Running Locally

run `docker-compose up` to initialize all micro services, a mongo instance, a redis instance, and a rabbitMQ instance

## API

#### Get healthz

Endpoint `/_healthz`

This accepts GET requests. Here is an example development request:

```bash
curl 'http://0.0.0.0:8000/_healthz' \
  -H 'Content-Type: application/json' -v
```

Example response:

```json
"ok"
```

#### Get all

Endpoint `/all`

This accepts GET requests. Here is an example development request:

```bash
curl 'http://0.0.0.0:8000/all' \
  -H 'Content-Type: application/json' -v
```

Example response:

```json
[
    {
        "by": "OrangeTux",
        "title": "Nobel Prize in Physiology or Medicine 2021",
        "url": "https://www.nobelprize.org/prizes/medicine/2021/press-release/",
        "text": "",
        "type": "story",
        "descendants": 110,
        "id": 28745101,
        "score": 212,
        "time": 1633344129,
        "parent": 0,
        "poll": 0,
        "kids": [
            28745277,
            28745403,
            28745264,
            28746006,
            28746466,
            28746087,
            28745363,
            28745334,
            28753783,
            28779080,
            28745249
        ],
        "parts": null,
        "deleted": false,
        "dead": false
    },
    {
        "by": "thomaspaulmann",
        "title": "Raycast (YC W20) Is Hiring Product Designers",
        "url": "https://www.raycast.com/jobs/product-designer",
        "text": "",
        "type": "job",
        "descendants": 0,
        "id": 28743902,
        "score": 1,
        "time": 1633330807,
        "parent": 0,
        "poll": 0,
        "kids": null,
        "parts": null,
        "deleted": false,
        "dead": false
    }
]
```

#### Get stories

Endpoint `/stories`

This accepts GET requests. Here is an example development request:

```bash
curl 'http://0.0.0.0:8000/stories' \
  -H 'Content-Type: application/json' -v
```

Example response:

```json
[
    {
        "by": "OrangeTux",
        "title": "Nobel Prize in Physiology or Medicine 2021",
        "url": "https://www.nobelprize.org/prizes/medicine/2021/press-release/",
        "text": "",
        "type": "story",
        "descendants": 110,
        "id": 28745101,
        "score": 212,
        "time": 1633344129,
        "parent": 0,
        "poll": 0,
        "kids": [
            28745277,
            28745403,
            28745264,
            28746006,
            28746466,
            28746087,
            28745363,
            28745334,
            28753783,
            28779080,
            28745249
        ],
        "parts": null,
        "deleted": false,
        "dead": false
    }
]
```

#### Get jobs

Endpoint `/jobs`

This accepts GET requests. Here is an example development request:

```bash
curl 'http://0.0.0.0:8000/jobs' \
  -H 'Content-Type: application/json' -v
```

Example response:

```json
[
    {
        "by": "thomaspaulmann",
        "title": "Raycast (YC W20) Is Hiring Product Designers",
        "url": "https://www.raycast.com/jobs/product-designer",
        "text": "",
        "type": "job",
        "descendants": 0,
        "id": 28743902,
        "score": 1,
        "time": 1633330807,
        "parent": 0,
        "poll": 0,
        "kids": null,
        "parts": null,
        "deleted": false,
        "dead": false
    }
]
```

## Environment Variables

There are a number of environment variables that are required within the .env file in the root dir. 

To specify where the application runs:

*   API_HOST - The host for the application, defaults to 0.0.0.0
*   API_PORT - The port for the application, defaults to 8000

To specify for the database:

*   DB_PORT - The port the database instance is running on
*   DB_NAME - The name of the database
*   DB_USER - The username to access the database
*   DB_PASSWORD - The password to access the database

To specify for the redis instance:

*   REDIS_HOST - The host url for the redis instance

To specify for the rabbitMQ instance: 

*   RABBIT_MQ_USER - The username to access the queue
*   RABBIT_MQ_PASS - The password to access the queue
*   RABBIT_MQ_HOST - The host url for the rabbitMQ instance

