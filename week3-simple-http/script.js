import http from 'k6/http';
import { check } from 'k6';


export let options = {
  vus: 20,
  iterations: 300000,
};

export default function() {
  let res =  http.get('http://localhost:3001/safe-counter');
  check(res, { 'status was 200': r => r.status == 200 });
}