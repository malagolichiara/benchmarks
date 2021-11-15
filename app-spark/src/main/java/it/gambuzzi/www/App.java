package it.gambuzzi.www;
import static spark.Spark.*;
import java.util.HashMap;

/**
 * Hello world!
 *
 */
public class App 
{
    public static HashMap<String, String> values;

    public static void main(String[] args) {
        values = new HashMap<>();
        values.put("1","one");

        port(8000);
        get("/", (req, res) -> "Hello World");
        get("/items/:idx", (req, res) -> {
                String value = values.get(req.params(":idx"));
                return value;
        });
        post("/items/:idx", (req, res) -> {
                values.put(req.params(":idx"), req.body());
                return "OK";
        });
    }
}
