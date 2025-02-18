create table links (
  id serial primary key,
  slug varchar(7) unique not null,
  original_url text not null,
  click_count integer default 0,
  created_at timestamptz default current_timestamp,
  updated_at timestamptz default current_timestamp
);

create index idx_links_slug on links(slug);

INSERT INTO links (slug, original_url, click_count, created_at, updated_at) VALUES
('abc123', 'http://example.com', 10, current_timestamp, current_timestamp),
('def456', 'http://anotherexample.com', 5, current_timestamp, current_timestamp),
('ghi789', 'http://yetanotherexample.com', 0, current_timestamp, current_timestamp),
('jkl012', 'http://example.org', 20, current_timestamp, current_timestamp),
('mno345', 'http://example.net', 15, current_timestamp, current_timestamp);