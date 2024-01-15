FROM ubuntu:20.04

# 설치 시, 사용자 입력 요구 방지 
# ENV DEBIAN_FRONTEND=noninteractive

# 다운로드 속도 빨라지게 하기 위함 
RUN sed -i 's@archive.ubuntu.com@kr.archive.ubuntu.com@g' /etc/apt/sources.list

# container에서 실행 
RUN mkdir /image_process

# container에서 작업 수행하고자 하는 디렉토리 경로 
WORKDIR /image_process
RUN apt update -y
RUN apt install software-properties-common -y
RUN add-apt-repository ppa:deadsnakes/ppa -y
RUN apt install python3.9 -y

# [err] -y 로 설치 확인 여부 안해줘서 설치할 때 계속 오류 발생 
RUN apt install python3-pip -y

RUN pip install fastapi
RUN pip install uvicorn
RUN pip install pydantic

# COPY [local에 위치하는 복사하고자 하는 파일 경로] [container내에 복사]
# 없으면 오류 발생 
COPY ./app /image_process/app

# 포트포워딩 [ $docker run -it -p 8080:8000 ]가능 -> 8080으로 로컬에서 들어와야 함 
CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "8000"]

## CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "8000"] 인경우
### [ $docker run -it -p 8000:8000 ]
### [ $docker run -it -p 8080:8000 ] 모두 가능 local에서만 8080으로 들어오면 됨

## CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "80"] 인경우
### [ $docker run -it -p 8000:80 ]
### [ $docker run -it -p 8080:80 ]가능

# CMD ["uvicorn", "app.main:app", "--reload"]
## 안됨 