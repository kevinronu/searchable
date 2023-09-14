export type Documents = {
  _shards: Shards;
  hits: Hits;
  timed_out: boolean;
  took: number;
};

export type Shards = {
  failed: number;
  skipped: number;
  successful: number;
  total: number;
};

export type Hits = {
  hits: Hit[];
  max_score: number;
  total: Total;
};

export type Hit = {
  "@timestamp": Date;
  _id: string;
  _index: string;
  _score: number;
  _source: Source;
  _type: string;
};

export type Source = {
  "@timestamp": Date;
  bcc: string[];
  body: string;
  cc: string[];
  date: Date;
  from: string;
  message_id: string;
  subject: string;
  to: string[];
};

export type Total = {
  value: number;
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
