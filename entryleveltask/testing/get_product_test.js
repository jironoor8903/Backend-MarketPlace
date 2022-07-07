// import http from 'k6/http';
// import { sleep } from 'k6';

// export let options ={
//     insecureSkipTLSVerify: true,
//     noConnectionReuse: false,
//     stages : [
//         {duration: '10s', target : 100},
//         {duration: '10s', target : 100},
//         {duration: '10s', target : 0},
//     ],

//     thresholds: {
//         http_req_duration: ['p(99)<150'],
//     },
// };

// export default() => {
//     http.get("http://localhost:8080/products")
//     http.get("http://localhost:8080/products/category/pants")
//     http.get("http://localhost:8080/products/view/1")
//     // http.get("http://localhost:8080/products/title/White Pants")

//     sleep(1);

// };
import http from "k6/http";

export const options = {
  vus: 2000,
  duration: '1m'
}

export default function() {
  // pick random id
  const id = Math.floor(Math.random() * 100000)

  http.get(`http://localhost:8080/products`);
}