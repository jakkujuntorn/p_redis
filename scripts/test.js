import http from "k6/http";

export let options = {
  vus: 5,
  duration: "5s",
};

export default function () {
  // ******* ยิงจากข้างนอก *******
  // http.get('http://localhost:8000/hello')

  // ********** ยิงจากข้างใน **********
  http.get("http://host.docker.internal:8000/products");
}
