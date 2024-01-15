#[macro_use] extern crate rocket;

use rocket::serde::{json::Json, Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
struct Response {
    nickname: String,
    email: String,
}

#[post("/", format = "json", data = "<response>")]
fn post(response: Json<Response>) -> Json<Response> {
    response
}

#[get("/")]
fn get() -> Json<Response> {
    Json(Response {
        nickname: "cikim".to_string(),
        email: "hanbat.ac.kr".to_string(),
    })
}

#[launch]
fn rocket() -> _ {
    rocket::build().mount("/", routes![get, post])
}
