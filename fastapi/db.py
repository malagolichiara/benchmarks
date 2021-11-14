import asyncio
import sqlite3

DB = "/tmp/fastapi_test.sqlite3"
conn = sqlite3.connect(DB)


async def fetchall_async(conn, query, params=None):
    loop = asyncio.get_event_loop()
    return await loop.run_in_executor(
        None, lambda: conn.cursor().execute(query, params).fetchall()
    )


async def get_item(_id) -> str:
    sql = "select * from items where id=?"
    return await fetchall_async(conn, sql, (_id,))


async def set_item(_id: int, value: str):
    def write():
        conn.cursor().execute("INSERT OR UPDATE INTO items (id = ?, value = ?);", (_id, value))
        conn.commit()

    loop = asyncio.get_event_loop()
    await loop.run_in_executor(None, lambda: write())


def init():
    curs = conn.cursor()

    curs.execute(
        """
    CREATE TABLE IF NOT EXISTS items (id INTEGER PRIMARY KEY, value TEXT);
    """
    )

    conn.commit()


init()
