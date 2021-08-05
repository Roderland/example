SRC_DIR="."
DST_DIR="./generate"

protoc -I=$SRC_DIR \
--go_out=$DST_DIR --go_opt=paths=source_relative \
--go-grpc_out=$DST_DIR --go-grpc_opt=paths=source_relative \
$SRC_DIR/addressbook.proto