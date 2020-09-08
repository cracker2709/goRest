# goRest

- Get docker image with

```shell script
docker pull cracker2709/public-restful-go
```
- Run if with
```shell script
docker run -p 8080:8080 cracker2709/public-restful-go:latest
```
Install it quickly on a kubernetes cluster
```
# Launch a temporary pod which will be destroyed when exiting the session
kubectl run tmp-go-rest-pod --rm -i --tty --image cracker2709/public-restful-go:latest --namespace <your_namespace>

# Port forward it so you are able to browse it locally
kubectl port-forward tmp-go-rest-pod 8080:8080
```

- Browse through a web browser or better with postman
```
http://localhost:8080/
http://localhost:8080/api/v1
http://localhost:8080/api/v1/user/1/comment/2
```
