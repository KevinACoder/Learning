# set up ssh # https://docs.docker.com/engine/examples/running_ssh_service/
# build image with Dockerfile
docker build -t eg_sshd .
# run container with assigned port
docker run -d -P --name test_sshd -p 10122:22 eg_sshd
# remote connect to container
ssh root@192.168.0.125 -p 10122
# remove container
docker container stop test_sshd
docker container rm test_sshd