# get image
docker image ls
docker pull ubuntu:18.04
docker pull busybox

# list images
docker image ls
docker images

# run image
docker run -it ubuntu:18.04 bash

# check manual
man docker-images

# tag image
docker tag ubuntu:18.04 mubuntu

# inspect image
docker image inspect mubuntu
# check one item of image info
docker image inspect -f {{".Architecture"}} mubuntu

# list info of each layer
docker history mubuntu --no-trunc

# search on-line images
docker search --filter=is-official=true ubuntu