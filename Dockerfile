FROM ubuntu:latest
COPY akariai /opt/akariai/akariai
COPY assets /opt/akariai/assets
RUN apt-get update && \
    apt-get install -y ca-certificates && \
    rm -rf /var/lib/apt/lists/* && \
    update-ca-certificates && \
    useradd akariai && \
    chown akariai:akariai -R /opt/akariai && \
    chmod 700 /opt/akariai/akariai
USER akariai
WORKDIR /opt/akariai
CMD ["/opt/akariai/akariai"]
