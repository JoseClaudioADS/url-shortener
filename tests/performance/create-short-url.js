import http from 'k6/http';
import { sleep, check } from 'k6';
export const options = {
  vus: 1160,
  duration: '1s',
};
export default function () {
  const res = http.post('http://localhost:3000', JSON.stringify({
    originalUrl: "http://google.com"
  }), {
    headers: { 'Content-Type': 'application/json' },
  });
  check(res, {
    'is status 201': (r) => r.status === 201,
  });
  sleep(1);
}