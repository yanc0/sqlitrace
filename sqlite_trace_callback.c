#include <sqlite3ext.h>

SQLITE_EXTENSION_INIT1
void _SQLITE_EXTENSION_INIT2(sqlite3_api_routines *pApi)
{
    SQLITE_EXTENSION_INIT2(pApi)
}


int _sqlite3_trace_v2(
    sqlite3 *db,
    unsigned uMask,
    int (*xCallback)(unsigned, void *, void *, void *),
    void *pCtx)
{
    return sqlite3_trace_v2(db, uMask, xCallback, pCtx);
}


const char *_sqlite3_sql(sqlite3_stmt *stmt)
{
    return sqlite3_sql(stmt);
}