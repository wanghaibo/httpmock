From alpine
ADD https://github.com/wanghaibo/httpmock/releases/download/v1.0.0/mock /usr/local/bin/mock
RUN chmod u+x /usr/local/bin/mock
ENTRYPOINT ["/usr/local/bin/mock", "--datapath", "/data/mock"]
