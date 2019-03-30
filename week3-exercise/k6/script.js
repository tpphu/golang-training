import http from "k6/http";
import { check, sleep } from "k6";

export default function() {
  let res = http.get("http://192.168.40.138:8081/note/1");
  check(res, {
    "status was 200": (r) => r.status == 200,
    "transaction time OK": (r) => r.timings.duration < 10
  });
};