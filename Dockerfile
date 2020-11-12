FROM alpine

RUN mkdir -p /files /etc/treasury && echo "{\"reads\": [\"/files\"], \"writes\": [\"/files\"]}" > /etc/treasury/config.json
ADD build/treasury /bin/treasury

CMD ["/bin/treasury"]
