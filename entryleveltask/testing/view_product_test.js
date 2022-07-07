import http from "k6/http";

export const options = {
  vus: 2000,
  duration: '1m'
}

export default function() {
  // pick random id
  const id = Math.floor(Math.random() * 100000)

  http.get(`http://localhost:8080/products/view/${id}`);
}