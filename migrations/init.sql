create table links (
  id serial primary key,
  slug varchar(7) unique not null,
  original_url text not null,
  click_count integer default 0,
  created_at timestamptz default current_timestamp,
  updated_at timestamptz default current_timestamp
);

create index idx_links_slug on links(slug);