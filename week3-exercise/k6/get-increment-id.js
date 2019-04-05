import http from "k6/http";
import { check } from "k6";

var allUniqueIds = {};
export default function() {  
  let res = http.get("http://host.docker.internal:8081/get-increment-id");
  check(res, {
    "status was 200": (r) => r.status == 200,
    "transaction time OK": (r) => r.timings.duration < 50,
    "data should not be dup": (r) => {
      var incre = r.json().incre;
      if (allUniqueIds[incre]) {
        return false;
      }
      allUniqueIds[incre] = true;
      return true;
    }
  });
};