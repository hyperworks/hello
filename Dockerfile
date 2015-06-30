FROM scratch
ADD hello /hello
EXPOSE 8080
CMD ["/hello"]
