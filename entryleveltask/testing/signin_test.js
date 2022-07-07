import http from "k6/http";
import { URL } from 'https://jslib.k6.io/url/1.0.0/index.js';
import { randomString } from 'https://jslib.k6.io/k6-utils/1.1.0/index.js';

export const options = {
  vus: 2000,
  duration: '1m'
}


export default function() {
  const url = 'http://localhost:8080/registration';
  const payload = JSON.stringify({
    username: `Jironoor1`,
    email: 'jiro@gmail.com',
    confirm_password: 'jiroiscool'
  });

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  http.post(url, payload, params);
  // pick random queries
}