FROM oryd/kratos:v0.7.6
COPY ./scripts/kratos.yaml /home/ory
COPY ./scripts/kratos.identity.schema.json /home/ory