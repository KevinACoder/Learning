# copy file from host to container
docker container cp python3.tar test:/tmp/
# cmp diff in file sys
docker container diff test
