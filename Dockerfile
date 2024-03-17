FROM scratch
COPY ./blackhole /
EXPOSE 8080
EXPOSE 2525
WORKDIR /
ENTRYPOINT ["/blackhole"]
