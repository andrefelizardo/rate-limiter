import http from "k6/http";
import { sleep } from "k6";

export const options = {
  scenarios: {
    // low_load: {
    //   executor: "constant-arrival-rate",
    //   rate: 5,
    //   timeUnit: "1s",
    //   duration: "1m",
    //   preAllocatedVUs: 10,
    //   maxVUs: 20,
    // },
    // ramp_up: {
    //   executor: "ramping-arrival-rate",
    //   startRate: 500,
    //   timeUnit: "1s",
    //   preAllocatedVUs: 2000,
    //   maxVUs: 20000,
    //   gracefulStop: "5s",
    //   stages: [
    //     { target: 1000, duration: "20s" },
    //     { target: 3000, duration: "20s" },
    //     { target: 4000, duration: "30s" },
    //     { target: 4500, duration: "30s" },
    //     { target: 5000, duration: "30s" },
    //     { target: 6000, duration: "30s" },
    //     { target: 0, duration: "10s" },
    //   ],
    // },
    steady: {
      executor: "constant-arrival-rate",
      rate: 1800,
      timeUnit: "1s",
      duration: "1m",
      preAllocatedVUs: 8000,
      maxVUs: 20000,
    },
  },
  thresholds: {
    http_req_failed: ["rate<0.01"],
    http_req_duration: ["p(95)<3500"],
    dropped_iterations: ["count<100"],
  },
};

export default function () {
  http.get("http://localhost:3333/test", { timeout: "10s" });
}
