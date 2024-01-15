import requests

url = "http://192.168.50.202:8080/run/?name=name&image=image&tag=tag"

params = {
    "name": "name",
    "image": "image",
    "tag": "tag"
}

res = requests.post(url, json = params)
print(res.json())