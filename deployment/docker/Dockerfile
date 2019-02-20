FROM alpine:3.8

WORKDIR /home/
COPY ofriends.bin .
RUN chmod +x ofriends.bin

# COPY web/assets ./web/assets
COPY web/dist ./web
# COPY web/index.html ./web/index.html

EXPOSE 8080
CMD ["./ofriends.bin"]
