from fastapi import FastAPI
from pydantic.main import BaseModel

app = FastAPI()

# get요청으로 name, image, tag 데이터 전달 하면
# docker run --it name image:tag의 String으로 반환

class Container(BaseModel):
  name: str
  image: str
  tag: int
  
@app.get("/")
def container():
  return "HI"
  
# get 요청 하기
@app.get('/run')
def container_with_anme(name: str, image: str, tag: int):
  return f'name : {name}, iamge : {image}, tag : {tag}'

@app.post("/run/post")
def container_post(request: Container):
  return 'name{}, iamge{}, tag{}'.format(request.name, request.image, request.tag)