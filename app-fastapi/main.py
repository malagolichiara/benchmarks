from typing import Optional

from fastapi import FastAPI, Form

app = FastAPI()
values = {}


@app.get("/")
async def read_root():
    return {"Hello": "World"}


@app.get("/items/{item_id}")
async def read_item(item_id: int, q: Optional[str] = None):
    return {"item_id": item_id, "value": values.get(item_id, None)}


@app.post("/items/{item_id}")
async def write_item(
    item_id: int,
    value: str = Form(...),
    q: Optional[str] = None,
):
    values[item_id] = value
