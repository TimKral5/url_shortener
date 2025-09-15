docker="docker run --rm -u 1000:1000 -v $PWD:/local openapitools/openapi-generator-cli"

yq . openapi.yaml > openapi.json
$docker generate -i /local/openapi.yaml -g go -o /local/build/go
$docker generate -i /local/openapi.yaml -g html -o /local/build/html
$docker generate -i /local/openapi.yaml -g html2 -o /local/build/html2
