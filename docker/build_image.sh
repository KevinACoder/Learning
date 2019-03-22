# build image through docker file
docker image build -t python:3 .
# check newly build image info
docker images|grep python