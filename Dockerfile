# ベースとなるDockerイメージ指定
FROM golang:latest
# コンテナ内に作業ディレクトリを作成
RUN mkdir /go/src/go_test
# コンテナログイン時のディレクトリ指定
WORKDIR /go/src/go_test
# ホストのファイルをコンテナの作業ディレクトリに移行
COPY . /go/src/go_test