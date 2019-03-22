# get system info
uname -a #core info
cat /proc/version

# manual install for raspberry pi
sudo apt-get update && sudo apt-get upgrade -y
curl -sSL https://get.docker.com | sh
#add permission to pi user
sudo usermod -aG docker pi
docker run armhf/hello-world

#Removes docker and dependencies
sudo apt-get remove --auto-remove docker
#Removes all data 
sudo rm -rf /var/lib/docker 