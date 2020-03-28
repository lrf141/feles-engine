MAKE             = make
GO               = go
CC               = gcc
INSTALL          = install
SHARED           = -shared
C_TARGET         = *.c
LIB_NAME         = libfeles.so
OUTPUT_DIR       = output_library_dir
BINARY           = feles-engine
BUILD_LIB        = build_shared_object
BUILD_DAEMON     = build_go_server
INSTALL_DIR      = /usr/bin/local

all:	$(BUILD_LIB) $(BUILD_DAEMON)

$(BUILD_DAEMON):
	$(GO) build .

$(BUILD_LIB):
	$(CC) $(SHARED) $(C_TARGET) -o $(OUTPUT_DIR)/$(LIB_NAME)

install:
	$(INSTALL) $(BINARY) $(INSTALL_DIR)

uninstall:
	rm $(INSTALL_DIR)/$(BINARY)

clean:
	rm $(BINARY)
	rm $(OUTPUT_DIR)/*.so