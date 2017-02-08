From alpine
ADD https://github.com/wanghaibo/httpmock/releases/download/v1.0.0/mock /mock
RUN chmod u+x /mock
ENTRYPOINT ["/mock", "--datapath", "/mock.txt"]
