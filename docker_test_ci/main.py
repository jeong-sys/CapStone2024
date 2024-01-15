from fastapi import FastAPI, Body, HTTPException
from pydantic import BaseModel

class Item(BaseModel):
      name: str
      image: str
      tag: str


app = FastAPI()


@app.get("/run/")
def item(name:str, image:str, tag:str):
  return f"name : {name}, image : {image}, tag : {tag}" 


@app.post("/run/")
def item(test : Item):
  return f"name : {test.name}, image {test.image}, image {test.tag}"