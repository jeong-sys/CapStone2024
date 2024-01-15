FROM ubuntu:20.04
ENV DEBIAN_FRONTEND=noninteractive
RUN sed -i 's@archive.ubuntu.com@kr.archive.ubuntu.com@g' /etc/apt/sources.list
RUN apt update && apt install -y python3.9 \
   python3-pip 
   
COPY ./ /DOCKERFILE-FOLDER

RUN pip install fastapi && pip install "fastapi [all]" 

WORKDIR /DOCKERFILE-FOLDER

EXPOSE 8000

CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--reload"]