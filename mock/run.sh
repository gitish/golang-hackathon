docker run -it --rm \
  -p 9080:8080 \
  --name wiremock \
  -v $PWD:/home/wiremock \
  wiremock/wiremock:2.35.0