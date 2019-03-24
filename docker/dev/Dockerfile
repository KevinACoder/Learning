FROM ubuntu:16.04

#########Enable SSH service
RUN apt-get update && apt-get install -y openssh-server
RUN mkdir /var/run/sshd
RUN echo 'root:docker' | chpasswd
RUN sed -i 's/PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config

# SSH login fix. Otherwise user is kicked off after login
RUN sed 's@session\s*required\s*pam_loginuid.so@session optional pam_loginuid.so@g' -i /etc/pam.d/sshd
ENV NOTVISIBLE "in users profile"
RUN echo "export VISIBLE=now" >> /etc/profile

##########Enable SSHFS service
RUN apt-get install sshfs -y
RUN apt-get install sshpass -y
RUN ln -s /config /root/.ssh
# USER nobody
COPY ./entrypoint.sh /entrypoint.sh
VOLUME ["/mount", "/config"]
ENV UID=0 GID=0 PORT=22 IDENTITY_FILE=/config/id_ed25519 SSHPASS=
ENTRYPOINT ["/entrypoint.sh"]
CMD ["chmod +x /entrypoint.sh"]

EXPOSE 22
CMD ["/usr/sbin/sshd", "-D"]