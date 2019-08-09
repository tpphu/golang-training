import http from "k6/http";
import { check } from "k6";
import { group } from "k6";

// k6 run --vus 10 --duration 30s k6.js
var url = "http://localhost:8080/register";
var params = { headers: { "Content-Type": "application/json" } };
var start = new Date();
var end = new Date();
export default function () {
  group("register 1", function() {
    start = new Date(end.getTime());
    start.setHours(start.getHours() + 1);
    end = new Date(start.getTime());
    end.setHours(end.getHours() + 2);
    let payload = JSON.stringify({
      "code": "ABC",
      "discount": 0.05,
      "start": start.toISOString(),
      "end": end.toISOString(),
    });
    http.post(url, payload, params);
  });
}
