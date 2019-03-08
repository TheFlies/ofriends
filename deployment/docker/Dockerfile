FROM alpine:3.8

WORKDIR /home/
COPY ofriends.bin .
RUN chmod +x ofriends.bin

COPY web/dist ./web

EXPOSE 8080
CMD ["./ofriends.bin"]
