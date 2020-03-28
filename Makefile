MAKE = make
GO = go
CC = gcc
SHARED = -shared
C_TARGET = *.c
LIB_NAME = libfeles.so
OUTPUT_DIR = output_library_dir
EXEC_NAME = feles-engine
BUILD_LIB = build_shared_object
BUILD_DAEMON = build_go_server

all:	$(BUILD_LIB) $(BUILD_DAEMON)

$(BUILD_DAEMON):
	$(GO) build .

$(BUILD_LIB):
	$(CC) $(SHARED) $(C_TARGET) -o $(OUTPUT_DIR)/$(LIB_NAME)

clean:
	rm $(EXEC_NAME)
	rm $(OUTPUT_DIR)/*.so