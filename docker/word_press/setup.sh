# install docker-compose
sudo curl -L "https://github.com/docker/compose/releases/download/1.23.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
#for arm https://dietpi.com/phpbb/viewtopic.php?f=11&t=4472
sudo curl -L "https://github.com/javabean/arm-compose/releases/download/1.21.2/docker-compose-Linux-armv7l" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
docker-compose --version
sudo rm /usr/local/bin/docker-compose

# build project with compose file
# https://docs.docker.com/compose/wordpress/s
docker-compose up -d
docker-compose down