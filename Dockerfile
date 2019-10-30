FROM scratch

ENV ADDR 0.0.0.0:8080
ENV KEY test-key

COPY "bin/genpass.bin" /
ENTRYPOINT ["/genpass.bin"]
