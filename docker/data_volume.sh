# create data volume
docker volume create -d local data
sudo ls -l /var/lib/docker/volumes
# check volume details
docker volume inspect data
#docker run -d -P --name web --mounttype=bind,source=/webapp,destination=/opt/webapp training/webapp python app.python
docker run -d -P --name web -v webapp:/opt/webapp training/webapp python app.py