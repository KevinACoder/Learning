# build image through docker file
docker image build -t python:3 .
# check newly build image info
docker images|grep python

# export image
docker save -o python3.tar python:3
# import image
docker load -i python3.tar