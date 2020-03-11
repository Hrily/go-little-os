FROM i386/debian:stable-slim
WORKDIR /littleosbook
RUN  apt-get update
RUN  apt-get install gcc gccgo nasm make -y
RUN  apt-get install gcc-multilib -y
RUN  apt-get install g++-multilib libc6-dev -y
RUN  apt-get install genisoimage -y
COPY . /littleosbook
RUN  make clean all
