go test ./...
if ($LASTEXITCODE -eq 0) 
{
    docker build -f Dockerfile -t lykalon/app:v1 .
    docker push lykalon/app:v1
    docker-compose up app -d
}