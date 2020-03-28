CC = gcc
SHARED = -shared
C_TARGET = *.c
LIB_NAME = libfeles.so
OUTPUT_DIR = output_library_dir


build_shared_object:	$(LIB_NAME)
	$(CC) $(SHARED) $(C_TARGET) -o $(OUTPUT_DIR)/$(LIB_NAME)