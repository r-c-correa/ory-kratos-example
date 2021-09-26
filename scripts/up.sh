docker network create kratos

docker run --name pg -p 5432:5432 --restart always --network kratos -e POSTGRES_DB=kratos -e POSTGRES_PASSWORD=CHANGE_PASSWORD -d postgres:13.4

docker build -f ./scripts/kratos.dockerfile -t ory-kratos-example:0.1 .

docker run --network kratos --rm ory-kratos-example:0.1 migrate sql "postgres://postgres:CHANGE_PASSWORD@pg:5432/kratos" --yes

docker run --name kratos -d --network kratos -p 4000:4000 -p 4001:4001 --restart always ory-kratos-example:0.1 serve --config "/home/ory/kratos.yaml"