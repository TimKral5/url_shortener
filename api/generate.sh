docker="docker run --rm -u $(id -u):$(id -g) -v $PWD:/local openapitools/openapi-generator-cli:v7.15.0"

yq . openapi.yaml > openapi.json
$docker generate -i /local/openapi.yaml -g go -o /local/build/go
$docker generate -i /local/openapi.yaml -g html -o /local/build/html
$docker generate -i /local/openapi.yaml -g html2 -o /local/build/html2
