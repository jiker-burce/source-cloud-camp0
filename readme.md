## cloud tec study

### homework/http_server_simple

- run

```
cd homework/http_server_simple

// we should login docker before releasing our image.
docker login

make release

// you can release the image using the following command if you can't find it in the https://hub.docker.com/u/your_user_name
docker push storefeegmail/httpserver:v1.0

docker run --publish 9090:9090 storefeegmail/httpserver:v1.0

curl http://localhost:9090/healthz
```