#[macro_use]
extern crate rocket;

use rocket::form::Form;
use rocket::response::{Flash, Redirect};

#[get("/items/<idx>")]
fn items(idx: &str) -> String {
    format!("{}", idx)
}

#[derive(FromForm)]
struct Value<'r> {
    value: &'r str,
}

#[post("/items/<idx>", data = "<value>")]
fn new(idx: &str, value: Form<Value<'_>>) -> Flash<Redirect> {
    if value.value.is_empty() {
        Flash::error(Redirect::to(uri!(index)), "Cannot be empty.")
    } else {
        Flash::success(Redirect::to(uri!(index)), format!("{} Task added.", idx))
    }
}

#[get("/")]
fn index() -> &'static str {
    "Hello, world!"
}

#[launch]
fn rocket() -> _ {
    rocket::build().mount("/", routes![index, new, items])
}
