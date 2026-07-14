#ifndef ELRA_C_H
#define ELRA_C_H

#include <stddef.h>

#ifdef __cplusplus
extern "C" {
#endif

typedef struct {
    char id[64];
    char name[128];
    char arch[16];
    char mode[16];
    char path[256];
    int pid;
    int auto_inject;
    char status[16];
} ELRAInstance;

int elra_init(const char* db_path);
int elra_create_instance(const ELRAInstance* inst);
int elra_get_instance(const char* id, ELRAInstance* inst);
int elra_list_instances(ELRAInstance* instances, size_t* count);
int elra_update_status(const char* id, const char* status);
int elra_delete_instance(const char* id);
void elra_close(void);

#ifdef __cplusplus
}
#endif

#endif
