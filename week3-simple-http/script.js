import http from 'k6/http';
import { check } from 'k6';

// k6.io
export let options = {
  vus: 20, // 20 concurrency
  iterations: 300000, // 300k request
};

export default function() {
  let res =  http.get('http://localhost:3000/safe-counter');
  check(res, { 'status was 200': r => r.status == 200 });
}
