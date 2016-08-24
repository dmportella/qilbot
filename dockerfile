FROM scratch
MAINTAINER Daniel Portella

ARG CONT_IMG_VER
ARG BINARY

LABEL version ${CONT_IMG_VER}

ADD ${BINARY} /qilbot

ENTRYPOINT ["/qilbot"]