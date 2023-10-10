FROM debian:stable-slim

ENV PORT 8080

COPY rss-aggregator /bin/rss-aggregator

CMD ["/bin/rss-aggregator"]
