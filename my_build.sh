set -e #exit if any command return non zero status
# set -x #debug modeset
set -u #error and exit when env var is not set

# login to docker hub
echo "================== logging into docker hub =================="
docker login -u ${DOCKER_USERNAME} -p ${DOCKER_PASSWORD}

# Build docker images
echo "================== Building an docker image =================="

tags=("latest" "${GITHUB_SHA:0:6}")
for tag in "${tags[@]}"; do
    echo "Building docker image with tag ${tag}"
    docker build -t ${DOCKER_USERNAME}/${DOCKER_IMAGE_NAME}:$tag .
done

# push images to docker hub
echo "================== pushung image to docker hub =================="

for tag in "${tags[@]}"; do
    echo "Pushing docker image with tag ${tag}"
    docker push ${DOCKER_USERNAME}/${DOCKER_IMAGE_NAME}:$tag
done