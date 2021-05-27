import http from 'k6/http';
import { check, sleep } from 'k6';

export default function () {
  let res = http.get('https://httpbin.org/json');
    check(res, {
        'status was 200': (r) => r.status == 200,
        'exist slideshow': (r) => r.json().slideshow != "",
        'contain title': (r) => r.json().slideshow.title.indexOf("Example") >= 0,
    })
}
