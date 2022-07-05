docker_build:
	docker build -f dockerfile -t ssanikin/statics:0.1.0
docker_push:
	docker push ssanikin/statics:0.1.0

docker_run:
	docker run -p 8080:8080 ssanikin/statics:0.1.0