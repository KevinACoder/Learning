docker create -it python:3
docker start python:3

# run container in daemonized
docker run -d --name test mubuntu:1 /bin/sh -c "while true; do echo hello world; sleep 1; done"
# ls container info
docker ps
# get output of container, continue output
#  with time stamp
docker logs -f -t test
# pause container
docker pause test
# send SIGTERM and wait for 10 sec and
#   send SIGKILL
docker stop test

# attach container
docker run -itd --name test ubuntu
docker attach test
# -it maintain input stream
docker exec -it test /bin/bash

# remove container
docker ps -a
docker stop test
docker rm test
# force to remove container
docker rm -f test

# inspect container info
docker container inspect test
# inspect process inside container
docker top test
# get stats of container resource usage
docker stats test