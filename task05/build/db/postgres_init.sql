CREATE TABLE ads
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(200) NOT NULL,
    date        TIMESTAMP NOT NULL,
    price       INTEGER NOT NULL,
    description VARCHAR(1000)
);

CREATE TABLE photos
(
    id      SERIAL PRIMARY KEY,
    link    TEXT NOT NULL
);

CREATE TABLE ads_photos
(
    ad_id       INTEGER NOT NULL,
    photo_id    INTEGER NOT NULL,
    is_main     BOOLEAN NOT NULL,
    PRIMARY KEY (ad_id, photo_id),
    FOREIGN KEY (ad_id) REFERENCES ads(id),
    FOREIGN KEY (photo_id) REFERENCES photos(id)
);

-- Example data
INSERT INTO ads(name, date, price, description) VALUES
    ('Квартира по адресу 1', '2021-08-30 19:10:25-07', 12000, 'Обычная квартира в центре мира'),
    ('Квартира по адресу 2', '2021-08-30 19:11:25-07', 25500, 'Роскошные хоромы'),
    ('Машина 1', '2021-08-30 19:12:25-07', 6000, 'Обычная машина'),
    ('Билет в казино', '2021-08-30 19:13:25-07', 1, 'Вы обязательно победите!'),
    ('PC', '2021-08-30 19:14:25-07', 1000, 'Точно игровой (с) консультант'),
    ('Машина 2', '2021-08-30 19:15:25-07', 12000, 'Необычная машина'),
    ('Хлеб', '2021-08-30 19:16:25-07', 1, NULL),
    ('Мышь', '2021-08-30 19:17:25-07', 50, 'Живая или компьютерная?');

INSERT INTO photos(link) VALUES
    ('https://picsum.photos/id/101/200/200'),
    ('https://picsum.photos/id/102/200/200'),
    ('https://picsum.photos/id/103/200/200'),
    ('https://picsum.photos/id/201/200/200'),
    ('https://picsum.photos/id/202/200/200'),
    ('https://picsum.photos/id/203/200/200'),
    ('https://picsum.photos/id/301/200/200'),
    ('https://picsum.photos/id/302/200/200'),
    ('https://picsum.photos/id/303/200/200'),
    ('https://picsum.photos/id/401/200/200'),
    ('https://picsum.photos/id/402/200/200'),
    ('https://picsum.photos/id/403/200/200'),
    ('https://picsum.photos/id/501/200/200'),
    ('https://picsum.photos/id/502/200/200'),
    ('https://picsum.photos/id/503/200/200'),
    ('https://picsum.photos/id/601/200/200'),
    ('https://picsum.photos/id/602/200/200'),
    ('https://picsum.photos/id/603/200/200'),
    ('https://picsum.photos/id/701/200/200'),
    ('https://picsum.photos/id/702/200/200'),
    ('https://picsum.photos/id/703/200/200'),
    ('https://picsum.photos/id/801/200/200'),
    ('https://picsum.photos/id/802/200/200'),
    ('https://picsum.photos/id/803/200/200');


INSERT INTO ads_photos(ad_id, photo_id, is_main) VALUES
    (1, 1, TRUE),
    (1, 2, FALSE),
    (1, 3, FALSE),
    (2, 4, FALSE),
    (2, 5, TRUE),
    (2, 6, FALSE),
    (3, 7, FALSE),
    (3, 8, FALSE),
    (3, 9, TRUE),
    (4, 10, FALSE),
    (4, 11, TRUE),
    (4, 12, FALSE),
    (5, 13, TRUE),
    (5, 14, FALSE),
    (5, 15, FALSE),
    (6, 16, TRUE),
    (6, 17, FALSE),
    (6, 18, FALSE),
    (7, 19, FALSE),
    (7, 20, TRUE),
    (7, 21, FALSE),
    (8, 22, FALSE),
    (8, 23, FALSE),
    (8, 24, TRUE);
