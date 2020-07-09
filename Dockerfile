FROM golang:1.14-buster as builder

WORKDIR /go

ENV GOPATH ""

COPY . ./
RUN cd cmd/analyticproxy && go build; cd -; \
    cd cmd/medifor && go build; cd -

FROM python:3.8-slim-buster
COPY --from=builder /go/python /app/bin/
RUN pip install -U pip
RUN pip install -r requirements.txt


ENV PATH ${PATH}:/app/bin
RUN mkdir -p /app/bin

COPY --from=builder /go/cmd/analyticproxy/analyticproxy /app/bin/
COPY --from=builder /go/cmd/medifor/medifor /app/bin/

EXPOSE 50051

CMD ['python', '-m', 'medifor', '--help']
