#!/bin/bash

echo "Running '$0' for '${NAME}'"

dev() {
	echo "minikube"
	docker save ${IMG} | (eval $(minikube docker-env) && docker load)
	kubectl apply -f tmp/k8s.yml
}

cloud() {
	### The following section will push your container image to ECR. The `$NAME` variable is provided from our
	### Makefile under 'deploy:' rule, which is set to the name of the component/module/service.
	###
	docker tag ${NAME}:${CI_COMMIT_SHORT_SHA} $CI_REGISTRY/crypto-game-portal/monorepo/${NAME}:${CI_COMMIT_SHORT_SHA}
	docker tag ${NAME}:${CI_COMMIT_SHORT_SHA} $CI_REGISTRY/crypto-game-portal/monorepo/${NAME}:latest
	docker images
	docker push $CI_REGISTRY/crypto-game-portal/monorepo/${NAME}:${CI_COMMIT_SHORT_SHA}
	docker push $CI_REGISTRY/crypto-game-portal/monorepo/${NAME}:latest
	
	### Apply the changes to our cluster
	
	# echo "kubectl apply..."
	# kubectl apply -f backend/${NAME}/k8s.yml	
}

if [ ${NAME} == "local" ]
then
	dev
else
	cloud
fi
