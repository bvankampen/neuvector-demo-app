FROM golang:alpine
RUN mkdir /code
RUN mkdir /webapp

COPY ./cmd /code/cmd
COPY ./static /webapp/static
COPY ./templates /webapp/templates
WORKDIR /code
RUN go build -o /webapp/frontend /code/cmd/frontend/main.go
RUN go build -o /webapp/backend /code/cmd/backend/main.go
RUN rm -rf /code
