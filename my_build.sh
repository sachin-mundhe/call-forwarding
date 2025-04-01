set -e #exit if any command return non zero status
# set -x #debug modeset
set -u #error and exit when env var is not set


# login to docker hub
echo "================== logging into docker hub =================="
docker login -u ${DOCKER_USERNAME} -p ${DOCKER_PASSWORD}

# Build docker image
echo "================== Building an docker image =================="
docker build -t sachinmundhe/call-forwarding:latest .

# push image to docker hub
echo "================== pushung image to docker hub =================="
docker push sachinmundhe/call-forwarding:latest

