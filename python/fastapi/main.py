from _typeshed import StrOrBytesPath, SupportsItemAccess
from datetime import datetime, timezone
from fastapi import FastAPI, HTTPException
from random import randint
from sqlalchemy import create_engine
from sqlmodel import SQLModel

# -----------------------------------------------------------------------
sqlite_file_name = "database.db"
sqlite_url = f"sqlite://{sqlite_file_name}"
connect_args = {"check_same_thread": False}
engine = create_engine(sqlite_url, connect_args=connect_args)

def create_db_and_tables():
    SQLModel.metadata.create_all(engine)

def get_session():
    """"""
# -----------------------------------------------------------------------


app = FastAPI(root_path="/api/v1")

# root route
@app.get("/")
async def root():
    return {
        "message":"Hello, World!"
}

data = [
    {
        "id": 1,
        "name": "IBM Bob Hackathon",
        "due_date": datetime.now(timezone.utc),
        "created_at": datetime.now(timezone.utc)

    },
    {
        "id": 2,
        "name": "Devs Meet India",
        "due_date": datetime.now(timezone.utc),
        "created_at": datetime.now(timezone.utc)
    },
    {
        "id": 3,
        "name": "HackerGoa",
        "due_date": datetime.now(timezone.utc),
        "created_at": datetime.now(timezone.utc)
    }
]

# get all campaigns
@app.get("/campaigns")
async def read_campaigns():
    return {
        "campaigns": data
}

# get campaign by id
@app.get("/campaigns/{id}")
async def read_campign(id: int):
    for campaign in data:
        if campaign.get("id") == id:
            return {
                "campaign": campaign
            }
    raise HTTPException(status_code=404)

# add campaign
@app.post("/campaigns")
async def create_campaign(body: dict[str, Any]):
    new_campaign = {
        "id": randint(100, 1000),
        "name": body.get("name"),
        "due_date": body.get("due_date"),
        "created_at": datetime.now(timezone.utc)
    }

    data.append(new_campaign)
    return {
        "campaign": new_campaign
    }

# update campaign
@app.put("/campaigns/{id}")
async def update_campaign(id:int, body:dict[str, Any]):
    for campaign in data:
        if campaign.get("id") == id:
            campaign.update(body)
            return {
                "campaign": campaign
            }
    raise HTTPException(status_code=404)

# delete campaign
@app.delete("/campaigns/{id}")
async def delete_campaign(id: int):
    for campaign in data:
        if campaign.get("id") == id:
            data.remove(campaign)
            return {
                "message": "Campaign deleted"
            }
    raise HTTPException(status_code=404)
