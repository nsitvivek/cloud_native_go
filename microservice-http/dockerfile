FROM alpine:3.5
MAINTAINER  M.-nsitvivek 
COPY ./cloud_native_go /app/cloud_native_go
RUN chmod +x /app/cloud_native_go
ENV PORT 8080
EXPOSE 8080
ENTRYPOINT app/cloud_native_go