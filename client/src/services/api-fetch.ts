const tokenKey = import.meta.env.VITE_TOKEN_KEY;
const BASE_URI = import.meta.env.VITE_BASE_URI;

export default async function apiFetch(
  endpoint: string,
  { method, headers, body } = {} as RequestInit
) {
  const token = sessionStorage.getItem(tokenKey);

  if (token) {
    headers = {
      Authorization: `Bearer ${token}`,
      ...headers,
    };
  }

  if (body) {
    headers = {
      "Content-Type": "application/json",
      ...headers,
    };
  }

  const config = {
    method: method || (body ? "POST" : "GET"),
    headers,
    body,
  };

  const response = await fetch(BASE_URI + endpoint, config);

  let data;

  if (!response.ok) {
    try {
      data = await response.json();
    } catch (error) {
      throw new Error(response.statusText);
    }
    throw new Error(JSON.stringify(data));
  }

  try {
    data = await response.json();
  } catch (error) {
    data = response.statusText;
  }

  return data;
}
