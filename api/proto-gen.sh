#!/bin/sh

PROTO_MODULE="."
BACKEND_API="../backend/api"
FRONTEND_API="../frontend/api/src"
BACKEND_OUT="$GOPATH/src"
PROTO_GEN_VALIDATE_LIB="$GOPATH/src/github.com/envoyproxy/protoc-gen-validate"

mkdir -p "$BACKEND_API"
mkdir -p "$FRONTEND_API"

MODULES=$(ls "$PROTO_MODULE")

for NAME in ${MODULES}; do
  MODULE="$PROTO_MODULE/$NAME"
  if [ -d "$MODULE" ]
  then
    if [ "$NAME" = "$1" ] || [ -z "$1" ]
    then
      PROTO_FILES=$(find "$MODULE" -name "*.proto" -not -path "*/vendor/*")
      for PROTO_FILE in ${PROTO_FILES}; do
        case "$NAME" in
          # Handle generate for proto in config directory 
          config)
            protoc -I. -I"${PROTO_GEN_VALIDATE_LIB}" \
            --go_out="$BACKEND_OUT" \
            --validate_out="lang=go:$BACKEND_OUT" \
            "${PROTO_FILE}"
            ;;
          # Handle generate for proto in others directory
          *)
            protoc -I. -I"${PROTO_GEN_VALIDATE_LIB}" \
              --go_out="$BACKEND_OUT" \
              --go-grpc_out="$BACKEND_OUT" \
              --validate_out="lang=go:$BACKEND_OUT" \
              --js_out=import_style=commonjs:"$FRONTEND_API" \
              --grpc-web_out=import_style=commonjs,mode=grpcwebtext:"$FRONTEND_API" \
              "${PROTO_FILE}"
            ;;
        esac
        echo "protoc $PROTO_FILE ---> DONE"
      done
    fi
  fi
done

