#include<uuid/uuid.h>

void getUuid(char *uuid_s) {
	uuid_t uuid;

	uuid_generate(uuid);
	uuid_unparse(uuid, uuid_s);
}