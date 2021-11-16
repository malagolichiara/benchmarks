#[macro_use]
extern crate lazy_static;
#[macro_use]
extern crate rocket;

use rocket::form::Form;

use std::collections::HashMap;
use std::sync::Mutex;

lazy_static! {
    static ref HASHMAP: Mutex<HashMap<String, String>> = Mutex::new(HashMap::new());
}

#[get("/items/<idx>")]
fn items(idx: &str) -> String {
    let map = HASHMAP.lock().unwrap();
    format!(
        "{{ \"value\": \"{}\"}}",
        map.get(idx).unwrap_or(&String::from("")),
    )
}

#[derive(FromForm)]
struct Value<'r> {
    value: &'r str,
}

#[post("/items/<idx>", data = "<value>")]
fn new(idx: &str, value: Form<Value<'_>>) -> &'static str {
    if value.value.is_empty() {
        return "Value cannot be empty.";
    } else {
        let mut map = HASHMAP.lock().unwrap();
        map.insert(String::from(idx), String::from(value.value));
        return "OK";
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
