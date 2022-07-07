import http from "k6/http";

export const options = {
  vus: 200,
  duration: '1m'
}

export default function() {
  // pick random id
  const id = Math.floor(Math.random() * 100000)
  const url = 'http://localhost:8080/products/category/pants'

  http.get(url);
}