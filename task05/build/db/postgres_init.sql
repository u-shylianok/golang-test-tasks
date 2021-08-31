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
    ad_id   INTEGER NOT NULL,
    is_main BOOLEAN NOT NULL,
    link    TEXT NOT NULL,
    FOREIGN KEY (ad_id) REFERENCES ads(id)
);

-- Example data
INSERT INTO ads(name, date, price, description) VALUES
    ('Квартира по адресу 1', '2021-08-30 19:10:25-07', 12000, 'Обычная квартира в центре мира'),
    ('Квартира по адресу 2', '2021-08-30 19:11:25-07', 25500, 'Роскошные хоромы'),
    ('Машина 1', '2021-08-30 19:12:25-07', 6000, 'Обычная машина'),
    ('Билет в казино', '2021-08-30 19:13:25-07', 1, 'Вы обязательно победите!'),
    ('PC', '2021-08-30 19:14:25-07', 1000, 'Точно игровой (с) консультант'),
    ('Машина 2', '2021-08-30 19:15:25-07', 12000, 'Необычная машина'),
    ('Хлеб', '2021-08-30 19:16:25-07', 1, 'Мягкий, вкусный'),
    ('Мышь', '2021-08-30 19:17:25-07', 50, 'Живая или компьютерная?');

INSERT INTO photos(ad_id, is_main, link) VALUES
    (1, TRUE, 'https://picsum.photos/id/101/200/200'),
    (1, FALSE, 'https://picsum.photos/id/102/200/200'),
    (1, FALSE, 'https://picsum.photos/id/103/200/200'),
    (1, FALSE, 'https://picsum.photos/id/104/200/200'),
    (2, TRUE, 'https://picsum.photos/id/201/200/200'),
    (2, FALSE, 'https://picsum.photos/id/202/200/200'),
    (2, FALSE, 'https://picsum.photos/id/203/200/200'),
    (2, FALSE, 'https://picsum.photos/id/204/200/200'),
    (2, FALSE, 'https://picsum.photos/id/205/200/200'),
    (3, FALSE, 'https://picsum.photos/id/301/200/200'),
    (3, TRUE, 'https://picsum.photos/id/302/200/200'),
    (3, FALSE, 'https://picsum.photos/id/303/200/200'),
    (4, FALSE, 'https://picsum.photos/id/401/200/200'),
    (4, TRUE, 'https://picsum.photos/id/402/200/200'),
    (4, FALSE, 'https://picsum.photos/id/403/200/200'),
    (5, FALSE, 'https://picsum.photos/id/501/200/200'),
    (5, FALSE, 'https://picsum.photos/id/502/200/200'),
    (5, TRUE, 'https://picsum.photos/id/503/200/200'),
    (5, FALSE, 'https://picsum.photos/id/504/200/200'),
    (6, TRUE, 'https://picsum.photos/id/601/200/200'),
    (6, FALSE, 'https://picsum.photos/id/602/200/200'),
    (6, FALSE, 'https://picsum.photos/id/603/200/200'),
    (6, FALSE, 'https://picsum.photos/id/604/200/200'),
    (7, TRUE, 'https://picsum.photos/id/701/200/200'),
    (7, FALSE, 'https://picsum.photos/id/702/200/200'),
    (7, FALSE, 'https://picsum.photos/id/703/200/200'),
    (8, FALSE, 'https://picsum.photos/id/801/200/200'),
    (8, FALSE, 'https://picsum.photos/id/802/200/200'),
    (8, TRUE, 'https://picsum.photos/id/803/200/200');
