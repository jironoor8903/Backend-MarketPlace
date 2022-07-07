import http from 'k6/http';
import { sleep } from 'k6';

export let options = {
    insecureSkipTLSVerify: true,
    noConnectionReuse: false,
  stages: [
    { duration: '10s', target: 100 }, // below normal load
    { duration: '1m', target: 100 },
    { duration: '10s', target: 1400 }, // normal load
    { duration: '3m', target: 1400 },
    { duration: '10s', target: 100 }, // around the breaking point
    { duration: '3m', target: 100 },
    { duration: '10s', target: 0 }, // beyond the breaking point
    // scale down. Recovery stage.
  ],
};


  const API_BASE_URL = 'http://localhost:8080'; // make sure this is not production


export default () => {
    http.batch([
        ['GET', '${API_BASE_URL}/products'],
        ['GET', '${API_BASE_URL}/products/category/pants']
    ]);


  sleep(1);
};
