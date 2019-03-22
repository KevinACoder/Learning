# get image
docker image ls
docker pull ubuntu:18.04
docker pull busybox

# list images
docker image ls
docker images
docker ps -a #list all local containers

# run image
docker run -it ubuntu:18.04 bash
# run app in container
docker run mubuntu echo "hello docker"

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
docker search --filter=is-official=true --filter=stars=4 ubuntu

# remove images
docker image rm ubuntu:18.04
# force to delete an image even if
#  it is running
docker image rm -f ubuntu:18.04
# delete tmp image file layers
docker image prune -f

# modify the container and save changes
docker run -it mubuntu /bin/bash
# commit container by id,
#   commit with comments and author name
docker container commit -m "add new folder" -a "Kevin" cd6a1fb8b715 mubuntu:1