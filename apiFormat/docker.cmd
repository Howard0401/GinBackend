docker pull jaegertracing/all-in-one

docker run -d --name jaeger -e COLLEECTOR_ZIPKIN_HTTP_PORT=9411 -p 5775:5775/udp -p 6831:6831/udp -p 6832:6832/udp -p 5778:5778 -p 16686:16686 -p 14268:14268 -p 9411:9411 jaegertracing/all-in-one

go get -u github.com/opentracing/opentracing-go
go get -u github.com/uber/jaeger-client-go

報錯；if pull access denied for jargertracing/all-in-one, repository does not exist or may require 'docker login': denied: requested access to the resource is denied:
重新登錄
docker logout
docker login
https://stackoverflow.com/questions/41984399/denied-requested-access-to-the-resource-is-denied-docker