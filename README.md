# Ory Kratos Example
This is an example of using Ory Kratos with Golang.

## Getting Started
#### 1. Create a exclusive network for communication between Database and Kratos:
```bash
docker network create kratos
```

#### 2. Start a database:
```bash
docker run --name pg -p 5432:5432 --restart always --network kratos -e POSTGRES_DB=kratos -e POSTGRES_PASSWORD=CHANGE_PASSWORD -d postgres:13.4
```

#### 3. Build a container image with custom configuration:
```bash
docker build -f ./scripts/kratos.dockerfile -t ory-kratos-example:0.1 .
```

#### 4. Run Kratos Migration (create tables on database):
```bash
docker run --network kratos --rm ory-kratos-example:0.1 migrate sql "postgres://postgres:CHANGE_PASSWORD@pg:5432/kratos" --yes
```

#### 5. Run Kratos:
```bash
docker run --name kratos -d --network kratos -p 4000:4000 -p 4001:4001 --restart always ory-kratos-example:0.1 serve --config /home/ory/kratos.yaml
```

#### 6. Run Application:
```bash
go run main.go web
```

#### 7. Open the web application:
[http://127.0.0.1:4002/](http://127.0.0.1:4002/)

## Clean Up Environment
For clean the environment:
```bash
docker rm -f kratos
docker rm -f pg
docker rmi ory-kratos-example:0.1
docker network rm kratos
```