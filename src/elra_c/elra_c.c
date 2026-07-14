#include "elra_c.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sqlite3.h>

static sqlite3* db = NULL;

int elra_init(const char* db_path) {
    if (db_path == NULL) {
        return -1;
    }
    int rc = sqlite3_open(db_path, &db);
    if (rc != SQLITE_OK) {
        return -1;
    }
    const char* schema = 
        "CREATE TABLE IF NOT EXISTS instances ("
        "id TEXT PRIMARY KEY,"
        "name TEXT,"
        "arch TEXT,"
        "mode TEXT,"
        "path TEXT,"
        "pid INTEGER,"
        "auto_inject INTEGER,"
        "status TEXT"
        ");";
    char* err = NULL;
    rc = sqlite3_exec(db, schema, NULL, NULL, &err);
    if (rc != SQLITE_OK) {
        sqlite3_free(err);
        return -1;
    }
    return 0;
}

int elra_create_instance(const ELRAInstance* inst) {
    if (db == NULL || inst == NULL) {
        return -1;
    }
    char sql[1024];
    snprintf(sql, sizeof(sql),
        "INSERT INTO instances (id, name, arch, mode, path, pid, auto_inject, status) "
        "VALUES ('%s', '%s', '%s', '%s', '%s', %d, %d, '%s')",
        inst->id, inst->name, inst->arch, inst->mode,
        inst->path, inst->pid, inst->auto_inject, inst->status);
    char* err = NULL;
    int rc = sqlite3_exec(db, sql, NULL, NULL, &err);
    if (rc != SQLITE_OK) {
        sqlite3_free(err);
        return -1;
    }
    return 0;
}

void elra_close(void) {
    if (db != NULL) {
        sqlite3_close(db);
        db = NULL;
    }
}
