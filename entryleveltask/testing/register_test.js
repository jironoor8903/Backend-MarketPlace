import http from "k6/http";
import { URL } from 'https://jslib.k6.io/url/1.0.0/index.js';
import { randomString } from 'https://jslib.k6.io/k6-utils/1.1.0/index.js';

export const options = {
  vus: 2000,
  duration: '1m'
}


export default function() {
  const url = 'http://localhost:8080/users/register';
  const payload = JSON.stringify({
    username:"jiro",
    email: `k6_user_${randomString(12)}@${randomString(5)}.ksix`,
    password: 'password',
  });

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  http.post(url, payload, params);
  // pick random queries
}