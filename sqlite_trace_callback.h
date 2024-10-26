#include <sqlite3ext.h>

int _sqlite3_trace_v2(
    sqlite3 *db,
    unsigned uMask,
    int (*xCallback)(unsigned, void *, void *, void *),
    void *pCtx);
void _SQLITE_EXTENSION_INIT2(sqlite3_api_routines *pApi);
typedef void (*closure)();
int go_trace_v2_callback(unsigned t, void *c, void *p, void *x);
const char *_sqlite3_sql(sqlite3_stmt *);
