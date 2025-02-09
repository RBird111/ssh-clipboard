FROM ubuntu:24.04

RUN apt-get update && apt-get install -y openssh-server
RUN mkdir /var/run/sshd
RUN echo 'root:root' | chpasswd
RUN sed -Ei 's/#(PermitRootLogin )prohibit-password/\1yes/' /etc/ssh/sshd_config

RUN perl -le 'print int rand 100 for 1 .. 25' > ~/rand.txt

EXPOSE 22

CMD ["/usr/sbin/sshd", "-D"]
