export type SearchResult = {
  total_found: number;
  documents: Document[];
};

export type Document = {
  id: string;
  subject: string;
  from: string;
  to: string[];
  cc: string[];
  body: string;
  date: Date;
};

type Pagination = {
  from: number;
  size: number;
};

export type Body = {
  query: string;
  sort: string;
  pagination: Pagination;
};

type HttpMethod =
  | "GET"
  | "POST"
  | "PUT"
  | "PATCH"
  | "DELETE"
  | "OPTIONS"
  | "HEAD";

export type RequestOptions = {
  method?: HttpMethod;
  headers?: Record<string, string>;
  body?: Body;
};
