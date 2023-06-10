import http from 'k6/http';
import { sleep, check } from 'k6';
export const options = {
  vus: 1160,
  duration: '1s',
};
const baseURL = 'http://localhost:3000';
const numUrls = 30;


export function setup() {
  const shortUrls = [];
  for (let i = numUrls; i > shortUrls.length; i--) {
    const payload = {
      originalUrl: `https://www.example.com`,
    };
  
    const headers = { 'Content-Type': 'application/json' };
  
    const createResponse = http.post(
      `${baseURL}`,
      JSON.stringify(payload),
      { headers }
    );
  
    if (createResponse.status === 201) {
      const shortUrl = createResponse.json('shortUrl');
      
      shortUrls.push(shortUrl);
    } else {
      console.error(`Error creating shorted URL: ${createResponse.status}`);
    }
  
  }

  return {shortUrls};
}

export default function (data) {
  const randomIndex = Math.floor(Math.random() * data.shortUrls.length);
  const randomUrl = `${baseURL}/${data.shortUrls[randomIndex]}`;

  const getRequest = http.get(randomUrl);

  check(getRequest, {
    'is status 200': (r) => r.status === 200,
  });

  sleep(1);
}
