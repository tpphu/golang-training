import http from 'k6/http';
import { check } from 'k6';

// k6.io
export let options = {
  vus: 10, // 30 concurrency
  duration: '5s', // 10 seconds 
};

export default function() {
  let res =  http.get('http://localhost:8088/order');
  check(res, { 'status was 200': r => r.status == 200 });
  check(res, { 'total should 560k': r => r.body == "Tong so tien cua 3 san pham: 560000" });
}