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

# mount docker filesystem to host http://programster.blogspot.com/2014/08/dev-inside-docker-container-with-sshfs.html
docker run --privileged -d -P --name test_sshd -p 10122:22 eg_sshd
sudo sshfs -o allow_other root@192.168.0.125:10122:/ /home/pi/Code/learning/docker/dev_envir/share1

docker plugin install vieux/sshfs
docker volume create -d vieux/sshfs -o sshcmd=<pi@192.168.0.125:/home/pi/Code/learning/docker/dev_envir/share2> -o password= [-o port=] sshvolume sshvolume
docker volume create -d vieux/sshfs -o sshcmd=<pi@192.168.0.125:/home/pi/Code/learning/docker/dev_envir/share2> -o password=kevin123 [-o port=22] sshvolume sshvolume
docker volume create -d vieux/sshfs -o sshcmd=<user@host:path> -o password=<password> [-o port=<port>] [-o <any_sshfs_-o_option> ] sshvolume sshvolume
docker volume ls DRIVER VOLUME NAME local 
docker run -it -v sshvolume: eg_sshd ls

# use volume to share data
docker volume create hello
docker volume create --driver local \
    --opt type=nfs \
    --opt o=addr=192.168.0.125,rw \
    --opt device=:/home/pi/Code/learning/docker/dev_envir/share2 \
    hello3
docker volume ls
docker volume rm hello3
docker run -d -v hello3:/world3 eg_sshd ls /world3

docker volume create -d local data
ls -l /var/lib/docker/volumes
docker run -d -P --name dev_env -v /data:/opt/data eg_sshd

docker run -d --rm \
    --cap-add SYS_ADMIN \
    --device /dev/fuse \
    --name dev_env \
    -e UID=1000 \
    -e GID=100 \
    -v $PWD/mnt:/mount:shared \
    -v ~/.ssh/id_ed25519:/config/id_ed25519:ro \
    eg_sshd \
    root@192.168.0.125:/data