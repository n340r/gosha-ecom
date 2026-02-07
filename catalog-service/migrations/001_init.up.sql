CREATE TABLE products (
  id INT PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT NOT NULL,
  manufacturer TEXT NOT NULL,
  image_url TEXT NOT NULL,
  active BOOLEAN NOT NULL DEFAULT true,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE offers (
  id INT PRIMARY KEY,
  product_id INT NOT NULL REFERENCES products(id) ON DELETE CASCADE,
  name TEXT NOT NULL,
  price INT NOT NULL,
  images TEXT[] NOT NULL DEFAULT '{}',
  size TEXT,
  color TEXT,
  quantity INT NOT NULL DEFAULT 0,
  active BOOLEAN NOT NULL DEFAULT true
);

-- seed 1 product + a couple offers (so frontend works immediately)

INSERT INTO products (id, name, description, manufacturer, image_url, active)
VALUES
  (74, 'порно студия[black]', 'оверсайз фит/', 'SERVISEX',
  'https://s3-s1.retailcrm.tech/ru-central1/retailcrm/goshamartynovich-3d1331c972115a9776c29ace80fb18eb/product/68cc60f1c9f5e-porno-black_1-copy.jpg',
  true
  );

INSERT INTO offers (id, product_id, name, price, images, size, quantity, active)
VALUES
  (199, 74, 'порно студия[black], S', 5000,
   ARRAY['https://s3-s1.retailcrm.tech/ru-central1/retailcrm/goshamartynovich-3d1331c972115a9776c29ace80fb18eb/product/68cc60f1c9f5e-porno-black_1-copy.jpg'],
   'S', 3, true
  ),
  (200, 74, 'порно студия[black], M', 5000,
   ARRAY['https://s3-s1.retailcrm.tech/ru-central1/retailcrm/goshamartynovich-3d1331c972115a9776c29ace80fb18eb/product/68cc60f1c9f5e-porno-black_1-copy.jpg'],
   'M', 4, true
  );
