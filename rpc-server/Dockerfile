FROM golang:1.18
WORKDIR /app
COPY . .
RUN sh ./build.sh
EXPOSE 8888
RUN wget -O /usr/local/bin/wait-for-it.sh https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh
RUN chmod +x /usr/local/bin/wait-for-it.sh
ENTRYPOINT [ "bash", "/app/entrypoint.sh" ]
# CMD ["./output/bootstrap.sh"]